import { ApolloClient, createHttpLink, InMemoryCache } from "@apollo/client";
import { relayStylePagination } from "@apollo/client/utilities";

const token = localStorage.getItem("token");

const cache = new InMemoryCache({
  typePolicies: {
    Query: {
      fields: {
        files: relayStylePagination(),
      },
    },
  },
});

// TODO: Move to redux and env variables
const client = new ApolloClient({
  link: createHttpLink({
    uri: "http://192.168.1.18:8080/graphql",
    headers: {
      Authorization: token ? `Bearer ${token}` : "",
    },
  }),

  cache: cache,
});

export default client;
