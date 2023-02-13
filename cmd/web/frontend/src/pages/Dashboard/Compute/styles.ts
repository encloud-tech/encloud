import styled from "styled-components";

export const SectionBox = styled.div`
  color: #697a8d;

  td {
    white-space: normal !important;
    word-wrap: break-word;
  }

  table {
    table-layout: fixed;
  }
`;

export const KeyBoxedContent = styled.div`
  width: 100%;

  &.active {
    background: #d94964;
    border: 2px solid #ffffff50;
  }
`;

export const KeyPairsSection = styled.div`
  color: #cc3366;

  h3 {
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
