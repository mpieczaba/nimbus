import styled from "styled-components";

import { colors } from "../../themes/colors";

const Input = styled.input`
  display: flex;
  width: 100%;
  padding-bottom: 0.5rem;
  background: none;
  border: none;
  outline: none;
  font-size: 1rem;
  color: ${colors.text};
  border-bottom: 2px solid ${colors.accent};
`;

export const InputWrapper = styled.div`
  display: flex;
  flex-direction: column;
  width: 100%;
`;

export const InputAndIconWrapper = styled.div`
  display: flex;
  flex-direction: row;
`;

export const InputIcon = styled.div`
  display: flex;
  align-self: center;
  justify-content: center;
  margin-right: 1.5rem;
  color: ${colors.textGray};
`;

export const InputError = styled.div`
  display: flex;
  height: 1rem;
  margin: 0 0 1rem 3rem;
  color: ${colors.error};
`;

export default Input;
