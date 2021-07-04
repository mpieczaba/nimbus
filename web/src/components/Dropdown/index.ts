import styled from "styled-components";

import { colors } from "../../themes/colors";

const Dropdown = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
  position: fixed;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  background: rgba(0, 0, 0, 0.4);
`;

export const DropdownItemsWrapper = styled.div`
  display: flex;
  flex-direction: column;
  width: 100%;
  gap: 1rem;
  padding: 1rem;
  background: ${colors.gray};
  font-weight: 600;
  border-radius: 10px 10px 0 0;
  box-shadow: ${colors.boxShadow};
`;

export default Dropdown;
