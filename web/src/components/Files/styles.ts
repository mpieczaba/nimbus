import styled from "styled-components";

import { colors } from "../../themes/colors";

export const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  padding: 1rem;
`;

export const OptionsWrapper = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: space-between;
`;

export const ViewOptionsWrapper = styled.div`
  display: flex;
  flex-direction: row;
  background: ${colors.gray};
  border-radius: 10px;
  box-shadow: ${colors.boxShadow};
`;

interface ViewOptionProps {
  active?: boolean;
}

export const ViewOption = styled.button<ViewOptionProps>`
  display: flex;
  align-self: center;
  justify-content: center;
  width: 2.25rem;
  height: 2.25rem;
  padding: 0.5rem;
  background: ${(props) => (props.active ? colors.accent : colors.gray)};
  color: ${colors.text};
  font-size: 1.25rem;
  border: none;
  border-radius: 10px;
  cursor: pointer;
`;

export const FileCardWrapper = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(150px, 2fr));
  column-gap: 1rem;
  grid-template-rows: 1fr auto;
  row-gap: 1rem;
  width: 100%;
  padding: 1rem 0;
`;

export const FileListWrapper = styled.div`
  display: flex;
  flex-direction: column;
  gap: 1rem;
  padding: 1rem 0;
`;
