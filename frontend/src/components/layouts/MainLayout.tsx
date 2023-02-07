import { DashboardWrapper, LayoutWrapper } from "./styles";
import LeftNavbar from "./LeftNavbar";
import TopNavbar from "./TopNavbar";

const MainLayout = (props: any) => {
  return (
    <LayoutWrapper className="layout-content-navbar">
      <div className="layout-container">
        <LeftNavbar />
        <div className="layout-page">
          <TopNavbar />
          <div className="layout-inner-content">
            <DashboardWrapper>{props.children}</DashboardWrapper>
          </div>
        </div>
      </div>
    </LayoutWrapper>
  );
};

export default MainLayout;
