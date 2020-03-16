const amqp = require("amqplib");

async function run() {
  const conn = await amqp.connect("amqp://guest:guest@localhost:5672");
  const channel = await conn.createChannel();

  await channel.assertQueue("moderation-final", { durable: true });
  await channel.assertQueue("moderation-upload-photo", { durable: true });
  await channel.assertQueue("moderation-image-queue", { durable: true });
  await channel.assertQueue("moderation-manager-image-queue", {
    durable: true
  });
  await channel.assertQueue("moderation-copy-done", { durable: true });
  await channel.assertQueue("moderation-image-waiting-for-crop-queue", {
    durable: true
  });
}

run().catch(console.log);
