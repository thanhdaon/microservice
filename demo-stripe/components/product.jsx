import Link from "next/link";

function Product({ id, name, description, image, price, buyable = true } = {}) {
  return (
    <div className="max-w-xs bg-white shadow-lg rounded-lg overflow-hidden my-10">
      <div className="px-4 py-2">
        <h1 className="text-gray-900 font-bold text-3xl uppercase">{name}</h1>
        <p className="text-gray-600 text-sm mt-1">{description}</p>
      </div>
      <img className="h-56 w-full object-cover mt-2" src={image} alt={name} />
      <div className="flex items-center justify-between px-4 py-2 bg-gray-900">
        <h1 className="text-gray-200 font-bold text-xl">${price}</h1>
        {buyable && (
          <Link as={`/checkout/${id}`} href="/checkout/[id]">
            <button className="px-3 py-1 bg-gray-200 text-sm text-gray-900 font-semibold rounded">
              Buy
            </button>
          </Link>
        )}
      </div>
    </div>
  );
}

export default Product;
