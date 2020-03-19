const fs = require("fs");
const cheerio = require("cheerio");

async function run() {
  const html = await readTextFile("static/fb_article.txt");
  const $ = cheerio.load(html);

  const owner = $("div._4g34 > h3 > span > strong:nth-child(1) > a").text();
  const text = $("div._5rgt._5nk5 > ._3w8y").text();
  const comments = [];

  $("div._333v._45kb")
    .children("div._2a_i")
    .each(function() {
      comments.push({
        actor: $(this)
          .find("div._2b05 > a")
          .text(),
        text: $(this)
          .find("[data-commentid]")
          .text()
      });
    });

  console.log({ owner, text, comments });
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

// function writeTextFile(pathToFile, data) {
//   return new Promise((resolve, reject) => {
//     fs.writeFile(pathToFile, data, error => {
//       if (error) {
//         reject(error);
//       }
//       resolve();
//     });
//   });
// }

run().catch(console.log);
