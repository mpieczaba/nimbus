import React, { ChangeEvent, FormEvent, useState } from "react";
import { useHistory } from "react-router-dom";
import { IconSearch } from "@tabler/icons";

import { useTagsLazyQuery } from "../../generated/graphql";

import {
  Wrapper,
  SearchIcon,
  SearchInput,
  SearchResultsWrapper,
  SearchResults,
} from "./styles";

import Tag from "../Tag";

const Search: React.FC = () => {
  const history = useHistory();

  const [getTags, { data }] = useTagsLazyQuery();

  const [tags, setTags] = useState<Array<string>>(Array<string>());
  const [searchInput, setSearchInput] = useState<string>("");
  const [searchResults, showHideSearchResults] = useState<boolean>(false);

  const handleSearchSubmit = (e: FormEvent) => {
    e.preventDefault();

    const urlSearchParams = new URLSearchParams();

    if (searchInput.length > 0) urlSearchParams.set("search", searchInput);

    tags.map((tag) => urlSearchParams.append("tag", tag));

    history.push({ pathname: "/files", search: urlSearchParams.toString() });
  };

  const handleSearchInputChange = (e: ChangeEvent<HTMLInputElement>) => {
    setSearchInput(e.target.value);
    showHideSearchResults(false);

    if (e.target.value.startsWith("#")) {
      showHideSearchResults(true);

      getTags({
        variables: {
          name: `${e.target.value.substring(1)}%`,
          first: 10,
        },
      });
    }
  };

  const handleTagClick = (tagName: string) => {
    setTags(tags?.concat(tagName));
    setSearchInput("");
    showHideSearchResults(false);
  };

  const handleTagRemove = (index: number) => {
    setTags(tags.filter((item, j) => index !== j));
  };

  return (
    <Wrapper>
      <SearchIcon>
        <IconSearch />
      </SearchIcon>

      {tags?.map((name, index) => (
        <Tag
          key={index}
          removable
          tagName={`#${name}`}
          handleTagRemove={() => handleTagRemove(index)}
        />
      ))}

      <form onSubmit={handleSearchSubmit}>
        <SearchInput
          type="text"
          name="search"
          placeholder="Search for files..."
          value={searchInput}
          onChange={handleSearchInputChange}
        />
      </form>

      <SearchResultsWrapper>
        {searchResults && data?.tags?.edges ? (
          <SearchResults>
            {data?.tags?.edges?.map((edge, index) => (
              <Tag
                key={index}
                tagName={`#${edge!.node.name}`}
                onClick={() => handleTagClick(edge!.node.name)}
              />
            ))}
          </SearchResults>
        ) : null}
      </SearchResultsWrapper>
    </Wrapper>
  );
};

export default Search;
