import React, { useState } from "react";
import { useQuery } from "@apollo/react-hooks";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import dayjs from "dayjs";
import fileSize from "filesize";

import files, { FilesData } from "../../apollo/queries/files";

import FileCard from "../FileCard";
import FileList from "../FileList";

import {
  Wrapper,
  OptionsWrapper,
  ViewOptionsWrapper,
  ViewOption,
  FileCardWrapper,
  FileListWrapper,
} from "./styles";

const Files = () => {
  const { loading, error, data } = useQuery<FilesData>(files);

  const [view, setView] = useState<boolean>(true);

  const handleOptionChange = () => {
    setView(!view);
  };

  if (loading) {
    return <>Loading...</>;
  }
  if (error) return <>Error! {error.message}</>;

  return (
    <Wrapper>
      <OptionsWrapper>
        <div />
        <ViewOptionsWrapper>
          <ViewOption onClick={handleOptionChange} active={view}>
            <FontAwesomeIcon icon="grip-vertical" />
          </ViewOption>
          <ViewOption onClick={handleOptionChange} active={!view}>
            <FontAwesomeIcon icon="th-list" />
          </ViewOption>
        </ViewOptionsWrapper>
      </OptionsWrapper>
      {view ? (
        <FileCardWrapper>
          {data &&
            data.files.edges.map((el) => (
              <FileCard fileName={el.node.name} thumbnail={el.node.url} />
            ))}
        </FileCardWrapper>
      ) : (
        <FileListWrapper>
          {data &&
            data.files.edges.map((el) => (
              <FileList
                file={{
                  name: el.node.name,
                  modificationDate: dayjs(el.node.updatedAt)
                    .format("D MMMM YYYY H:m")
                    .toString(),
                  size: fileSize(el.node.size).toString(),
                }}
                thumbnail={el.node.url}
              />
            ))}
        </FileListWrapper>
      )}
    </Wrapper>
  );
};

export default Files;
