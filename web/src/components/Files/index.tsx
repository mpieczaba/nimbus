import React, { useState } from "react";
import { useLocation } from "react-router-dom";
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
import InfiniteScroll from "react-infinite-scroll-component";
import { extractSearchOperators } from "../../utils/search";

const Files: React.FC = () => {
  const location = useLocation();

  const urlSearchParams = new URLSearchParams(location.search);

  const { loading, error, data, fetchMore } = useFilesQuery({
    variables: {
      first: 20,
      name: extractSearchOperators(urlSearchParams.get("search")),
      tags: urlSearchParams.getAll("tag").map((tag) => `#${tag}`),
    },
  });

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
                tags: urlSearchParams.getAll("tag").map((tag) => `#${tag}`),
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
      </InfiniteScroll>
    </Wrapper>
  );
};

export default Files;
