import styled from "styled-components";

import { colors } from "../../themes/colors";

export const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: 100%;
  background: ${colors.gray};
  border-radius: 10px;
  box-shadow: ${colors.boxShadow};
  cursor: pointer;
`;

export const Thumbnail = styled.div`
  width: 100%;

  div {
    width: 100%;
    height: 100px;
  }

  svg {
    display: flex;
    width: 100%;
    height: 100%;
    padding: 2rem 2rem 1em 2rem;
    filter: grayscale(100%);
  }
`;
