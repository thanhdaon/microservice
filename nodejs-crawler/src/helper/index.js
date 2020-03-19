const fs = require("fs");

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

function writeTextFile(pathToFile, data) {
  return new Promise((resolve, reject) => {
    fs.writeFile(pathToFile, data, error => {
      if (error) {
        reject(error);
      }
      resolve();
    });
  });
}

module.exports = Object.freeze({ readTextFile, writeTextFile });
