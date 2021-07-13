import React from "react";
import { Link } from "react-router-dom";

import FileInfo from "../FileInfo";

import { Wrapper, Thumbnail } from "./styles";
import FileThumbnail from "../FileThumbnail";

interface Props {
  file: {
    id: string;
    name: string;
    url: string;
  };
  thumbnail?: string;
}

const FileCard: React.FC<Props> = ({ file, thumbnail }) => {
  return (
    <Link to={`/files/${file.id}`}>
      <Wrapper>
        <Thumbnail>
          <FileThumbnail thumbnail={thumbnail} />
        </Thumbnail>

        <FileInfo file={{ id: file.id, name: file.name, url: file.url }} />
      </Wrapper>
    </Link>
  );
};

export default FileCard;
