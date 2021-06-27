import { ApolloClient, createHttpLink, InMemoryCache } from "@apollo/client";

// TODO: Move to redux and env variables
const client = new ApolloClient({
  link: createHttpLink({
    uri: "http://192.168.1.19:8080/graphql",
    headers: {
      authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjUwNDg1MTMsImlkIjoiYzMxN2xyb2l0bjcxYm1hM2RoYzAiLCJ1c3IiOiJkdXBhIiwia2luZCI6IkFETUlOIn0.94SpPeWokH_YEa0DhmTsjpI2UZI3Rnv_L75dnHH7EjQ`,
    },
  }),

  cache: new InMemoryCache(),
});

export default client;
