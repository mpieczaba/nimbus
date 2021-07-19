import styled from "styled-components";

import { colors } from "../../themes/colors";
import { SidebarItem } from "../Sidebar";

export const Wrapper = styled.div`
  display: flex;
  flex-direction: row;
  gap: 1rem;
  align-items: center;
  justify-content: space-between;
  padding: 0.5rem 1rem;
  position: fixed;
  width: 100%;
  background: ${colors.background};
  border-bottom: 2px solid ${colors.accent};
`;

export const NavButton = styled.button`
  display: flex;
  align-items: center;
  justify-content: center;
  height: 2.5rem;
  width: 2.5rem;
  padding: 0;
  background: none;
  border: none;
  color: ${colors.text};
`;

export const SidebarItemLogo = styled(SidebarItem)`
  padding: 0.5rem 1rem;
  border-bottom: 2px solid ${colors.accent};
`;

export const Logo = styled.span`
  display: flex;
  height: 2.5rem;
  margin-left: 0.5rem;
  font-size: 2rem;
  color: ${colors.accent};
  font-weight: 700;
  text-transform: lowercase;
  text-align: center;
  line-height: 3.5rem;
  font-family: "Baloo Tammudu 2", sans-serif;
`;
