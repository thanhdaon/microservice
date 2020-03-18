const puppeteer = require("puppeteer");

async function run() {
  const browser = await puppeteer.launch();

  const page = await browser.newPage();
  await page.goto("https://www.linkedin.com/login");
  await page.type("#username", "daongocthanh98hy@gmail.com");
  await page.type("#password", "daongocthanh98");
  await page.click("[type=submit]");
  await page.waitFor(3000);
  await page.screenshot({ path: "static/3.png" });
  await page.goto(
    "https://www.linkedin.com/feed/update/urn:li:activity:6634997148599586816/"
  );
  const hrefs = await page.$$eval("a", as => as.map(a => a.href));
  console.log(hrefs);
  const html = await page.evaluate(() => document.body.innerHTML);
  console.log(findAllEmail(html));
  await browser.close();
}

function findAllEmail(html) {
  return html.match(/([a-zA-Z0-9._+-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9._-]+)/gi);
}

run().catch(console.log);
