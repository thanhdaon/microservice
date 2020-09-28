package ratelimit

import "github.com/go-redis/redis/v8"

var (
	enqueueScript = redis.NewScript(`
		-- e.g. evalsha <sha> 0 arena_namespace some_queue payload
		if #ARGV < 3 then
			return error("USAGE: arena queue payload")
		end
		local arena_name = ARGV[1]
		local queue_name = ARGV[2]
		local payload = ARGV[3]
		local rate_queues = arena_name .. ":queues"
		local queue_name = arena_name .. ":" .. queue_name
		-- schedule the job.
		local queue_length = redis.call('RPUSH', queue_name, payload)
		if queue_length == 1 then
		--new queue, so add to rate list with immediate start time.
		redis.call('ZADD', rate_queues, 0, queue_name)
		end
		return queue_length
	`)

	dequeueScript = redis.NewScript(`
		-- e.g. evalsha <sha> 0 arena_namespace, time uuid() 10
		if #ARGV < 4 then
			return error("USAGE: arena current_time txnid reservation_duration")
		end

		local arena_name = ARGV[1]
		local current_time = ARGV[2]
		local txnid = ARGV[3]
		local reservation_duration = ARGV[4]
		local rate_queues = arena_name .. ":queues"
		
		while true do -- loop until there are no ready queues or we find an unempty queue.
			-- pick a queue with a read time from the dawn of time until now, but only one queue.
			local queue_names = redis.call("ZRANGEBYSCORE", rate_queues, 0, current_time, "LIMIT", 0, 1)
			if #queue_names == 0 then
				-- no queue is allowed to run yet.
				return false
			end
			
			local queue_name = queue_names[1]
			local item = redis.call("LPOP", queue_name)
			if item == false then 
				--the queue was empty; remove it from the managed queues.
				redis.call("ZREM", rate_queues, queue_name)
			else
				-- rate_limits is an assumed hash with calls allowed per queue per second.
				-- for sharded throughput, set it to overall_limit / #shards
				local rate_delay = (1000 / redis.call("HGET", "rate_limits", queue_name))
				-- reschedule the next check of this queue
				redis.call("ZADD", rate_queues, current_time + rate_delay, queue_name)
				-- add the item to in-flight transactions
				redis.call("HMSET", "txn:" .. txnid, "data", item, "queue_name", queue_name)
				-- and schedule transaction reaping if not committed in time.
				redis.call("ZADD", "transactions", current_time + reservation_duration, txnid)
				return item
			end
		end    
	`)

	commitScript = redis.NewScript(`
		-- e.g. evalsha <sha> 0 previous_uuid
		if #ARGV < 1 then
			return error("USAGE: txnid")
		end
		local txnid = ARGV[1]
		redis.call("DEL", "txn:" .. txnid)
		return redis.call("ZREM", "transactions", txnid)
	`)

	reapScript = redis.NewScript(`
		-- e.g. evalsha <sha> time.time()
		if #ARGV < 1 then
			return error("USAGE: current_time")
		end
		local time = ARGV[1]
		local txns = "transactions"
		-- find expired transactions
		local expired_txns = redis.call("ZRANGEBYSCORE", txns, 0, time)
		for i, txnid in ipairs(expired_txns) do
			local item_key = "txn:" .. txnid
			local data, queue_name = unpack(redis.call("HMGET", item_key, "data", "queue_name"))
			--add the txn back to the input queue
			redis.call("LPUSH", queue_name, data)
			-- remove the in-flight transaction
			redis.call("DEL", item_key)
		end
		--remove the expired transactions from the reaping set.
		return redis.call("ZREMRANGEBYSCORE", txns, 0, time)     
	`)
)
