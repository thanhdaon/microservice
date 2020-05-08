const { Magic } = require("@magic-sdk/admin");
const app = require("express")();

async function run() {
  const magic = new Magic("sk_test_42BECDA74B708D20");

  app.get("/login", async (req, res) => {
    try {
      const DIDToken = req.headers.authorization.substring(7);
      console.log(DIDToken);
      const issuer = magic.token.getIssuer(DIDToken);
      console.log({ issuer });
    } catch (error) {
      console.log(error);
    } finally {
      res.end();
    }
  });

  app.listen(8000, () => {
    console.log("[INFO] App is running on port: 8000");
  });
}

run();
