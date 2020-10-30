import React from "react";
import NotificationItem from "./NotificationItem";
import { notifications } from "./data";

const AppNotification = () => {
  return (
    <>
      <div className="gx-popover-header">
        <h3 className="gx-mb-0">Notifications</h3>
        <i className="gx-icon-btn icon icon-charvlet-down" />
      </div>
      <ul className="gx-sub-popover">
        {notifications.map((notification, index) => (
          <NotificationItem key={index} notification={notification} />
        ))}
      </ul>
    </>
  );
};

export default AppNotification;
