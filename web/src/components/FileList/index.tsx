import React from "react";
import { Link } from "react-router-dom";

import { Wrapper } from "./styles";

import FileInfo from "../FileInfo";

interface Props {
  file: {
    id: string;
    name: string;
    extension: string;
    size?: string;
    url: string;
    downloadURL: string;
    modificationDate?: string;
  };
  thumbnail?: string;
}

const FileList: React.FC<Props> = ({ file }) => {
  return (
    <Link to={`/files/${file.id}`}>
      <Wrapper>
        <FileInfo rich file={file} />
      </Wrapper>
    </Link>
  );
};

export default FileList;
