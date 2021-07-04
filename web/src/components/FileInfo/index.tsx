import React, { useState } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

import {
  Wrapper,
  Thumbnail,
  FileInfoWrapper,
  FileName,
  FileInfoElements,
  FileInfoElement,
  FileMenuButton,
  DropdownItem,
} from "./styles";

import FileThumbnail from "../FileThumbnail";
import Dropdown, { DropdownItemsWrapper } from "../Dropdown";

interface Props {
  rich?: boolean;
  file: {
    name: string;
    size?: string;
    url: string;
    modificationDate?: string;
  };
  thumbnail?: string;
}

const FileInfo: React.FC<Props> = ({ rich, file, thumbnail }) => {
  const [dropdown, showHideDropdown] = useState<boolean>(false);

  const handleDropdownShowHide = () => {
    showHideDropdown(!dropdown);
  };

  const handleDownload = () => {
    window.open(file.url, "_blank");
  };

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
      <FileMenuButton onClick={handleDropdownShowHide}>
        <FontAwesomeIcon icon="ellipsis-v" />
      </FileMenuButton>
      {dropdown ? (
        <Dropdown onClick={handleDropdownShowHide}>
          <DropdownItemsWrapper>
            <DropdownItem>Details</DropdownItem>
            <DropdownItem onClick={handleDownload}>Download</DropdownItem>
          </DropdownItemsWrapper>
        </Dropdown>
      ) : null}
    </Wrapper>
  );
};

export default FileInfo;
