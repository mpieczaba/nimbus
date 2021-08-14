import React, { useState, MouseEvent } from "react";
import {
  IconDownload,
  IconPencil,
  IconInfoCircle,
  IconShare,
  IconTrash,
  IconDotsVertical,
} from "@tabler/icons";

import { useAppDispatch } from "../../hooks/store";
import { setScrollable } from "../../store/actions/uiActions";

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
import Dropdown, {
  DropdownItemsWrapper,
  DropdownItem,
  DropdownItemIcon,
} from "../Dropdown";

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
  const dispatch = useAppDispatch();

  const [dropdown, showHideDropdown] = useState<boolean>(false);

  const handleDropdownShowHide = (e: MouseEvent) => {
    e.preventDefault();

    dispatch(setScrollable(!dropdown));

    showHideDropdown(!dropdown);
  };

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

      <FileMenuButton onClick={handleDropdownShowHide}>
        <IconDotsVertical />
      </FileMenuButton>

      {dropdown ? (
        <Dropdown onClick={handleDropdownShowHide}>
          <DropdownItemsWrapper>
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
          </DropdownItemsWrapper>
        </Dropdown>
      ) : null}
    </Wrapper>
  );
};

export default FileInfo;
