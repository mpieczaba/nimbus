import styled from "styled-components";

import { colors } from "../../themes/colors";

const Sidebar = styled.div`
  display: flex;
  flex-direction: column;
  position: fixed;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  background: ${colors.overlay};
`;

export const SidebarItemsWrapper = styled.div`
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  width: 80%;
  height: 100%;
  background: ${colors.background};
  box-shadow: ${colors.boxShadow};
  animation: 100ms ease-out 0s 1 goRight;
`;

export const SidebarItem = styled.div`
  display: flex;
  flex-direction: row;
  align-items: center;
  padding: 0 1rem;
  font-size: 1rem;
  color: ${colors.text};
  font-weight: 400;
  text-align: center;
`;

export const SidebarItemIcon = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  height: 2.5rem;
  width: 2.5rem;
  margin-right: 1rem;
  color: ${colors.textGray};
`;

export default Sidebar;
