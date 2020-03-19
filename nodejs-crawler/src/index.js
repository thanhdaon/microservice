const { parseHtmlUserGroups } = require("crawler/facebook_m");

async function run() {
  // await crawlHtml({
  //   cookieFilePath: "static/haile_cookie.txt",
  //   crawedPath: "static/crawed.txt",
  //   url: "https://m.facebook.com/groups/"
  // });
  await parseHtmlUserGroups();
}

run().catch(console.log);
