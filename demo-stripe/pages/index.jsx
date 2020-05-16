import { Product } from "components";
import { products } from "data";

function Products() {
  return (
    <div>
      <h1 className="p-8 text-3xl underline">Online Store</h1>
      <div className="w-full flex justify-around items-center">
        {products.map((product, i) => (
          <Product key={`index-${i}`} {...product} />
        ))}
      </div>
    </div>
  );
}

export default Products;
