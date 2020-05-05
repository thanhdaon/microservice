const amqp = require("amqplib");

async function run() {
  const connection = await amqp.connect("amqp://guest:guest@localhost:5672");
  const channel = await connection.createChannel();

  channel.prefetch(1);

  await channel.assertQueue("test", { durable: true });
  console.log("AMQP connected");

  let counter = 1;

  const { consumerTag } = await channel.consume("test", (msg) => {
    console.log(msg.content.toString());
    channel.ack(msg);
    counter = 1;
  });

  const timer = setInterval(() => {
    console.log(counter);
    if (++counter > 10) {
      channel.cancel(consumerTag, console.log);
      console.log("cancel consumer: ", consumerTag);
      clearInterval(timer);
    }
  }, 1000);
}

run().catch(console.log);
