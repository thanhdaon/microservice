import { ConfigProvider, Layout } from "antd";
import { IntlProvider } from "react-intl";
import { useSelector } from "react-redux";

import Sidebar from "layout/Sidebar";
import Topbar from "layout/Topbar";

import AppLocale from "lang";

function AppLayout({ children }) {
  const locale = useSelector((state) => state.settings.locale);
  const localConfig = AppLocale[locale.locale];
  return (
    <ConfigProvider locale={localConfig.antd}>
      <IntlProvider locale={localConfig.locale} messages={localConfig.messages}>
        <Layout className="gx-app-layout">
          <Sidebar />
          <Layout>
            <Topbar />
            <Layout.Content className="gx-layout-content gx-container-wrap">
              <div className="gx-main-content-wrapper">{children}</div>
              <Layout.Footer>
                <div className="gx-layout-footer-content">
                  Copyright Company Name © 2020
                </div>
              </Layout.Footer>
            </Layout.Content>
          </Layout>
        </Layout>
      </IntlProvider>
    </ConfigProvider>
  );
}

export default AppLayout;
