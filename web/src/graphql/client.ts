import { ApolloClient, createHttpLink, InMemoryCache } from "@apollo/client";

const token = localStorage.getItem("token");

// TODO: Move to redux and env variables
const client = new ApolloClient({
  link: createHttpLink({
    uri: "http://192.168.1.18:8080/graphql",
    headers: {
      Authorization: token ? `Bearer ${token}` : "",
    },
  }),

  cache: new InMemoryCache(),
});

export default client;
