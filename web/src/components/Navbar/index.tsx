import React, { useState } from "react";
import { useDispatch } from "react-redux";
import { Link, useHistory, useLocation } from "react-router-dom";
import {
  IconMenu2,
  IconFiles,
  IconTag,
  IconSettings,
  IconLogout,
} from "@tabler/icons";

import { setToken } from "../../actions/authActions";

import { Wrapper, NavButton, SidebarItemLogo, Logo } from "./styles";

import Sidebar, {
  SidebarItemsWrapper,
  SidebarItem,
  SidebarItemIcon,
} from "../Sidebar";

import Search from "../Search";

const Navbar: React.FC = () => {
  const dispatch = useDispatch();
  const history = useHistory();
  const location = useLocation();

  const [sidebar, showHideSidebar] = useState<boolean>(false);

  const handleSidebarShowHide = () => {
    showHideSidebar(!sidebar);
  };

  const handleSignOut = () => {
    dispatch(setToken(""));

    history.push("/");
  };

  return (
    <>
      {location.pathname !== "/login" ? (
        <Wrapper>
          <NavButton onClick={handleSidebarShowHide}>
            <IconMenu2 size="2rem" />
          </NavButton>

          <Search />

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

                <SidebarItem onClick={handleSignOut}>
                  <SidebarItemIcon>
                    <IconLogout />
                  </SidebarItemIcon>
                  Sign out
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
