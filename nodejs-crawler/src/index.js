const {
  crawlHtmlWithAutoScroll,
  parseHtmlUserFriends
} = require("crawler/facebook_m");

async function run() {
  // await crawlHtmlWithAutoScroll({
  //   cookieFilePath: "static/thanhdao_cookie.txt",
  //   crawedPath: "static/crawed.txt",
  //   url: "https://m.facebook.com/friends/center/friends"
  // });

  await parseHtmlUserFriends();
}

run().catch(console.log);
