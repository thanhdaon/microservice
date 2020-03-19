const puppeteer = require("puppeteer-extra");
const StealthPlugin = require("puppeteer-extra-plugin-stealth");
const { readTextFile, writeTextFile } = require("helper");

async function crawHtml({ cookieFilePath, url, crawedPath }) {
  puppeteer.use(StealthPlugin());

  const options = {
    args: [
      "--user-agent=`Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36`"
    ]
  };
  const browser = await puppeteer.launch(options);
  const page = await browser.newPage();
  await loadFBCookie(page, cookieFilePath);

  await page.goto(url);
  await page.screenshot({ path: "static/fb.png" });
  const html = await page.evaluate(() => document.body.innerHTML);
  await writeTextFile(crawedPath, html);
  await browser.close();
}

async function loadFBCookie(page, pathToCookieFile) {
  const urls = [
    "https://www.facebook.com",
    "https://web.facebook.com",
    "https://m.facebook.com",
    "https://mbasic.facebook.com",
    "https://developers.facebook.com",
    "https://upload.facebook.com",
    "https://mobile.facebook.com",
    "https://business.facebook.com"
  ];
  const cookieText = await readTextFile(pathToCookieFile);
  const ca = cookieText.split(";");
  const promises = ca.map(item => {
    if (item) {
      const name = item.split("=")[0].trim();
      const value = item.split("=")[1].trim();
      const cookies = urls.map(url => ({ url, name, value }));
      return page.setCookie(...cookies);
    }
  });

  await Promise.all(promises);
}

module.exports = crawHtml;
