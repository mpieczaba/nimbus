import React from "react";
import { withRouter } from "react-router-dom";

import Files from "../../components/Files";

const Home: React.FC = () => {
  return <Files />;
};

export default withRouter(Home);
