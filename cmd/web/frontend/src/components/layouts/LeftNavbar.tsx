import { useState } from "react";
import { SidebarMenu } from "./styles";
import { Image } from "react-bootstrap";
import { Link, useLocation } from "react-router-dom";

import encloudLogo from "../../assets/images/encloud-logo.png";

// Images
import padlockIcon from "../../assets/images/padlock.png";
import manageIcon from "../../assets/images/manage.png";
import settingsIcon from "../../assets/images/settings.png";
import uploadIcon from "../../assets/images/upload.png";
import downloadIcon from "../../assets/images/download.png";
import shareIcon from "../../assets/images/share.png";

const LeftNavbar = () => {
  const [openContent, SetOpenContent] = useState(false);
  const [openCompute, SetOpenCompute] = useState(false);
  const location = useLocation();

  const { pathname } = location;

  //Javascript split method to get the name of the path in array
  const splitLocation = pathname.split("/");

  return (
    <>
      <SidebarMenu
        id="layout-menu"
        className="layout-menu menu-vertical menu bg-menu-theme"
      >
        <div className="app-brand">
          <Link to="" className="app-brand-link">
            <Image src={encloudLogo} fluid />
          </Link>
        </div>
        <div className="menu-inner-shadow"></div>
        <ul className="menu-inner py-1">
          <li
            className={
              splitLocation[1] === "" && !splitLocation[2]
                ? "menu-item active"
                : "menu-item"
            }
            onClick={() => {
              SetOpenContent(false);
              SetOpenCompute(false);
            }}
          >
            <Link to="" className="menu-link">
              <Image src={padlockIcon} className="menuIcon" />
              <div>Manage Key Pair</div>
            </Link>
          </li>
          <li
            className={openContent ? "menu-item open" : "menu-item"}
            onClick={() => {
              SetOpenContent(!openContent);
              SetOpenCompute(false);
            }}
          >
            <Link to="#" className="menu-link menu-toggle">
              <Image src={manageIcon} className="menuIcon" />
              <div>Manage Content</div>
            </Link>

            <ul className="menu-sub">
              <li
                className={
                  splitLocation[2] === "upload"
                    ? "menu-item active"
                    : "menu-item"
                }
                onClick={(e) => e.stopPropagation()}
              >
                <Link to="upload" className="menu-link">
                  <Image src={uploadIcon} className="menuIcon" />
                  <div>Upload</div>
                </Link>
              </li>
              <li
                className={
                  splitLocation[2] === "list" ? "menu-item active" : "menu-item"
                }
                onClick={(e) => e.stopPropagation()}
              >
                <Link to="list" className="menu-link">
                  <Image src={downloadIcon} className="menuIcon" />
                  <div>Retrieve</div>
                </Link>
              </li>
              <li
                className={
                  splitLocation[2] === "retrieve-shared-content"
                    ? "menu-item active"
                    : "menu-item"
                }
                onClick={(e) => e.stopPropagation()}
              >
                <Link to="retrieve-shared-content" className="menu-link">
                  <Image src={shareIcon} className="menuIcon" />
                  <div>Shared Content</div>
                </Link>
              </li>
            </ul>
          </li>
          <li className="menu-item">
            <Link to="configuration" className="menu-link">
              <Image src={settingsIcon} className="menuIcon" />
              <div>Configuration</div>
            </Link>
          </li>
        </ul>
      </SidebarMenu>
    </>
  );
};

export default LeftNavbar;
