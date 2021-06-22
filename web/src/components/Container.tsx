import styled from "styled-components";

import colors from "../utils/colors";

export const Container = styled.main`
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100%;
  min-height: 100vh;
  background: ${colors.background};
`;
