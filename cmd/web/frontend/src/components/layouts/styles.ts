import { Navbar } from "react-bootstrap";
import styled from "styled-components";

export const LayoutWrapper = styled.div`
  width: 100%;
  display: flex;
  flex: 1 1 auto;
  align-items: stretch;
  .layout-container {
    width: 100%;
    display: flex;
    flex: 1 1 auto;
    align-items: stretch;
    min-height: 100vh;
  }

  .layout-page {
    padding-left: 16.25rem;
    flex-basis: 100%;
    flex-direction: column;
    width: 0;
    min-width: 0;
    max-width: 100%;
    display: flex;
    flex: 1 1 auto;
    align-items: stretch;

    .layout-inner-content {
      padding-right: 1.625rem;
      padding-left: 1.625rem;
    }
  }
`;

export const NavbarWrapper = styled(Navbar)`
  background: radial-gradient(circle at right bottom, #ffcc83 23%, #d94964 71%);
  padding: 15px 1.625rem;
  position: sticky;
  top: 0;
  z-index: 50;
  height: 70px;

  @media only screen and (max-width: 767.98px) {
    position: fixed;
    width: 100%;
  }

  .siteLogo {
    width: 100%;
    max-width: 150px;
    height: auto;
    filter: brightness(0);
  }
`;

export const DashboardWrapper = styled.div`
  display: flex;
  justify-content: flex-start;
  flex-direction: column;
  text-align: left;
  color: #d94964;
  padding: 15px 0;
`;

export const BoxWrapper = styled.div`
  width: 100%;
  position: relative;
  background: white;
  margin-bottom: 30px;
  border: 3px solid #f76e5c;
  box-shadow: 0px 0px 20px 0px #0000001c;
  border-radius: 15px;

  .content {
    padding: 1em 2em;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: flex-start;
    text-align: left;
    transition: width 1s ease-in-out;
    z-index: 3;

    @media only screen and (max-width: 767.98px) {
      flex-direction: column;
      text-align: left;
    }

    h3 {
      color: #d94964;
      font-size: 1.5em;
      margin-bottom: 0;
      margin-bottom: 15px;
      display: block;
      width: 100%;
      padding-bottom: 0.5em;
      border-bottom-width: 1px;
      border-bottom-style: solid;
      border-image: linear-gradient(
          90deg,
          #d94964 0%,
          #ffcc83 52.08%,
          #d94964 100%
        )
        1 / 1 / 0 stretch;

      @media only screen and (max-width: 991.98px) {
        font-size: 2em;
      }

      @media only screen and (max-width: 767px) {
        font-size: 28px;
      }
    }
  }
`;

export const KeyBoxedContent = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;

  @media only screen and (max-width: 767px) {
    flex-direction: column;
  }

  .boxInner {
    flex: 1;
    padding: 0px 30px 0px 0px;
    position: relative;
    height: 80px;
    display: flex;
    justify-content: flex-start;
    align-items: center;
    cursor: pointer;

    &:not(:last-child):after {
      content: "";
      position: absolute;
      height: 100%;
      top: 0;
      right: 15px;
      width: 2px;
      background-image: linear-gradient(
        180deg,
        rgba(247, 110, 92, 0),
        #dee2e6,
        rgba(247, 110, 92, 0)
      );
    }

    .iconBox {
      background: transparent !important;
      /* background-color: #f76e5c;
            background-image: linear-gradient(90deg, rgb(247, 110, 92) 0%, #cc3366 100%); */
      width: 50px;
      height: 50px;
      display: flex;
      justify-content: center;
      align-items: center;
      border-radius: 8px;
      margin-right: 0.75rem;

      img {
        width: 26px;
        height: auto;
        /* filter: brightness(0) invert(1); */
      }
    }

    p.boxTitle {
      font-weight: 700;
      font-size: 1.5rem;
      color: #d94964;
      line-height: 1;
      margin-top: 0;
      margin-bottom: 0;
    }
  }
`;

export const PageHeader = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 15px 0 30px;

  h2 {
    color: #0a0b33 !important;
    margin: 0;

    span {
      font-size: 1.5rem;
      font-weight: 700 !important;
      line-height: 1;
    }

    .titleIcon {
      width: 24px;
      height: auto;
      margin-right: 15px;
    }
  }
`;

