import { ApolloClient, createHttpLink, InMemoryCache } from "@apollo/client";

// TODO: Move to redux and env variables
const client = new ApolloClient({
  link: createHttpLink({
    uri: "http://192.168.1.19:8080/graphql",
    headers: {
      authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjU2Nzg5NTgsImlkIjoiYzMxN2xyb2l0bjcxYm1hM2RoYzAiLCJ1c3IiOiJkdXBhIiwia2luZCI6IkFETUlOIn0.rV3P6jI21p9x-u_3-hIlp9sx97UcdkGSr3c_O1-vC68`,
    },
  }),

  cache: new InMemoryCache(),
});

export default client;
