import { Wrapper } from "./styles";

import FileInfo from "../FileInfo";

interface Props {
  file: {
    name: string;
    size?: string;
    modificationDate?: string;
  };
  thumbnail?: string;
}

const FileList = ({ file, thumbnail }: Props) => {
  return (
    <Wrapper>
      <FileInfo rich file={file} thumbnail={thumbnail} />
    </Wrapper>
  );
};

export default FileList;
