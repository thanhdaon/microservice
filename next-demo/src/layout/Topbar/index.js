import React from "react";
import Link from "next/link";
import { Button, Dropdown, Layout, Menu, message, Popover } from "antd";
import Icon from "@ant-design/icons";

import { useDispatch, useSelector } from "react-redux";

import CustomScrollbars from "components/CustomScrollbars";
import SearchBox from "components/SearchBox";
import UserInfo from "components/UserInfo";
import AppNotification from "components/AppNotification";
import MailNotification from "components/MailNotification";
import HorizontalNav from "layout/Topbar/HorizontalNav";

import { SettingActions } from "app-redux/settings";
import languages from "layout/Topbar/languages";

const menu = (
  <Menu onClick={() => message.info("Click on menu item.")}>
    <Menu.Item key="1">Products</Menu.Item>
    <Menu.Item key="2">Apps</Menu.Item>
    <Menu.Item key="3">Blogs</Menu.Item>
  </Menu>
);

function TopBar() {
  const dispatch = useDispatch();
  const locale = useSelector((state) => state.settings.locale);
  const navCollapsed = useSelector((state) => state.settings.navCollapsed);

  const languageMenu = () => (
    <CustomScrollbars className="gx-popover-lang-scroll">
      <ul className="gx-sub-popover">
        {languages.map((language) => (
          <li
            className="gx-media gx-pointer"
            key={language.languageId}
            onClick={(e) => dispatch(SettingActions.switchLanguage(language))}
          >
            <i className={`flag flag-24 gx-mr-2 flag-${language.icon}`} />
            <span className="gx-language-text">{language.name}</span>
          </li>
        ))}
      </ul>
    </CustomScrollbars>
  );

  return (
    <div className="gx-header-horizontal gx-header-horizontal-dark gx-inside-header-horizontal">
      <Layout.Header className="gx-header-horizontal-main">
        <div className="gx-container">
          <div className="gx-header-horizontal-main-flex">
            <div className="gx-d-block gx-d-lg-none gx-linebar gx-mr-xs-3 6e">
              <i
                className="gx-icon-btn icon icon-menu"
                onClick={() => {
                  dispatch(
                    SettingActions.toggleCollapsedSideNav(!navCollapsed)
                  );
                }}
              />
            </div>
            <Link href="/">
              <img
                alt=""
                className="gx-d-block gx-d-lg-none gx-pointer gx-mr-xs-3 gx-pt-xs-1 gx-w-logo"
                src={"/images/w-logo.png"}
              />
            </Link>
            <Link href="/">
              <img
                alt=""
                className="gx-d-none gx-d-lg-block gx-pointer gx-mr-xs-5 gx-logo"
                src={"/images/logo.png"}
              />
            </Link>

            <div className="gx-header-horizontal-nav gx-header-horizontal-nav-curve gx-d-none gx-d-lg-block">
              <HorizontalNav />
            </div>
            <ul className="gx-header-notifications gx-ml-auto">
              <li className="gx-notify gx-notify-search">
                <Popover
                  overlayClassName="gx-popover-horizantal"
                  placement="bottomRight"
                  content={
                    <div className="gx-d-flex">
                      <Dropdown overlay={menu}>
                        <Button>
                          Category <Icon type="down" />
                        </Button>
                      </Dropdown>
                      <SearchBox
                        styleName="gx-popover-search-bar"
                        placeholder="Search in app..."
                      />
                    </div>
                  }
                  trigger="click"
                >
                  <span className="gx-pointer gx-d-block">
                    <i className="icon icon-search-new" />
                  </span>
                </Popover>
              </li>

              <li className="gx-notify">
                <Popover
                  overlayClassName="gx-popover-horizantal"
                  placement="bottomRight"
                  content={<AppNotification />}
                  trigger="click"
                >
                  <span className="gx-pointer gx-d-block">
                    <i className="icon icon-notification" />
                  </span>
                </Popover>
              </li>

              <li className="gx-msg">
                <Popover
                  overlayClassName="gx-popover-horizantal"
                  placement="bottomRight"
                  content={<MailNotification />}
                  trigger="click"
                >
                  <span className="gx-pointer gx-status-pos gx-d-block">
                    <i className="icon icon-chat-new" />
                    <span className="gx-status gx-status-rtl gx-small gx-orange" />
                  </span>
                </Popover>
              </li>
              <li className="gx-language">
                <Popover
                  overlayClassName="gx-popover-horizantal"
                  placement="bottomRight"
                  content={languageMenu()}
                  trigger="click"
                >
                  <span className="gx-pointer gx-flex-row gx-align-items-center">
                    <i className={`flag flag-24 flag-${locale.icon}`} />
                  </span>
                </Popover>
              </li>
              <li className="gx-user-nav">
                <UserInfo />
              </li>
            </ul>
          </div>
        </div>
      </Layout.Header>
    </div>
  );
}

export default TopBar;
