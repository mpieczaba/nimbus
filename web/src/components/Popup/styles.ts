import styled from "styled-components";

import { colors } from "../../themes/colors";

export const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
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

export const PopupItemsWrapper = styled.div`
  display: flex;
  position: relative;
  flex-direction: column;
  gap: 1.5rem;
  margin: 1rem;
  padding: 1rem 0;
  background: ${colors.gray};
  font-weight: 600;
  border-radius: 20px;
  z-index: 1100;
`;

export const PopupItem = styled.div<{ right?: boolean }>`
  display: flex;
  flex-direction: row;
  justify-content: ${(props) => (props.right ? "flex-end" : "flex-start")};
  align-items: center;
  padding: 0 0.5rem;
  margin: 0 1rem;
  font-size: 1rem;
  color: ${colors.text};
  font-weight: 400;
  text-align: center;
`;

export const PopupItemTitle = styled(PopupItem)`
  margin: 0;
  padding: 0.5rem 1.5rem 1rem 1.5rem;
  font-size: 1rem;
  font-weight: 600;
  border-bottom: 2px solid ${colors.accent};
`;

export const PopupItemIcon = styled.div`
  display: flex;
  align-self: center;
  justify-content: center;
  margin-right: 1.5rem;
  color: ${colors.textGray};
`;