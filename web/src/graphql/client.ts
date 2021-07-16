import { ApolloClient, createHttpLink, InMemoryCache } from "@apollo/client";

// TODO: Move to redux and env variables
const client = new ApolloClient({
  link: createHttpLink({
    uri: "http://192.168.1.18:8080/graphql",
    headers: {
      authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjY2MzA2MjcsImlkIjoiYzMxN2xyb2l0bjcxYm1hM2RoYzAiLCJ1c3IiOiJkdXBhIiwia2luZCI6IkFETUlOIn0.0NBBKOihBTHqPa76K_Rhmy00TA0tOKJ0vPSt9ikTF64`,
    },
  }),

  cache: new InMemoryCache(),
});

export default client;
