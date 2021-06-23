import React from "react";
import styled from "styled-components";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

import colors from "../utils/colors";

import { Row } from "./Row";
import { Column } from "./Column";
import { FileIcon } from "./FileIcon";

const Wrapper = styled(Column)`
  background: ${colors.gray};
  border-radius: 10px;
  box-shadow: ${colors.boxShadow};
  cursor: pointer;
`;

const ThumbnailWrapper = styled.div`
  display: flex;
  width: 100%;
  height: 100px;
`;

const Thumbnail = styled.img`
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 10px;
  box-shadow: ${colors.boxShadow};
`;

const ThumbnailIcon = styled.div`
  display: flex;
  width: 100%;
  height: 100%;
  padding: 2rem 2rem 1em 2rem;
  filter: grayscale(100%);
`;

const FileInfo = styled(Row)`
  justify-content: space-between;
`;

const FileInfoIcon = styled.div`
  display: flex;
  width: 1.5rem;
  height: 1.5rem;
  margin: 0.5rem 0 0.5rem 0.5rem;
`;

const FileName = styled.div`
  display: inline-block;
  align-self: center;
  padding: 0.5rem 0;
  font-size: 0.8rem;
  font-weight: 600;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
`;

const Menu = styled.button`
  display: flex;
  align-self: center;
  padding: 0.5rem;
  background: none;
  border: none;
  cursor: pointer;
`;

export const FileCard: React.FC<{ name: string; thumbnail?: string }> = ({
  name,
  thumbnail,
}) => {
  return (
    <Wrapper>
      <ThumbnailWrapper>
        {thumbnail ? (
          <Thumbnail src={thumbnail} />
        ) : (
          <ThumbnailIcon>
            <FileIcon icon="image" />
          </ThumbnailIcon>
        )}
      </ThumbnailWrapper>
      <FileInfo>
        <FileInfoIcon>
          <FileIcon icon="image" />
        </FileInfoIcon>
        <FileName>{name}</FileName>
        <Menu>
          <FontAwesomeIcon icon="ellipsis-v" color={colors.text} />
        </Menu>
      </FileInfo>
    </Wrapper>
  );
};
