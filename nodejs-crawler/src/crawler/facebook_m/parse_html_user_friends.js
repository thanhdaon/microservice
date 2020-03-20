const cheerio = require("cheerio");
const { readTextFile } = require("helper");

async function parseHtmlUserGroups() {
  const html = await readTextFile("static/crawed.txt");
  const $ = cheerio.load(html);

  const friendsCount = $("div#friends_center_main header._5lm6 > h3").text();
  $("div#friends_center_main div._55wp._7om2").each(function() {
    const name = $(this)
      .find("h3 > a")
      .text();
    const link = $(this)
      .find("h3 > a")
      .attr("href");

    console.log({ name, link });
  });
  console.log({ friendsCount });
}

module.exports = parseHtmlUserGroups;
