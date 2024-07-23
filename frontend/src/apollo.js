import { ApolloClient, createHttpLink, InMemoryCache } from '@apollo/client/core'


const httpLink = createHttpLink({
  uri: 'http://localhost:8081/graphql', // Go server endpoint
});

// const authLink = setContext((_, { headers }) => {
//   const token = localStorage.getItem('auth-token');
//   return {
//     headers: { ...headers, authorization: token ? `Bearer ${token}` : '' },
//   };
// });
const cache = new InMemoryCache()
export const apolloClient = new ApolloClient({
  link: httpLink,
  cache,
});

