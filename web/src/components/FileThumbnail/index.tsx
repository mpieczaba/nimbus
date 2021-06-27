import { Wrapper, Thumbnail } from "./styles";
import FileIcon from "../FileIcon";
import React from "react";

interface Props {
  thumbnail?: string;
}

const FileThumbnail = ({ thumbnail }: Props) => {
  return (
    <Wrapper>
      {thumbnail ? <Thumbnail src={thumbnail} /> : <FileIcon icon="image" />}
    </Wrapper>
  );
};

export default FileThumbnail;
