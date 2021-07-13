import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import { createGlobalStyle } from "styled-components";

import { style, AppWrapper } from "./themes/styles";

import Home from "./views/Home";

const GlobalStyle = createGlobalStyle`${style}`;

export const App: React.FC = () => {
  return (
    <>
      <GlobalStyle />
      <AppWrapper>
        <Router>
          <Switch>
            <Route path="/" component={Home} />
          </Switch>
        </Router>
      </AppWrapper>
    </>
  );
};
