const amqp = require('amqplib')

async function run() {
  const connection = await amqp.connect(
      'amqp://congtyio_email_crawler:FQ914bquqmkcW8N5aDhg6qzfIBDNLX8r@congty.io:5672/congtyio_email_crawler');

  const channel = await connection.createChannel();

  await channel.purgeQueue('test')
}

run().catch(console.log);
