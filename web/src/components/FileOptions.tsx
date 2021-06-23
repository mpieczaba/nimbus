import React from "react";
import styled from "styled-components";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

import colors from "../utils/colors";

import { Row } from "./Row";

const Wrapper = styled(Row)`
  justify-content: space-between;
  margin: 0 1rem;
  padding-top: 1rem;
`;

const Filter = styled(Row)`
  gap: 1rem;
  height: 2.25rem;
  padding: 1rem;
  background: ${colors.gray};
  border-radius: 10px;
`;

const FilterOption = styled(Row)`
  align-self: center;
  font-weight: 600;
`;

const FilerIcon = styled(Row)`
  align-self: center;
`;

const ViewOptions = styled(Row)`
  background: ${colors.gray};
  border-radius: 10px;
  box-shadow: ${colors.boxShadow};
`;

const ViewOption = styled.button`
  display: flex;
  align-self: center;
  justify-content: center;
  width: 2.25rem;
  height: 2.25rem;
  padding: 0.5rem;
  background: ${colors.gray};
  color: ${colors.text};
  font-size: 1.25rem;
  border: none;
  border-radius: 10px;

  &.active {
    background: ${colors.accent};
  }
`;

export const FileOptions: React.FC = () => {
  return (
    <Wrapper>
      <Filter>
        <FilterOption>Name</FilterOption>
        <FilerIcon>
          <FontAwesomeIcon icon="arrow-down" />
        </FilerIcon>
      </Filter>
      <ViewOptions>
        <ViewOption className="active">
          <FontAwesomeIcon icon="grip-vertical" />
        </ViewOption>
        <ViewOption>
          <FontAwesomeIcon icon="th-list" />
        </ViewOption>
      </ViewOptions>
    </Wrapper>
  );
};
