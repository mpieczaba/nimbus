import React from "react";

import { Wrapper, Overlay, DropdownItemsWrapper } from "./styles";

interface Props {
  active: boolean;
  hideDropdown: () => void;
}

const Dropdown: React.FC<Props> = ({ active, hideDropdown, children }) => {
  return (
    <>
      {active ? (
        <Wrapper>
          <DropdownItemsWrapper>{children}</DropdownItemsWrapper>
          <Overlay onClick={hideDropdown} />
        </Wrapper>
      ) : null}
    </>
  );
};

export { DropdownItem, DropdownItemIcon, DropdownItemTitle } from "./styles";
export default Dropdown;
