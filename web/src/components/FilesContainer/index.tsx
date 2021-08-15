import React, { useState } from "react";
import { useLocation } from "react-router-dom";
import dayjs from "dayjs";
import fileSize from "filesize";
import InfiniteScroll from "react-infinite-scroll-component";
import {
  IconArrowDown,
  IconArrowUp,
  IconLayoutGrid,
  IconLayoutList,
} from "@tabler/icons";

import { useFilesQuery } from "../../generated/graphql";

import Dropdown, { DropdownItem, DropdownItemTitle } from "../Dropdown";
import FileCard from "../FileCard";
import FileList from "../FileList";
import Tag from "../Tag";

import {
  Wrapper,
  OptionsWrapper,
  FilterOption,
  FilterButton,
  ViewOption,
  FileCardWrapper,
  FileListWrapper,
  Tags,
} from "./styles";

import { extractSearchOperators } from "../../utils/search";

const FilesContainer: React.FC = () => {
  const location = useLocation();

  const urlSearchParams = new URLSearchParams(location.search);

  const { loading, error, data, fetchMore } = useFilesQuery({
    variables: {
      first: 20,
      name: extractSearchOperators(urlSearchParams.get("search")),
      tags: urlSearchParams.getAll("tag"),
    },
  });

  const [dropdown, showHideDropdown] = useState<boolean>(false);
  const [order, setOrder] = useState<boolean>(true);
  const [view, setView] = useState<boolean>(true);

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
          <div onClick={() => showHideDropdown(true)}>Name</div>

          <Dropdown
            active={dropdown}
            hideDropdown={() => showHideDropdown(false)}
          >
            <DropdownItemTitle>Sort by</DropdownItemTitle>
            <DropdownItem>Name</DropdownItem>
            <DropdownItem>Modification date</DropdownItem>
            <DropdownItem>Size</DropdownItem>
          </Dropdown>

          <FilterButton onClick={handleOrderChange}>
            {order ? <IconArrowDown /> : <IconArrowUp />}
          </FilterButton>
        </FilterOption>
        <ViewOption onClick={handleViewOptionChange}>
          {view ? <IconLayoutList /> : <IconLayoutGrid />}
        </ViewOption>
      </OptionsWrapper>

      <Tags>
        {urlSearchParams.getAll("tag").map((tag, index) => (
          <Tag key={index} tagName={`#${tag}`} />
        ))}
      </Tags>

      <InfiniteScroll
        next={async () => {
          if (data?.files?.pageInfo.hasNextPage) {
            await fetchMore({
              variables: {
                after: data?.files.pageInfo.endCursor,
                first: 20,
                name: urlSearchParams.get("search"),
                tags: urlSearchParams.getAll("tag"),
              },
            });
          }
        }}
        hasMore={!!data.files?.pageInfo.hasNextPage}
        loader={<>Loading...</>}
        dataLength={data.files?.edges?.length || 0}
      >
        {view ? (
          <FileCardWrapper>
            {data.files?.edges?.map((edge, index) => (
              <FileCard key={index} file={edge!.node} />
            ))}
          </FileCardWrapper>
        ) : (
          <FileListWrapper>
            {data.files?.edges?.map((edge, index) => (
              <FileList
                key={index}
                file={{
                  ...edge!.node,
                  modificationDate: dayjs(edge!.node.updatedAt)
                    .format("D MMM YYYY H:m")
                    .toString(),
                  size: fileSize(edge!.node.size).toString(),
                }}
              />
            ))}
          </FileListWrapper>
        )}
      </InfiniteScroll>
    </Wrapper>
  );
};

export default FilesContainer;
