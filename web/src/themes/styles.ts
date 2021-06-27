import styled, { css } from "styled-components";

import { colors } from "./colors";

export const style = css`
  @import url("https://fonts.googleapis.com/css2?family=Open+Sans:wght@300;400;600;700&display=swap");

  *,
  *::before,
  *::after {
    box-sizing: border-box;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }

  body {
    font-family: "Open Sans", sans-serif;
    margin: 0;
    padding: 0;
    width: 100%;
    height: 100vh;
    background: ${colors.background};
    color: ${colors.text};
  }
`;

export const AppWrapper = styled.div`
  display: flex;
  width: 100%;
  height: 100vh;
  position: absolute;
  flex-direction: column;
`;
