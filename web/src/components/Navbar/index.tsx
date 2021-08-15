import React, { useState } from "react";
import { Link, useHistory, useLocation } from "react-router-dom";
import {
  IconMenu2,
  IconFiles,
  IconTag,
  IconSettings,
  IconLogout,
} from "@tabler/icons";

import { useAppDispatch } from "../../hooks/store";
import { setToken } from "../../store/actions/authActions";

import { Wrapper, NavButton, SidebarItemLogo, Logo } from "./styles";

import Sidebar, { SidebarItem, SidebarItemIcon } from "../Sidebar";

import Search from "../Search";

const Navbar: React.FC = () => {
  const dispatch = useAppDispatch();
  const history = useHistory();
  const location = useLocation();

  const [sidebar, showHideSidebar] = useState<boolean>(false);

  const handleSignOut = () => {
    dispatch(setToken(""));

    history.push("/");
  };

  return (
    <>
      {location.pathname !== "/login" ? (
        <Wrapper>
          <NavButton onClick={() => showHideSidebar(true)}>
            <IconMenu2 size="2rem" />
          </NavButton>

          <Search />

          <Sidebar active={sidebar} hideSidebar={() => showHideSidebar(false)}>
            <Link to="/">
              <SidebarItemLogo>
                <Logo>Nimbus</Logo>
              </SidebarItemLogo>
            </Link>

            <Link to="/files">
              <SidebarItem>
                <SidebarItemIcon>
                  <IconFiles />
                </SidebarItemIcon>
                Files
              </SidebarItem>
            </Link>

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
          </Sidebar>
        </Wrapper>
      ) : null}
    </>
  );
};

export default Navbar;
