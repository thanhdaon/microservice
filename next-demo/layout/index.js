import { Layout } from "antd";
import { useSelector } from "react-redux";

function AppLayout({ children }) {
  const { width } = useSelector((state) => state.settings);

  return (
    <Layout className="gx-app-layout">
      <Layout>
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
  );
}

export default AppLayout;
