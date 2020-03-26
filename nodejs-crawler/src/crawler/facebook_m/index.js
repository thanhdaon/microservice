const crawlHtml = require("crawler/facebook_m/crawl_html");
const crawlHtmlWithAutoScroll = require("crawler/facebook_m/crawl_html_auto_scroll");
const parseHtmlArticlesInGroup = require("crawler/facebook_m/parse_html_articles_in_group");
const parseHtmlUserGroups = require("crawler/facebook_m/parse_html_user_groups");
const parseHtmlUserFriends = require("crawler/facebook_m/parse_html_user_friends");
const demo = require("crawler/facebook_m/demo_auto_scroll");

module.exports = Object.freeze({
  crawlHtml,
  crawlHtmlWithAutoScroll,
  parseHtmlArticlesInGroup,
  parseHtmlUserGroups,
  parseHtmlUserFriends,
  demo
});
