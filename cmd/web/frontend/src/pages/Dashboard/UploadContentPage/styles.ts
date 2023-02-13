import styled from "styled-components";
import { ProgressBar } from "react-bootstrap";

export const DragArea = styled.div`
  border: 2px dashed #d94964;
  padding: 3em;
  width: 100%;
  border-radius: 5px;
  position: relative;
  display: flex;
  align-items: center;
  text-align: center;
  transition: 0.2s;
  margin-bottom: 30px;

  .choose-file-button {
    flex-shrink: 0;
    background-color: rgba(255, 255, 255, 0.5);
    border: 1px solid rgba(255, 255, 255, 1);
    border-radius: 3px;
    padding: 8px 15px;
    margin-right: 10px;
    font-size: 12px;
    text-transform: uppercase;
  }

  .file-message {
    font-weight: 400;
    line-height: 1.4;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .file-input {
    position: absolute;
    left: 0;
    top: 0;
    height: 100%;
    width: 100%;
    cursor: pointer;
    opacity: 0;
  }
`;

export const FileUploadBar = styled(ProgressBar)`
  background-color: #f76e5c50;
  margin-bottom: 30px;
  .progress-bar {
    background-color: #f76e5c;
  }
`;
