import { Product } from "components";
import { products } from "data";

function Products() {
  return (
    <div className="w-full flex justify-center items-center">
      {products.map((product, i) => (
        <Product key={`index-${i}`} {...product} />
      ))}
    </div>
  );
}

export default Products;
