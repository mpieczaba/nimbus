import styled from "styled-components";

import { colors } from "../../themes/colors";

export const Wrapper = styled.div`
  display: flex;
  flex-direction: row;
  gap: 0.5rem;
  align-items: center;
  width: 100%;
  height: 2.5rem;
  padding: 0 0.5rem;
  background: ${colors.gray};
  border-radius: 10px;
  overflow-x: auto;
  overflow-y: hidden;
`;

export const SearchIcon = styled.div`
  display: flex;
`;

export const SearchInput = styled.input`
  display: flex;
  width: 100%;
  min-width: 60vw;
  background: none;
  border: none;
  outline: none;
  font-size: 1rem;
  color: ${colors.text};
  z-index: 500;
`;

export const SearchResultsWrapper = styled.div`
  display: flex;
  flex-direction: column;
  position: absolute;
  top: 3rem;
  left: 4rem;
  right: 1rem;
  z-index: 100;
`;

export const SearchResults = styled.div`
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  gap: 0.5rem;
  margin-left: 2px;
  padding: 1rem;
  background: ${colors.gray};
  border-radius: 10px;
  border: 2px solid ${colors.accent};
`;
