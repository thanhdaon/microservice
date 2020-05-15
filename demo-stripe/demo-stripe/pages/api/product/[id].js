import { products } from "data";

export default (req, res) => {
  const { id } = req.query;
  res.status(200).json(products.find((product) => product.id === id));
};
