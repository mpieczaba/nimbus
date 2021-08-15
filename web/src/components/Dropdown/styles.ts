import styled from "styled-components";

import { colors } from "../../themes/colors";

export const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
  position: fixed;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  z-index: 1000;
`;

export const Overlay = styled.div`
  display: flex;
  position: absolute;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  background: ${colors.overlay};
`;

export const DropdownItemsWrapper = styled.div`
  display: flex;
  position: relative;
  flex-direction: column;
  width: 100%;
  gap: 1.5rem;
  padding: 1rem 0;
  background: ${colors.gray};
  font-weight: 600;
  border-radius: 20px 20px 0 0;
  box-shadow: ${colors.boxShadow};
  animation: 100ms ease-out 0s 1 goUp;
  z-index: 1100;
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

export const DropdownItemTitle = styled(DropdownItem)`
  margin: 0;
  padding: 0.5rem 1.5rem 1rem 1.5rem;
  font-size: 1rem;
  font-weight: 600;
  border-bottom: 2px solid ${colors.accent};
`;

export const DropdownItemIcon = styled.div`
  display: flex;
  align-self: center;
  justify-content: center;
  margin-right: 1.5rem;
  color: ${colors.textGray};
`;
