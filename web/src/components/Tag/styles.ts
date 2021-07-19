import styled from "styled-components";

import { colors } from "../../themes/colors";

export const Wrapper = styled.div`
  display: flex;
  flex-direction: row;
  gap: 0.5rem;
  padding: 0 0.5rem;
  font-size: 0.8rem;
  line-height: 1.5rem;
  background: ${colors.accent};
  color: ${colors.text};
  border-radius: 5px;
`;

export const TagName = styled.span`
  font-weight: 600;
`;

export const Button = styled.button`
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0;
  background: none;
  border: none;
  cursor: pointer;
`;
