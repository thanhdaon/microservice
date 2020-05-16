import { products } from "data";

export default (req, res) => {
  res.status(200).json(products);
};
