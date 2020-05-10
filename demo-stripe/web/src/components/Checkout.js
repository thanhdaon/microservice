import React from "react";
import { useParams } from "react-router-dom";
import { loadStripe } from "@stripe/stripe-js";
import {
  CardElement,
  Elements,
  useElements,
  useStripe,
} from "@stripe/react-stripe-js";

import Product from "components/Product";
import { products } from "data";

const stripePromise = loadStripe("pk_test_6pRNASCoBOKtIshFeQd4XMUh");

function Checkout() {
  const { productId } = useParams();
  console.log(productId);
  return (
    <Elements stripe={stripePromise}>
      <h1 className="px-8 text-3xl font-bold">Checkout</h1>
      <div className="flex items-center justify-around">
        <Product
          buyable={false}
          {...products.find((p) => p.id === productId)}
        />
        <div className="leading-loose ml-10">
          <form className="max-w-xl m-4 p-10 bg-white rounded shadow-xl">
            <p className="text-gray-800 font-medium">Customer information</p>
            <div className>
              <label className="block text-sm text-gray-00" htmlFor="cus_name">
                Name
              </label>
              <input
                className="w-full px-5 py-1 text-gray-700 bg-gray-200 rounded"
                id="cus_name"
                name="cus_name"
                type="text"
                required
                placeholder="Your Name"
                aria-label="Name"
              />
            </div>
            <div className="mt-2">
              <label className=" block text-sm text-gray-600">Address</label>
              <input
                className="w-full px-2 py-2 text-gray-700 bg-gray-200 rounded"
                id="cus_email"
                name="cus_email"
                type="text"
                required
                placeholder="Street"
                aria-label="Email"
              />
            </div>

            <div className="inline-block mt-2 w-1/2 pr-1">
              <label className="hidden block text-sm text-gray-600">
                Country
              </label>
              <input
                className="w-full px-2 py-2 text-gray-700 bg-gray-200 rounded"
                id="cus_email"
                name="cus_email"
                type="text"
                required
                placeholder="Country"
                aria-label="Email"
              />
            </div>
            <div className="inline-block mt-2 -mx-1 pl-1 w-1/2">
              <label
                className="hidden block text-sm text-gray-600"
                htmlFor="cus_email"
              >
                Zip
              </label>
              <input
                className="w-full px-2 py-2 text-gray-700 bg-gray-200 rounded"
                id="cus_email"
                name="cus_email"
                type="text"
                required
                placeholder="Zip"
                aria-label="Email"
              />
            </div>
            <p className="mt-4 text-gray-800 font-medium">
              Payment information
            </p>
            <div className>
              <label className="block text-sm text-gray-600" htmlFor="cus_name">
                Card
              </label>
              <input
                className="w-full px-2 py-2 text-gray-700 bg-gray-200 rounded"
                id="cus_name"
                name="cus_name"
                type="text"
                required
                placeholder="Card Number MM/YY CVC"
                aria-label="Name"
              />
            </div>
            <div className="mt-4">
              <button
                className="px-4 py-1 text-white font-light tracking-wider bg-gray-900 rounded"
                type="submit"
              >
                $3.00
              </button>
            </div>
          </form>
        </div>
      </div>
    </Elements>
  );
}

function CheckoutForm() {
  const stripe = useStripe();
  const elements = useElements();

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (!stripe || !elements) {
      return;
    }

    const cardElement = elements.getElement(CardElement);
    const { error, paymentMethod } = await stripe.createPaymentMethod({
      type: "card",
      card: cardElement,
    });

    if (error) {
      console.log("[error]", error);
    } else {
      console.log("[PaymentMethod]", paymentMethod);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="w-1/2">
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
      <button type="submit" disabled={!stripe}>
        Pay
      </button>
    </form>
  );
}

export default Checkout;
