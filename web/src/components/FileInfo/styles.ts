import styled from "styled-components";

import { colors } from "../../themes/colors";
import { DropdownItem } from "../Dropdown";

export const Wrapper = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  gap: 0.5rem;
  width: 100%;
  padding: 0.5rem;
`;

export const Thumbnail = styled.div<{ rich?: boolean }>`
  align-self: center;

  div {
    width: ${(props) => (props.rich ? "2.5rem" : "1.5rem")};
    height: ${(props) => (props.rich ? "2.5rem" : "1.5rem")};
  }
`;

export const FileInfoWrapper = styled.div`
  display: flex;
  flex-direction: column;
  flex: 1;
  min-width: 0;
  width: 100%;
  align-self: center;
`;

export const FileName = styled.div<{ rich?: boolean }>`
  display: inline-block;
  width: 100%;
  line-height: 1.5rem;
  font-size: 0.8rem;
  font-weight: 600;
  color: ${colors.text};
  text-align: ${(props) => (props.rich ? "flex-start" : "center")};
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
`;

export const FileInfoElements = styled.div`
  display: flex;
  flex-direction: row;
  gap: 1rem;
`;

export const FileInfoElement = styled.div`
  display: flex;
  align-self: flex-start;
  font-size: 0.7rem;
  color: ${colors.textGray};
`;

export const FileMenuButton = styled.button`
  display: flex;
  align-self: center;
  background: none;
  border: none;
  cursor: pointer;
  color: ${colors.textGray};
  font-size: 0.8rem;
`;

export const DropdownItemName = styled(DropdownItem)`
  margin: 0;
  padding: 0.5rem 1.5rem 1rem 1.5rem;
  font-size: 1rem;
  font-weight: 600;
  border-bottom: 2px solid ${colors.accent};

  div {
    width: 1.25rem;
    height: 1.25rem;
    margin-right: 1.5rem;
  }
`;

export const DropdownItemDelete = styled(DropdownItem)`
  color: ${colors.error};
`;
