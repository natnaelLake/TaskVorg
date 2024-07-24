import { ApolloClient, createHttpLink, InMemoryCache } from '@apollo/client/core'


const httpLink = createHttpLink({
  uri: 'http://localhost:8081/graphql',
});

const cache = new InMemoryCache()
export const apolloClient = new ApolloClient({
  link: httpLink,
  cache,
});