export const SidebarMenu = styled.aside`
  z-index: 1080;
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  margin-right: 0 !important;
  margin-left: 0 !important;
  background-color: #fff !important;
  color: #ffffff;
  flex: 1 0 auto;
  background: radial-gradient(circle at right bottom, #ffcc83 23%, #d94964 71%);

  a {
    color: #ffffff;

    &:hover {
      text-decoration: none;
    }
  }

  .app-brand {
    height: 70px;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 15px;
  }

  &.menu-vertical,
  &.menu-vertical .menu-block,
  &.menu-vertical .menu-inner > .menu-item,
  &.menu-vertical .menu-inner > .menu-header {
    width: 16.25rem;
  }

  &.menu-vertical {
    flex-direction: column;
  }

  .menu-inner-shadow {
    background: linear-gradient(
      #fff 41%,
      rgba(255, 255, 255, 0.11) 95%,
      rgba(255, 255, 255, 0)
    );
    display: none;
    position: absolute;
    top: 4.225rem;
    height: 3rem;
    width: 100%;
    pointer-events: none;
    z-index: 2;
  }

  .menu-inner {
    box-shadow: 0 0.6rem 0.375rem 0 rgb(161 172 184);
    display: flex;
    align-items: flex-start;
    justify-content: flex-start;
    margin: 0;
    padding: 0;
    height: 100%;
    flex-direction: column;
    flex: 1 1 auto;
    color: red;
    position: relative;

    & > .menu-item {
      margin: 0.0625rem 0;

      &.active:before {
        content: "";
        position: absolute;
        right: 0;
        width: 0.25rem;
        height: 2.5rem;
        border-radius: 0.375rem 0 0 0.375rem;
        background: #ffffff;
      }

      .menu-link {
        margin: 0rem 1rem;
        font-size: 0.9375rem;
        padding: 0.625rem 1rem;
        position: relative;
        display: flex;
        align-items: center;
        flex: 0 1 auto;
        transition-duration: 0.3s;
        transition-property: color, background-color;
        border-radius: 0.375rem !important;

        &:hover {
          background-color: rgba(255, 255, 255, 0.25);
        }

        .menuIcon {
          width: 20px;
          height: auto;
          filter: brightness(0) invert(1);
          margin-right: 15px;
        }
      }

      .menu-toggle {
        padding-right: calc(1rem + 1.26em);

        &::after {
          right: 1rem;
          transition-duration: 0.3s;
          transition-property: transform;
          content: "";
          position: absolute;
          top: 50%;
          display: block;
          width: 0.42em;
          height: 0.42em;
          border: 1px solid;
          border-bottom: 0;
          border-left: 0;
          transform: translateY(-50%) rotate(45deg);
        }
      }

      .menu-sub {
        display: none;
        > .menu-item {
          > .menu-link {
            padding-left: 4rem;
            padding-top: 0.625rem;
            padding-bottom: 0.625rem;
            font-size: 0.9375rem;

            &::before {
              content: "";
              position: absolute;
              left: 1.4375rem;
              width: 0.375rem;
              height: 0.375rem;
              border-radius: 50%;
              background-color: #b4bdc6 !important;
              display: none;
            }

            .menuIcon {
              position: absolute;
              left: 1.6rem;
            }
          }
        }
      }

      &.open {
        > .menu-toggle {
          &::after {
            transform: translateY(-50%) rotate(135deg);
          }
        }

        > .menu-sub {
          display: flex;
          flex-direction: column;
          margin: 0;
          padding: 0;
          padding-top: 0.3125rem;
          padding-bottom: 0.3125rem;

          .menu-item {
            &.active {
              > .menu-link {
                color: #ffffff;
                border: 1px solid #ffffff50;
                background-color: rgba(255, 255, 255, 0.25) !important;
              }
            }
          }
        }
      }

      &.active {
        > .menu-link {
          color: #ffffff;
          border: 1px solid #ffffff50;
          background-color: rgba(255, 255, 255, 0.25) !important;
        }
      }
    }
  }
`;
