import { useState } from "react";
import { useRouter } from "next/router";
import { loadStripe } from "@stripe/stripe-js";
import {
  CardElement,
  Elements,
  useElements,
  useStripe,
} from "@stripe/react-stripe-js";
import Link from "next/link";
import swal from "sweetalert";

import { Product } from "components";
import { products } from "data";

const stripePromise = loadStripe("pk_test_QPcSd2ybM4YsJxI5bnrfT2s4005PktZNSp");

function Checkout() {
  const { query } = useRouter();
  const product = products.find((p) => p.id === query.id);

  return (
    <Elements stripe={stripePromise}>
      <Link href="/">
        <a className="p-8 text-3xl underline">Back</a>
      </Link>
      <div className="flex items-center justify-around">
        <Product buyable={false} {...product} />
        <CheckoutForm {...product} />
      </div>
    </Elements>
  );
}

function CheckoutForm({ price = 0 }) {
  const stripe = useStripe();
  const elements = useElements();

  const [loading, setLoading] = useState(false);
  const [customerName, setCustomerName] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);

    if (!stripe || !elements) {
      return;
    }

    const { clientSecret } = await fetchPaymentIntent({
      amount: price * 100,
      currency: "usd",
    });

    const { error, paymentIntent } = await stripe.confirmCardPayment(
      clientSecret,
      {
        payment_method: {
          card: elements.getElement(CardElement),
          billing_details: {
            name: customerName,
          },
        },
      }
    );

    if (error) {
      console.log("[error]", error);
      swal({
        title: "Something wrong",
        text: error.message,
        icon: "error",
      });
    } else {
      console.log("[paymentIntent]", paymentIntent);
      swal({
        title: "Payment Success",
        text: `Payment method: ${paymentIntent.payment_method} \n Payment Indent: ${paymentIntent.id}`,
        icon: "success",
      });
    }
    setLoading(false);
  };

  return (
    <div className="leading-loose ml-10" style={{ width: 600 }}>
      <div className="max-w-xl m-4 p-10 bg-white rounded shadow-xl">
        <p className="text-gray-800 font-medium">Customer information</p>
        <div>
          <label className="block text-sm text-gray-00">Name</label>
          <input
            className="w-full px-5 py-1 text-gray-700 bg-gray-200 rounded"
            type="text"
            placeholder="Your Name"
            value={customerName}
            onChange={(e) => setCustomerName(e.target.value)}
          />
        </div>

        <p className="mt-4 text-gray-800 font-medium">Payment information</p>
        <div>
          <label className="block text-sm text-gray-600">Card</label>
          <form onSubmit={handleSubmit}>
            <CardElement
              options={{
                style: {
                  base: {
                    fontSize: "16px",
                    color: "#424770",
                    "::placeholder": {
                      color: "#aab7c4",
                    },
                  },
                  invalid: {
                    color: "#9e2146",
                  },
                },
              }}
            />

            <button
              type="submit"
              className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-1 px-6 mt-6 rounded"
              disabled={!stripe}
              children={
                loading ? <i className="fas fa-circle-notch fa-spin" /> : "Pay"
              }
            />
          </form>
        </div>
      </div>
    </div>
  );
}

function fetchPaymentIntent({ amount, currency }) {
  return fetch("http://localhost:3000/api/stripe/payment-intent", {
    method: "POST",
    body: JSON.stringify({ amount, currency }),
  }).then((res) => res.json());
}

export default Checkout;
