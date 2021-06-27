import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import { createGlobalStyle } from "styled-components";

import { style, AppWrapper } from "../themes/styles";

import Home from "../views/Home";

// Font awesome imports
import { library } from "@fortawesome/fontawesome-svg-core";
import { fas } from "@fortawesome/free-solid-svg-icons";
import {
  faSearch,
  faEllipsisV,
  faGripVertical,
  faThList,
  faArrowDown,
  faBars,
} from "@fortawesome/free-solid-svg-icons";

library.add(
  fas,
  faSearch,
  faEllipsisV,
  faGripVertical,
  faThList,
  faArrowDown,
  faBars
);

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
