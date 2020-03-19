const crawlHtml = require("crawler/facebook_m/crawl_html");
const parseHtmlArticlesInGroup = require("crawler/facebook_m/parse_html_articles_in_group");
const parseHtmlUserGroups = require("crawler/facebook_m/parse_html_user_groups");

module.exports = Object.freeze({
  crawlHtml,
  parseHtmlArticlesInGroup,
  parseHtmlUserGroups
});
