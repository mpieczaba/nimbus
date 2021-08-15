import styled from "styled-components";

import { colors } from "../../themes/colors";

import { PopupItemTitle } from "../Popup/styles";

export const Wrapper = styled.div`
  display: flex;
  position: fixed;
  align-items: center;
  justify-content: center;
  width: 3.5rem;
  height: 3.5rem;
  margin: 1rem;
  right: 0;
  bottom: 0;
  background: ${colors.accent};
  color: ${colors.text};
  border: none;
  border-radius: 50px;
  box-shadow: ${colors.boxShadow};
  z-index: 600;
`;

export const FileInputLabel = styled.label`
  display: flex;
  align-items: center;
  justify-content: center;
`;

export const Thumbnail = styled.div`
  align-self: center;

  div {
    width: 1.5rem;
    height: 1.5rem;
  }
`;

export const PopupItemName = styled(PopupItemTitle)`
  div {
    width: 1.25rem;
    height: 1.25rem;
    margin-right: 1.5rem;
  }

  span {
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }
`;
