import "App.css";

import React from "react";
import {
  CardElement,
  Elements,
  useStripe,
  useElements,
} from "@stripe/react-stripe-js";
import { loadStripe } from "@stripe/stripe-js";

const stripePromise = loadStripe("pk_test_QPcSd2ybM4YsJxI5bnrfT2s4005PktZNSp");

function App() {
  return (
    <Elements stripe={stripePromise}>
      <CheckoutForm />
    </Elements>
  );
}

function CheckoutForm() {
  const stripe = useStripe();
  const elements = useElements();

  async function onSubmit(e) {
    e.preventDefault();

    if (!stripe || !elements) {
      return;
    }

    const response = await fetch("http://localhost:8000/secret");
    const { client_secret } = await response.json();

    const result = await stripe.confirmCardPayment(client_secret, {
      payment_method: {
        card: elements.getElement(CardElement),
        billing_details: {
          name: "D.N.T",
        },
      },
    });

    if (result.error) {
      console.log(result.error.message);
    } else {
      if (result.paymentIntent.status === "succeeded") {
        console.log({ result });
      }
    }
  }

  return (
    <form onSubmit={onSubmit} style={{ width: 500 }}>
      <CardElement />
      <button type="submit" disabled={!stripe}>
        Pay
      </button>
    </form>
  );
}

export default App;
