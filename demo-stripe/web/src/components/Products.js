import React from "react";

import Product from "components/Product";

import { products } from "data";

function Checkout() {
  return (
    <div className="w-full flex justify-center items-center">
      {products.map((product, i) => (
        <Product key={`index-${i}`} {...product} />
      ))}
    </div>
  );
}

export default Checkout;
