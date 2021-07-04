import React from "react";

import FileInfo from "../FileInfo";

import { Wrapper, Thumbnail } from "./styles";
import FileThumbnail from "../FileThumbnail";

interface Props {
  fileName: string;
  fileURL: string;
  thumbnail?: string;
}

const FileCard: React.FC<Props> = ({ fileName, fileURL, thumbnail }) => {
  return (
    <Wrapper>
      <Thumbnail>
        <FileThumbnail thumbnail={thumbnail} />
      </Thumbnail>
      <FileInfo file={{ name: fileName, url: fileURL }} />
    </Wrapper>
  );
};

export default FileCard;
