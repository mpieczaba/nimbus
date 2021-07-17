import React from "react";
import { Link } from "react-router-dom";

import { Wrapper } from "./styles";

import FileInfo from "../FileInfo";

interface Props {
  file: {
    id: string;
    name: string;
    size?: string;
    url: string;
    modificationDate?: string;
  };
  thumbnail?: string;
}

const FileList: React.FC<Props> = ({ file, thumbnail }) => {
  return (
    <Link to={`/files/${file.id}`}>
      <Wrapper>
        <FileInfo rich file={file} thumbnail={thumbnail} />
      </Wrapper>
    </Link>
  );
};

export default FileList;
