import styled from "styled-components";

import { colors } from "../../themes/colors";

const Form = styled.form`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: 80%;
  padding: 2rem 1rem;
  background: ${colors.gray};
  border-radius: 10px;
  box-shadow: ${colors.boxShadow};
`;

export const Header = styled.h2`
  display: flex;
  justify-content: center;
  margin: 0;
  padding-bottom: 1rem;
  color: ${colors.text};
  font-family: "Baloo Tammudu 2", sans-serif;
  font-weight: 700;
  line-height: 2rem;
  font-size: 2rem;
`;

export default Form;
