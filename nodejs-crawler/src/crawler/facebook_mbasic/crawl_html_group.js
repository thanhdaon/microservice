const { readTextFile, writeTextFile } = require("helper");
const cheerio = require("cheerio");

async function run() {
  let counter = 0;
  const articles = [];
  const html = await readTextFile("static/fb.html");
  const $ = cheerio.load(html);
  $("[role='article']").each(function() {
    const article = cheerio.load($(this).html());
    const articleItem = {
      owner: {
        name: article("strong > a.eh").text(),
        link: article("strong > a.eh").attr("href")
      },
      text: article("div.en > span > p").text(),
      dataFt: $(this).attr("data-ft")
    };
    articles.push(articleItem);
    writeTextFile(`static/articles/article-${++counter}.html`, $(this).html());
  });

  console.log("article: ", articles.length);
  console.log(articles);
}

run().catch(console.log);
