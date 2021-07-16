import React from "react";
import { withRouter } from "react-router-dom";

import LoginForm from "../../components/LoginForm";

const Login: React.FC = () => {
  return <LoginForm />;
};

export default withRouter(Login);
