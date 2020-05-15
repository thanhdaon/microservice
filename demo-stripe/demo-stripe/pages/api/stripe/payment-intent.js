const stripe = require("stripe")(process.env.STRIPE_SECRET_KEY);

export default async (req, res) => {
  const payment = JSON.parse(req.body);

  if (!payment.amount || !payment.currency) {
    res.status(400).send("amount and currency can not null!");
    return;
  }

  const paymentIntent = await stripe.paymentIntents.create({
    amount: payment.amount,
    currency: payment.currency,
    metadata: { integration_check: "accept_a_payment" },
  });

  res.status(200).json({
    clientSecret: paymentIntent.client_secret,
  });
};
