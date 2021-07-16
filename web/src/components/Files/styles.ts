import styled from "styled-components";

import { colors } from "../../themes/colors";

export const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  margin-top: 3.5rem;
  padding: 1rem;
`;

export const OptionsWrapper = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: space-between;
`;

export const FilterOption = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  padding: 0.5rem 0.75rem;
  font-weight: 600;
  color: ${colors.textGray};
  cursor: pointer;
`;

export const FilterButton = styled.button`
  display: flex;
  align-self: center;
  justify-content: center;
  padding: 0;
  color: ${colors.textGray};
  background: none;
  border: none;
  border-radius: 10px;
  cursor: pointer;
`;

export const ViewOption = styled.button`
  display: flex;
  align-self: center;
  justify-content: center;
  padding: 0.5rem;
  background: none;
  color: ${colors.textGray};
  font-size: 1.25rem;
  border: none;
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
