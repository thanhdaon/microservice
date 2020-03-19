const fs = require("fs");
const puppeteer = require("puppeteer-extra");
const StealthPlugin = require("puppeteer-extra-plugin-stealth");

async function run() {
  const browser = await setupBrowser();
  const page = await browser.newPage();
  await loadFBCookie(page);

  // facebook logined
  await page.goto(
    "https://m.facebook.com/groups/284426014978595?view=permalink&id=2888141024607068"
  );
  await page.screenshot({ path: "static/fb_article.png" });
  const html = await page.evaluate(() => document.body.innerHTML);
  console.log(html);
  await browser.close();
}

async function setupBrowser() {
  puppeteer.use(StealthPlugin());

  const options = {
    args: [
      "--user-agent=`Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36`"
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
  const cookieText = await readTextFile("static/thanhdao_cookie.txt");
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
