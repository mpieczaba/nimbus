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
  gap: 1.5rem;
  padding: 1rem 0;
  background: ${colors.gray};
  font-weight: 600;
  border-radius: 20px 20px 0 0;
  box-shadow: ${colors.boxShadow};
  animation: 100ms ease-out 0s 1 goUp;

  @keyframes goUp {
    0% {
      transform: translateY(100px);
    }

    100% {
      transform: translateY(0);
    }
  }
`;

export const DropdownItem = styled.div`
  display: flex;
  flex-direction: row;
  align-items: center;
  padding: 0 0.5rem;
  margin: 0 1rem;
  font-size: 1rem;
  color: ${colors.text};
  font-weight: 400;
  text-align: center;
`;

export const DropdownItemIcon = styled.div`
  display: flex;
  align-self: center;
  justify-content: center;
  margin-right: 1.5rem;
  color: ${colors.textGray};
`;

export default Dropdown;
