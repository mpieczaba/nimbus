import React, { useState } from "react";
import dayjs from "dayjs";
import fileSize from "filesize";
import {
  IconArrowDown,
  IconArrowUp,
  IconLayoutGrid,
  IconLayoutList,
} from "@tabler/icons";

import { useFilesQuery } from "../../generated/graphql";

import Dropdown, {
  DropdownItemsWrapper,
  DropdownItem,
  DropdownItemTitle,
} from "../Dropdown";
import FileCard from "../FileCard";
import FileList from "../FileList";

import {
  Wrapper,
  OptionsWrapper,
  FilterOption,
  FilterButton,
  ViewOption,
  FileCardWrapper,
  FileListWrapper,
} from "./styles";

const Files: React.FC = () => {
  const { loading, error, data } = useFilesQuery();

  const [dropdown, showHideDropdown] = useState<boolean>(false);
  const [order, setOrder] = useState<boolean>(true);
  const [view, setView] = useState<boolean>(true);

  const handleDropdownShowHide = () => {
    showHideDropdown(!dropdown);
  };

  const handleViewOptionChange = () => {
    setView(!view);
  };

  const handleOrderChange = () => {
    setOrder(!order);
  };

  if (loading) return <>Loading...</>;
  if (error || !data) return <>Error!</>;

  return (
    <Wrapper>
      <OptionsWrapper>
        <FilterOption>
          <div onClick={handleDropdownShowHide}>Name</div>
          {dropdown ? (
            <Dropdown onClick={handleDropdownShowHide}>
              <DropdownItemsWrapper>
                <DropdownItemTitle>Sort by</DropdownItemTitle>
                <DropdownItem>Name</DropdownItem>
                <DropdownItem>Modification date</DropdownItem>
                <DropdownItem>Size</DropdownItem>
              </DropdownItemsWrapper>
            </Dropdown>
          ) : null}
          <FilterButton onClick={handleOrderChange}>
            {order ? <IconArrowDown /> : <IconArrowUp />}
          </FilterButton>
        </FilterOption>
        <ViewOption onClick={handleViewOptionChange}>
          {view ? <IconLayoutList /> : <IconLayoutGrid />}
        </ViewOption>
      </OptionsWrapper>

      {view ? (
        <FileCardWrapper>
          {data.files?.edges?.map((edge, index) => (
            <FileCard
              key={index}
              file={{
                id: edge!.node.id,
                name: edge!.node.name,
                url: edge!.node.url,
              }}
              thumbnail={edge!.node.url}
            />
          ))}
        </FileCardWrapper>
      ) : (
        <FileListWrapper>
          {data.files?.edges?.map((edge, index) => (
            <FileList
              key={index}
              file={{
                id: edge!.node.id,
                name: edge!.node.name,
                modificationDate: dayjs(edge!.node.updatedAt)
                  .format("D MMM YYYY H:m")
                  .toString(),
                size: fileSize(edge!.node.size).toString(),
                url: edge!.node.url,
              }}
              thumbnail={edge!.node.url}
            />
          ))}
        </FileListWrapper>
      )}
    </Wrapper>
  );
};

export default Files;
