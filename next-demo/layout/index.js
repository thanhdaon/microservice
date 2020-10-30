import { ConfigProvider, Layout } from "antd";
import { IntlProvider } from "react-intl";
import { useSelector } from "react-redux";

import Sidebar from "layout/Sidebar";
import Topbar from "layout/Topbar";

import AppLocale from "lngProvider";

function AppLayout({ children }) {
  const { locale } = useSelector((state) => state.settings);
  const currentLocale = AppLocale[locale.locale];
  return (
    <ConfigProvider locale={currentLocale.antd}>
      <IntlProvider
        locale={currentLocale.locale}
        messages={currentLocale.messages}
      >
        <Layout className="gx-app-layout">
          <Sidebar />
          <Layout>
            <Topbar />
            <Layout.Content className="gx-layout-content gx-container-wrap">
              <div className="gx-main-content-wrapper">{children}</div>
            </Layout.Content>
            <Layout.Footer>
              <div className="gx-layout-footer-content">
                Copyright Company Name © 2020
              </div>
            </Layout.Footer>
          </Layout>
        </Layout>
      </IntlProvider>
    </ConfigProvider>
  );
}

export default AppLayout;
