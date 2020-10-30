import React from "react";
import NotificationItem from "./NotificationItem";
import { notifications } from "./data";

const MailNotification = () => {
  return (
    <>
      <div className="gx-popover-header">
        <h3 className="gx-mb-0">Messages</h3>
        <i className="gx-icon-btn icon icon-charvlet-down" />
      </div>
      <div className="gx-popover-scroll">
        <ul className="gx-sub-popover">
          {notifications.map((notification, index) => (
            <NotificationItem key={index} notification={notification} />
          ))}
        </ul>
      </div>
    </>
  );
};

export default MailNotification;