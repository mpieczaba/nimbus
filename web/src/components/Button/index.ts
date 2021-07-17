import styled from "styled-components";

import { colors } from "../../themes/colors";

const Button = styled.button`
  display: flex;
  flex-direction: row;
  align-self: flex-end;
  padding: 0.5rem 1.5rem;
  background: ${colors.accent};
  border-radius: 10px;
  border: none;
  font-weight: 600;
  font-size: 1rem;
  color: ${colors.text};
  box-shadow: ${colors.boxShadow};
  cursor: pointer;
`;

export default Button;
