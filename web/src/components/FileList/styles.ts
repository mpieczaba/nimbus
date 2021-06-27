import styled from "styled-components";

import { colors } from "../../themes/colors";

export const Wrapper = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  background: ${colors.gray};
  border-radius: 10px;
  box-shadow: ${colors.boxShadow};
  cursor: pointer;
`;
