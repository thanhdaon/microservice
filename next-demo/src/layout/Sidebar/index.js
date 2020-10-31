import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { Drawer, Layout } from "antd";

import SidebarContent from "layout/Sidebar/SidebarContent";
import { SettingActions } from "app-redux/settings";
import { TAB_SIZE } from "constants/theme-settings";

const Sidebar = () => {
  const dispatch = useDispatch();

  const { navCollapsed, width } = useSelector((state) => state.settings);

  const onToggleCollapsedNav = () => {
    dispatch(SettingActions.toggleCollapsedSideNav(!navCollapsed));
  };

  useEffect(() => {
    const onResize = () => {
      dispatch(SettingActions.updateWindowWidth(window.innerWidth));
    };
    window.addEventListener("resize", onResize);
    return () => {
      window.removeEventListener("resize", onResize);
    };
  }, [dispatch]);

  return (
    <Layout.Sider
      className="gx-app-sidebar gx-collapsed-sidebar gx-layout-sider-dark"
      trigger={null}
      collapsed={false}
      theme="dark"
      collapsible
    >
      {width < TAB_SIZE ? (
        <Drawer
          className="gx-drawer-sidebar gx-drawer-sidebar-dark"
          placement="left"
          closable={false}
          onClose={onToggleCollapsedNav}
          visible={navCollapsed}
        >
          <SidebarContent />
        </Drawer>
      ) : (
        <SidebarContent />
      )}
    </Layout.Sider>
  );
};

export default Sidebar;
