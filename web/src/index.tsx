import React from "react";
import { render } from "react-dom";
import { ApolloProvider } from "@apollo/client";

import client from "./apollo/client";
import { App } from "./app";

render(
  <React.StrictMode>
    <ApolloProvider client={client}>
      <App />
    </ApolloProvider>
  </React.StrictMode>,
  document.getElementById("root")
);
