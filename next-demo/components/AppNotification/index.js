import React from "react";
import { Avatar } from "antd";

import CustomScrollbars from "components/CustomScrollbars";

import { notifications } from "components/AppNotification/data";

const AppNotification = () => {
  return (
    <>
      <div className="gx-popover-header">
        <h3 className="gx-mb-0">Notifications</h3>
        <i className="gx-icon-btn icon icon-charvlet-down" />
      </div>
      <CustomScrollbars className="gx-popover-scroll">
        <ul className="gx-sub-popover">
          {notifications.map((notification, index) => (
            <NotificationItem key={index} notification={notification} />
          ))}
        </ul>
      </CustomScrollbars>
    </>
  );
};

const NotificationItem = ({ notification }) => {
  const { icon, image, title, time } = notification;
  return (
    <li className="gx-media">
      <Avatar className="gx-size-40 gx-mr-3" alt={image} src={image} />
      <div className="gx-media-body gx-align-self-center">
        <p className="gx-fs-sm gx-mb-0">{title}</p>
        <i className={`icon icon-${icon} gx-pr-2`} />{" "}
        <span className="gx-meta-date">
          <small>{time}</small>
        </span>
      </div>
    </li>
  );
};

export default AppNotification;
