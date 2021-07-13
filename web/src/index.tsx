import React from "react";
import { render } from "react-dom";
import { ApolloProvider } from "@apollo/client";

import client from "./graphql/client";
import { App } from "./App";

render(
  <React.StrictMode>
    <ApolloProvider client={client}>
      <App />
    </ApolloProvider>
  </React.StrictMode>,
  document.getElementById("root")
);
