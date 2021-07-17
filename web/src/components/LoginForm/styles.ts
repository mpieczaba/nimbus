import styled from "styled-components";

import { colors } from "../../themes/colors";

export const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
`;

export const Logo = styled.h1`
  display: flex;
  margin: 0;
  padding: 1rem;
  font-family: "Baloo Tammudu 2", sans-serif;
  color: ${colors.accent};
  text-transform: lowercase;
  font-size: 2.5rem;
  font-weight: 700;
  line-height: 2rem;
`;

export const Error = styled.div`
  display: flex;
  justify-content: center;
  margin-bottom: 1rem;
  height: 2rem;
  color: ${colors.error};
`;
