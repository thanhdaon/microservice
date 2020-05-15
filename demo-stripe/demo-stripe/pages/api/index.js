export default (req, res) => {
  res.statusCode = 200;
  res.json({ msg: "server is OK!" });
};
