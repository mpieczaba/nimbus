import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import { createGlobalStyle } from "styled-components";

import PrivateRoute from "../utils/PrivateRoute";

import { style, AppWrapper } from "../themes/styles";

import Home from "../views/Home";
import Login from "../views/Login";

import Navbar from "../components/Navbar";

const GlobalStyle = createGlobalStyle`${style}`;

const App: React.FC = () => {
  return (
    <>
      <GlobalStyle />
      <AppWrapper>
        <Router>
          <Switch>
            <Route exact path="/login">
              <Login />
            </Route>
            <PrivateRoute path="/">
              <Home />
            </PrivateRoute>
          </Switch>
          <Navbar />
        </Router>
      </AppWrapper>
    </>
  );
};

export default App;
