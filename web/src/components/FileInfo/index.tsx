import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

import {
  Wrapper,
  Thumbnail,
  FileInfoWrapper,
  FileName,
  FileInfoElements,
  FileInfoElement,
  FileMenuButton,
} from "./styles";
import FileThumbnail from "../FileThumbnail";

interface Props {
  rich?: boolean;
  file: {
    name: string;
    size?: string;
    modificationDate?: string;
  };
  thumbnail?: string;
}

const FileInfo = ({ rich, file, thumbnail }: Props) => {
  return (
    <Wrapper>
      <Thumbnail rich={rich}>
        <FileThumbnail thumbnail={thumbnail} />
      </Thumbnail>
      <FileInfoWrapper>
        <FileName rich={rich}>{file.name}</FileName>
        {rich && file.size && file.modificationDate ? (
          <FileInfoElements>
            <FileInfoElement>{file.size}</FileInfoElement>
            <FileInfoElement>{file.modificationDate}</FileInfoElement>
          </FileInfoElements>
        ) : null}
      </FileInfoWrapper>
      <FileMenuButton>
        <FontAwesomeIcon icon="ellipsis-v" />
      </FileMenuButton>
    </Wrapper>
  );
};

export default FileInfo;
