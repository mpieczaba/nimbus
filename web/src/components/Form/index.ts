import styled from "styled-components";

import { colors } from "../../themes/colors";

const Form = styled.form`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 2rem 1rem;
  background: ${colors.gray};
  border-radius: 10px;
  box-shadow: ${colors.boxShadow};
`;

export default Form;
