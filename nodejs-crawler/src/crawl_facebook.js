const fs = require("fs");
const puppeteer = require("puppeteer-extra");
const StealthPlugin = require("puppeteer-extra-plugin-stealth");

async function run() {
  const browser = await setupBrowser();
  const page = await browser.newPage();
  await loadFBCookie(page);

  // facebook logined
  await page.goto("https://mbasic.facebook.com/groups/1902499916677004");
  await page.screenshot({ path: "static/fb.png" });
  const html = await page.evaluate(() => document.body.innerHTML);
  console.log(html);
  await browser.close();
}

async function setupBrowser() {
  puppeteer.use(StealthPlugin());

  const options = {
    args: [
      "--user-agent=`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36`"
    ]
  };

  return await puppeteer.launch(options);
}

async function loadFBCookie(page) {
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
  const cookieText = await readTextFile("static/cookie.txt");
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

function readTextFile(pathToFile) {
  return new Promise((resolve, reject) => {
    fs.readFile(pathToFile, "utf8", (error, contents) => {
      if (error) {
        reject(error);
      }
      resolve(contents);
    });
  });
}

run().catch(console.log);
