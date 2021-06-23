import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

import "./App.scss";

import Home from "../Home";

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

const App: React.FC = () => {
  return (
    <Router>
      <div className="app">
        <Switch>
          <Route path="/">
            <Home />
          </Route>
        </Switch>
      </div>
    </Router>
  );
};

export default App;
