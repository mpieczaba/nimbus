import { ApolloClient, createHttpLink, InMemoryCache } from "@apollo/client";

// TODO: Move to redux and env variables
const client = new ApolloClient({
  link: createHttpLink({
    uri: "http://192.168.1.19:8080/graphql",
    headers: {
      authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjYzNDE4MTAsImlkIjoiYzMxN2xyb2l0bjcxYm1hM2RoYzAiLCJ1c3IiOiJkdXBhIiwia2luZCI6IkFETUlOIn0.I1YQm7yS422ywj7nAFo1Gle5waRmFKeW1I33hL3L0ZM`,
    },
  }),

  cache: new InMemoryCache(),
});

export default client;
