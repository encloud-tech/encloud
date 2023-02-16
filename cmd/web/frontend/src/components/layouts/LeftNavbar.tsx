import { SidebarMenu } from "./styles";
import { Image } from "react-bootstrap";
import { Link, useLocation } from "react-router-dom";

import encloudLogo from "../../assets/images/encloud-logo.png";
// Images

import dsMenuImg from "../../assets/images/ds-menu.png";
import dsRefreshImg from "../../assets/images/ds-refresh.png";
import dsupload1Img from "../../assets/images/ds-upload-1.png";
import dsManageImg from "../../assets/images/ds-manage.png";
import { useState } from "react";

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
                  <Image src={dsupload1Img} className="menuIcon" />
                  <div>Upload</div>
                </Link>
              </li>
              <li
                className={
                  splitLocation[2] === "retrieve"
                    ? "menu-item active"
                    : "menu-item"
                }
                onClick={(e) => e.stopPropagation()}
              >
                <Link to="retrieve" className="menu-link">
                  <Image src={dsRefreshImg} className="menuIcon" />
                  <div>Retrieve</div>
                </Link>
              </li>
            </ul>
          </li>
          <li
            className={openCompute ? "menu-item open" : "menu-item"}
            onClick={() => {
              console.log(openCompute);
              SetOpenCompute(!openCompute);
              SetOpenContent(false);
            }}
          >
            <Link to="#" className="menu-link menu-toggle">
              <div>Compute</div>
            </Link>

            <ul className="menu-sub">
              <li
                className={
                  splitLocation[2] === "manage-compute"
                    ? "menu-item active"
                    : "menu-item"
                }
                onClick={(e) => e.stopPropagation()}
              >
                <Link to="manage-compute" className="menu-link">
                  <Image src={dsManageImg} className="menuIcon" />
                  <div>Manage Compute</div>
                </Link>
              </li>
              <li
                className={
                  splitLocation[2] === "list" ? "menu-item active" : "menu-item"
                }
                onClick={(e) => e.stopPropagation()}
              >
                <Link to="list" className="menu-link">
                  <Image src={dsMenuImg} className="menuIcon" />
                  <div>List</div>
                </Link>
              </li>
            </ul>
          </li>
          <li className="menu-item">
            <Link to="configuration" className="menu-link">
              <div>Configuration</div>
            </Link>
          </li>
        </ul>
      </SidebarMenu>
    </>
  );
};

export default LeftNavbar;
