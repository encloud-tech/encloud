import styled from "styled-components";
import encloudGradient from "../../../assets/images/encloud-gradient.jpg";

export const AuthWrapper = styled.div`
  display: flex;
  justify-content: center;
  flex-direction: column;
  text-align: left;
  height: 100vh;
  background-image: url(${encloudGradient});
  background-size: cover;
  background-position: center center;

  .siteLogo {
    text-align: center;
    margin: 0;
    line-height: 1;
    padding-bottom: 20px;
    max-width: 340px;
    width: 100%;
  }

  .auth-inner {
    width: 450px;
    margin: auto;
    background-color: rgba(255, 255, 255, 0.13);
    box-shadow: 0px 14px 80px rgba(34, 35, 58, 0.2);
    backdrop-filter: blur(10px);
    padding: 40px 55px 45px 55px;
    border-radius: 15px;
    transition: all 0.3s;
    border: 2px solid rgba(255, 255, 255, 0.1);

    .form-control {
      border-color: transparent !important;
      display: block;
      height: 50px;
      width: 100%;
      padding: 0 10px;
      margin-top: 8px;
      font-weight: 400;
      color: #303353;
      box-shadow: 0 0 40px rgba(8, 7, 16, 0.2);

      &:focus {
        border-color: transparent !important;
        box-shadow: none;
      }

      &::placeholder {
        color: #30335390;
      }
    }

    label {
      color: #ffffff;
      text-shadow: 0px 1px 2px #000000;
      display: block;
      margin-top: 30px;
      font-size: 16px;
      font-weight: 500;
    }

    .custom-control-label {
      font-weight: 400;
    }

    button {
      /* background: linear-gradient(90deg, #f76e5c 0%, #cc3366 100%) !important; */
      color: #ffffff;
      background-color: #f76e5c;
      border: none;
      margin-top: 50px;
      width: 100%;
      padding: 15px 0;
      font-size: 18px;
      font-weight: 400;
      border-radius: 8px;
      cursor: pointer;

      &:hover {
        background-color: #cc3366;
      }

      &:focus {
        border: none;
        box-shadow: unset;
        outline: none;
      }

      &:not(:disabled):not(.disabled):active:focus {
        box-shadow: unset;
      }
    }
  }
`;
