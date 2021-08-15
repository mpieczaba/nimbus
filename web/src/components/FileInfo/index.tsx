import React, { useState } from "react";
import {
  IconDownload,
  IconPencil,
  IconInfoCircle,
  IconShare,
  IconTrash,
  IconDotsVertical,
} from "@tabler/icons";

import { colors } from "../../themes/colors";

import {
  Wrapper,
  Thumbnail,
  FileInfoWrapper,
  FileName,
  FileInfoElements,
  FileInfoElement,
  FileMenuButton,
  DropdownItemName,
  DropdownItemDelete,
} from "./styles";

import FileThumbnail from "../FileThumbnail";
import Dropdown, { DropdownItem, DropdownItemIcon } from "../Dropdown";

interface Props {
  rich?: boolean;
  file: {
    id: string;
    name: string;
    extension: string;
    size?: string;
    url: string;
    downloadURL: string;
    modificationDate?: string;
  };
}

const FileInfo: React.FC<Props> = ({ rich, file }) => {
  const [dropdown, showHideDropdown] = useState<boolean>(false);

  const handleDownload = () => {
    window.open(file.downloadURL, "_blank");
  };

  return (
    <Wrapper>
      <Thumbnail rich={rich}>
        <FileThumbnail extension={file.extension} url={file.url} image={rich} />
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

      <FileMenuButton onClick={() => showHideDropdown(true)}>
        <IconDotsVertical />
      </FileMenuButton>

      <Dropdown active={dropdown} hideDropdown={() => showHideDropdown(false)}>
        <DropdownItemName>
          <Thumbnail>
            <FileThumbnail extension={file.extension} />
          </Thumbnail>
          <span>{file.name}</span>
        </DropdownItemName>

        <DropdownItem>
          <DropdownItemIcon>
            <IconInfoCircle />
          </DropdownItemIcon>
          Details
        </DropdownItem>

        <DropdownItem>
          <DropdownItemIcon>
            <IconShare />
          </DropdownItemIcon>
          Share
        </DropdownItem>

        <DropdownItem onClick={handleDownload}>
          <DropdownItemIcon>
            <IconDownload />
          </DropdownItemIcon>
          Download
        </DropdownItem>

        <DropdownItem>
          <DropdownItemIcon>
            <IconPencil />
          </DropdownItemIcon>
          Change name
        </DropdownItem>

        <DropdownItemDelete>
          <DropdownItemIcon>
            <IconTrash color={colors.error} />
          </DropdownItemIcon>
          Delete
        </DropdownItemDelete>
      </Dropdown>
    </Wrapper>
  );
};

export default FileInfo;
