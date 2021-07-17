import React from "react";

import { Wrapper, Thumbnail } from "./styles";

import FileIcon from "../FileIcon";

interface Props {
  thumbnail?: string;
}

const FileThumbnail: React.FC<Props> = ({ thumbnail }) => {
  return (
    <Wrapper>
      {thumbnail ? <Thumbnail src={thumbnail} /> : <FileIcon icon="image" />}
    </Wrapper>
  );
};

export default FileThumbnail;
