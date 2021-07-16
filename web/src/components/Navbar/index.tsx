import React, { useState } from "react";
import { Link, useLocation } from "react-router-dom";
import {
  IconMenu2,
  IconSearch,
  IconFiles,
  IconTag,
  IconSettings,
  IconLogout,
} from "@tabler/icons";

import {
  Wrapper,
  NavButton,
  SearchWrapper,
  SearchIcon,
  SearchInput,
  SidebarItemLogo,
  Logo,
} from "./styles";

import Sidebar, {
  SidebarItemsWrapper,
  SidebarItem,
  SidebarItemIcon,
} from "../Sidebar";

const Navbar: React.FC = () => {
  const location = useLocation();

  const [sidebar, showHideSidebar] = useState<boolean>(false);

  const handleSidebarShowHide = () => {
    showHideSidebar(!sidebar);
  };

  return (
    <>
      {location.pathname !== "/login" ? (
        <Wrapper>
          <NavButton onClick={handleSidebarShowHide}>
            <IconMenu2 size="2rem" />
          </NavButton>

          <SearchWrapper>
            <SearchIcon>
              <IconSearch />
            </SearchIcon>
            <SearchInput type="text" placeholder="Search..." />
          </SearchWrapper>

          {sidebar ? (
            <Sidebar onClick={handleSidebarShowHide}>
              <SidebarItemsWrapper>
                <Link to="/">
                  <SidebarItemLogo>
                    <Logo>Nimbus</Logo>
                  </SidebarItemLogo>
                </Link>

                <SidebarItem>
                  <SidebarItemIcon>
                    <IconFiles />
                  </SidebarItemIcon>
                  Files
                </SidebarItem>

                <SidebarItem>
                  <SidebarItemIcon>
                    <IconTag />
                  </SidebarItemIcon>
                  Tags
                </SidebarItem>

                <SidebarItem>
                  <SidebarItemIcon>
                    <IconSettings />
                  </SidebarItemIcon>
                  Settings
                </SidebarItem>

                <SidebarItem>
                  <SidebarItemIcon>
                    <IconLogout />
                  </SidebarItemIcon>
                  Logout
                </SidebarItem>
              </SidebarItemsWrapper>
            </Sidebar>
          ) : null}
        </Wrapper>
      ) : null}
    </>
  );
};

export default Navbar;
