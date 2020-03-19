const cheerio = require("cheerio");
const { readTextFile } = require("helper");

async function parseHtmlArticlesInGroup() {
  const data = [];
  const html = await readTextFile("static/fb.txt");
  const $ = cheerio.load(html);
  $("article").each(function() {
    const article = cheerio.load($(this).html());
    const dataStore = $(this).attr("data-store");
    const url = article("div._5rgt._5nk5._5msi > a._5msj").attr("href");
    data.push({ ...JSON.parse(dataStore), url });
  });

  console.log(data);
}

module.exports = parseHtmlArticlesInGroup;
