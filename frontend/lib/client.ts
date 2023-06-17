import { ApolloClient, HttpLink, from } from '@apollo/client';
import { registerApolloClient } from '@apollo/experimental-nextjs-app-support/rsc';
import { setContext } from '@apollo/client/link/context';
import { auth } from '@clerk/nextjs';
import { NextSSRInMemoryCache } from '@apollo/experimental-nextjs-app-support/ssr';

/*
 * This client is used on the server.
 */

const GRAPHQL_URL =
  process.env.NEXT_PUBLIC_GRAPHQL_URL ?? 'http://localhost:8080/query';
export const { getClient } = registerApolloClient(() => {
  const authMiddleware = setContext(async (_operation, { headers }) => {
    const { getToken } = auth();
    const token = await getToken();

    return {
      headers: {
        ...headers,
        authorization: `Bearer ${token}`,
      },
    };
  });

  const httpLink = new HttpLink({
    uri: GRAPHQL_URL,
    credentials: 'include',
  });

  return new ApolloClient({
    cache: new NextSSRInMemoryCache(),
    link: from([authMiddleware, httpLink]),
  });
});
