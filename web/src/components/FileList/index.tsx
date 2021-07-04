import React from "react";

import { Wrapper } from "./styles";

import FileInfo from "../FileInfo";

interface Props {
  file: {
    name: string;
    size?: string;
    url: string;
    modificationDate?: string;
  };
  thumbnail?: string;
}

const FileList: React.FC<Props> = ({ file, thumbnail }) => {
  return (
    <Wrapper>
      <FileInfo rich file={file} thumbnail={thumbnail} />
    </Wrapper>
  );
};

export default FileList;
