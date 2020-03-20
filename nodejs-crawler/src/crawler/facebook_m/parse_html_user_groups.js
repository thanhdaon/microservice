const cheerio = require("cheerio");
const { readTextFile } = require("helper");

async function parseHtmlUserFriends() {
  const html = await readTextFile("static/fb.txt");
  const $ = cheerio.load(html);

  const groups = [];

  $("div._4g34 > div")
    .children("div._7hkf")
    .each(function() {
      const avatar = $(this)
        .find("img")
        .attr("src");
      const name = $(this)
        .find("div._4ik4 > div._52je")
        .text();
      const link = $(this)
        .find("a._7hkg")
        .attr("href");
      groups.push({ avatar, name, link });
    });

  console.log(groups);
}

module.exports = parseHtmlUserFriends;
