import React from "react";
import { Link } from "react-router-dom";

import FileInfo from "../FileInfo";

import { Wrapper, Thumbnail } from "./styles";
import FileThumbnail from "../FileThumbnail";

interface Props {
  file: {
    id: string;
    name: string;
    extension: string;
    url: string;
    downloadURL: string;
  };
}

const FileCard: React.FC<Props> = ({ file }) => {
  return (
    <Link to={`/files/${file.id}`}>
      <Wrapper>
        <Thumbnail>
          <FileThumbnail extension={file.extension} url={file.url} image />
        </Thumbnail>

        <FileInfo file={file} />
      </Wrapper>
    </Link>
  );
};

export default FileCard;
