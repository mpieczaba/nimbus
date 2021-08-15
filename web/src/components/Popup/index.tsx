import React from "react";

import { Overlay, Wrapper, PopupItemsWrapper } from "./styles";

interface Props {
  active: boolean;
  hidePopup: () => void;
}

const Popup: React.FC<Props> = ({ active, hidePopup, children }) => {
  return (
    <>
      {active ? (
        <Wrapper>
          <PopupItemsWrapper>{children}</PopupItemsWrapper>
          <Overlay onClick={hidePopup} />
        </Wrapper>
      ) : null}
    </>
  );
};

export { PopupItem, PopupItemIcon, PopupItemTitle } from "./styles";
export default Popup;
