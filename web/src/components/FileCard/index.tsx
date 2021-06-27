import React from "react";

import FileInfo from "../FileInfo";

import { Wrapper, Thumbnail } from "./styles";
import FileThumbnail from "../FileThumbnail";

interface Props {
  fileName: string;
  thumbnail?: string;
}

const FileCard = ({ fileName, thumbnail }: Props) => {
  return (
    <Wrapper>
      <Thumbnail>
        <FileThumbnail thumbnail={thumbnail} />
      </Thumbnail>
      <FileInfo file={{ name: fileName }} />
    </Wrapper>
  );
};

export default FileCard;
