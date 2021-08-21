import { ApolloClient, InMemoryCache } from "@apollo/client";
// @ts-ignore
import { createUploadLink } from "apollo-upload-client";
import { relayStylePagination } from "@apollo/client/utilities";

const token = localStorage.getItem("token");

const cache = new InMemoryCache({
  typePolicies: {
    Query: {
      fields: {
        files: relayStylePagination(["name", "tags"]),
      },
    },
  },
});

// TODO: Move to redux and env variables
const client = new ApolloClient({
  link: createUploadLink({
    uri: `http://${process.env.REACT_APP_BACKEND_HOST}/graphql`,
    headers: {
      Authorization: token ? `Bearer ${token}` : "",
    },
  }),

  cache: cache,
});

export default client;
