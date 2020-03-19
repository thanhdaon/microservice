const fs = require("fs");
const cheerio = require("cheerio");

async function run() {
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
