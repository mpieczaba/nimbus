import React from "react";
import { render } from "react-dom";
import { Provider } from "react-redux";
import { ApolloProvider } from "@apollo/client";

import store from "./store";
import client from "./graphql/client";

import App from "./app";

render(
  <React.StrictMode>
    <Provider store={store}>
      <ApolloProvider client={client}>
        <App />
      </ApolloProvider>
    </Provider>
  </React.StrictMode>,
  document.getElementById("root")
);
