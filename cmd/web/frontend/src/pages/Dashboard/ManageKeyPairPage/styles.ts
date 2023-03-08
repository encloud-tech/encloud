import styled from "styled-components";

export const KeyBoxedContent = styled.div`
  width: 100%;
`;
export const KeyPairsSection = styled.div`
  color: #cc3366;

  h4 {
    color: #cc3366;
  }

  &.separator {
    position: relative;
    &:after {
      content: "";
      position: absolute;
      right: 0px;
      top: 0;
      height: 100%;
      width: 2px;
      border-radius: 10%;
      background: radial-gradient(
        circle at right bottom,
        #d94964,
        #ffcc83,
        #d94964
      );
    }
  }
`;

export const KeyPairs = styled.div`
  color: #cc3366;
  display: flex;
  justify-content: flex-start;
  align-items: center;
  margin-top: 20px;
`;

export const KeyBox = styled.label`
  display: block;
  position: relative;
  padding-left: 35px;
  margin-bottom: 12px;
  cursor: pointer;
  font-size: 16px;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
  margin-right: 15px;

  input {
    position: absolute;
    opacity: 0;
    cursor: pointer;
  }

  span:not(.checkmark) {
    line-height: 1;
  }

  .checkmark {
    position: absolute;
    top: 0;
    left: 0;
    height: 25px;
    width: 25px;
    background-color: #cc33662e;
    border-radius: 50%;
  }

  &:hover input ~ .checkmark {
    background-color: #ef9476;
  }

  input:checked ~ .checkmark {
    background-color: #f76e5c;
  }

  .checkmark:after {
    content: "";
    position: absolute;
    display: none;
  }

  input:checked ~ .checkmark:after {
    display: block;
  }

  .checkmark:after {
    top: 9px;
    left: 9px;
    width: 8px;
    height: 8px;
    border-radius: 50%;
    background: white;
  }
`;

export const StepBoxWrapper = styled.div`
  border: 3px solid #f76e5c;
  border-radius: 24px;
  margin-bottom: 50px;
  background: white;
  padding: 20px;
  box-shadow: 0px 0px 10px 0px #e5e5e5;

  &.inactive {
    position: relative;
    z-index: 1;
    overflow: hidden;
    opacity: 0.2;

    &::before {
      content: "";
      position: absolute;
      top: 0;
      left: 0;
      background-color: #ffffff;
      width: 100%;
      height: 100%;
      opacity: 0.6;
      z-index: 2;
    }
  }
`;

export const StepHeader = styled.div`
  display: flex;
  align-items: center;

  .stepTitle {
    color: #cc3366;

    font-size: 1.5rem;
    font-weight: 500;
    line-height: 120%;
  }

  .right-part {
    display: flex;
    align-items: center;
    justify-content: flex-end;
    margin-left: auto;

    span {
      font-size: 16px;
      line-height: 175%;
      font-weight: 600;
      color: #f76e5c;
    }
  }
`;

export const StepBody = styled.div`
  padding-top: 30px;
  .subHeader {
    margin-top: 10px;
    margin-bottom: 20px;
    font-size: 20px;
    opacity: 0.6;
    color: #0a0b33;
    line-height: 150%;
    font-weight: 700;
  }

  input {
    margin-bottom: 0px;
    color: rgba(10, 11, 51, 0.6);
    font-size: 16px;
    line-height: 30px;
    background-color: #f3f6fc;
    border: 1px solid #dee9ff;

    @media only screen and (max-width: 1024px) {
      font-size: 16px;
      line-height: 28px;
    }

    &:focus,
    &:active {
      box-shadow: unset;
      outline: unset;
      border-color: #dee9ff;
    }

    &.hide {
      webkit-filter: blur(2px);
      filter: blur(2px);
    }
  }

  .buttoncolumn {
    display: flex;
    justify-content: flex-end;
    margin-top: 25px;
    align-items: center;
  }
`;

export const InputGroupWrapper = styled.div`
  .input-group {
    // border: 2px solid #dee9ff;
    border-radius: 10px;
    // overflow: hidden;

    input {
      height: 55px;
      margin-bottom: 0px;
      color: rgba(10, 11, 51, 0.6);
      font-size: 16px;
      line-height: 30px;
      background-color: #f3f6fc;
      border: 1px solid #dee9ff;

      @media only screen and (max-width: 1024px) {
        font-size: 16px;
        line-height: 28px;
      }

      &:focus,
      &:active {
        box-shadow: unset;
        outline: unset;
        border-color: #dee9ff;
      }

      &.hide {
        webkit-filter: blur(2px);
        filter: blur(2px);
      }
    }
    button {
      border-left: none;
      width: auto;
      height: 55px;
      padding-right: 20px;
      padding-left: 20px;
      border: none;
      border-radius: 0px 10px 10px 0px;
      text-decoration: none;
      color: #ffffff;
      background-color: #f76e5c;
      display: flex;
      justify-content: center;
      align-items: center;
      font-size: 14px;
      font-weight: 700;
      border: 1px solid #f76e5c;

      @media only screen and (max-width: 1024px) {
        font-size: 14px;
        line-height: 24px;
        padding-right: 10px;
        padding-left: 10px;
      }

      &:hover {
        background-color: #ffffff;
        color: #f76e5c;

        img {
          filter: unset;
        }
        
      }
      &:focus, &:active {
        box-shadow: unset;
        outline: unset;
        background-color: #ffffff;
        color: #f76e5c;

        img {
          filter: unset;
        }
        
      }

      img {
        width: 20px;
        margin-right: 10px;
        filter: brightness(0) invert(1);
    }
  }
`;
