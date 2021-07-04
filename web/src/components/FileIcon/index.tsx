import React from "react";
import icons from "@uiw/file-icons/fonts/ffont.symbol.svg";

interface Props {
  icon: string;
}

const FileIcon: React.FC<Props> = ({ icon }) => {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      xmlnsXlink="http://www.w3.org/1999/xlink"
      aria-hidden="true"
    >
      <use xlinkHref={`${icons}#ffont-${icon}`} />
    </svg>
  );
};

export default FileIcon;
