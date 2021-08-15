import styled, { css } from "styled-components";

import { colors } from "./colors";

export const style = css`
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

  a {
    outline: none;
    text-decoration: none;
  }

  @keyframes goUp {
    0% {
      transform: translateY(100px);
    }

    100% {
      transform: translateY(0);
    }
  }

  @keyframes goRight {
    0% {
      transform: translateX(-100px);
    }

    100% {
      transform: translateX(0);
    }
  }
`;

export const AppWrapper = styled.div`
  display: flex;
  width: 100%;
  height: 100vh;
  position: absolute;
  flex-direction: column;
`;
