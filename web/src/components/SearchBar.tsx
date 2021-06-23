import React from "react";
import styled from "styled-components";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

import colors from "../utils/colors";

import { Row } from "./Row";

const Wrapper = styled(Row)`
  top: 0;
  position: sticky;
  background: ${colors.background};
  border-bottom: 3px solid ${colors.accent};
  z-index: 1000;
`;

const SearchMenuButton = styled.button`
  display: flex;
  align-self: center;
  margin: 1rem 0 1rem 1rem;
  padding: 0;
  background: none;
  font-size: 1.5rem;
  color: ${colors.text};
  border: none;
`;

const SearchInputWrapper = styled(Row)`
  height: 2.5rem;
  width: 100%;
  margin: 0.5rem 1rem;
  background: ${colors.gray};
  border-radius: 10px;
  box-shadow: ${colors.boxShadow};
`;

const SearchIcon = styled.div`
  display: flex;
  align-self: center;
  padding: 0 1rem;
  font-size: 1rem;
`;

const SearchInput = styled.input`
  display: flex;
  width: 100%;
  padding-right: 1rem;
  background: none;
  border: none;
  color: ${colors.text};
  font-size: 1rem;
  outline: none;
`;

export const SearchBar: React.FC = () => {
  return (
    <Wrapper>
      <SearchMenuButton>
        <FontAwesomeIcon icon="bars" />
      </SearchMenuButton>
      <SearchInputWrapper>
        <SearchIcon>
          <FontAwesomeIcon icon="search" />
        </SearchIcon>
        <SearchInput type="text" />
      </SearchInputWrapper>
    </Wrapper>
  );
};
