import React from "react";
import Link from "next/link";

const SidebarLogo = () => {
  return (
    <div className="gx-layout-sider-header">
      <Link href="/">
        <a className="gx-site-logo">
          <img alt="logo2" src="/images/logo.png" />
        </a>
      </Link>
    </div>
  );
};

export default SidebarLogo;
