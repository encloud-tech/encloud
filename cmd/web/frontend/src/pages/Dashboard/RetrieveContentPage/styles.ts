import { Button } from "react-bootstrap";
import styled from "styled-components";

export const SectionBox = styled.div`
  color: #697a8d;

  h4 {
    color: #d94964;
  }
`;

export const ColoredBtn = styled(Button)`
  color: #ffffff;
  border: 3px solid #f76e5c !important;
  background-color: #f76e5c;
  padding: 12px 20px;
  border-radius: 47px;
  text-align: center;
  font-weight: bold;
  letter-spacing: 0.5px;

  &:hover {
    color: #f76e5c;
    background: transparent;
  }

  &:focus,
  &:active {
    border: none;
    outline: unset;
    box-shadow: unset !important;
  }

  &:active {
    color: #f76e5c80 !important;
    background-color: unset !important;
  }

  &.nextBtn,
  &.submitBtn,
  &.modalBtn {
    width: 100%;
    max-width: 250px;
  }

  &.loginBtn {
    width: 100%;
    max-width: 250px;
  }

  &.outline {
    background: transparent;
    border: 3px solid #f76e5c !important;
    color: #0a0b33;

    &:hover {
      color: #f76e5c;
    }
  }

  &.step-button {
    display: inline-flex;
    justify-content: center;
    align-items: center;

    > div {
      display: inline-flex;
      justify-content: center;
      align-items: center;
      text-align: center;
      color: #ffffff;
    }

    &:hover {
      color: #f76e5c;
      background: transparent;
    }

    &.loadingStatus {
      color: #ffffff;

      > div > span {
        border-color: #ffffff !important;
        color: #ffffff;

        &.loadingText {
          margin-left: 5px;
        }
      }

      &:hover {
        > div > span {
          border-color: #f76e5c !important;
          color: #f76e5c;
        }
      }
    }
  }
`;
