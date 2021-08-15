import React from "react";
import { withRouter } from "react-router-dom";

import FilesContainer from "../../components/FilesContainer";
import UploadButton from "../../components/UploadButton";

const Files: React.FC = () => {
  return (
    <>
      <FilesContainer />
      <UploadButton />
    </>
  );
};

export default withRouter(Files);
