import React from "react";

import { Wrapper, SidebarItemsWrapper, Overlay } from "./styles";

interface Props {
  active: boolean;
  hideSidebar: () => void;
}

const Sidebar: React.FC<Props> = ({ active, hideSidebar, children }) => {
  return (
    <>
      {active ? (
        <Wrapper>
          <SidebarItemsWrapper>{children}</SidebarItemsWrapper>
          <Overlay onClick={hideSidebar} />
        </Wrapper>
      ) : null}
    </>
  );
};

export { SidebarItem, SidebarItemIcon } from "./styles";
export default Sidebar;
