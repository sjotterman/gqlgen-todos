'use client';

import {
  ApolloClient,
  ApolloLink,
  HttpLink,
  SuspenseCache,
} from '@apollo/client';
import {
  ApolloNextAppProvider,
  NextSSRInMemoryCache,
  SSRMultipartLink,
} from '@apollo/experimental-nextjs-app-support/ssr';
import { useAuth } from '@clerk/nextjs';
import { setContext } from '@apollo/client/link/context';

type UseAuth = ReturnType<typeof useAuth>;
type GetToken = UseAuth['getToken'];

/*
 * This client is used in the browser
 */

const GRAPHQL_URL =
  process.env.NEXT_PUBLIC_GRAPHQL_URL ?? 'http://localhost:8080/query';

let token: string | null | undefined = null;

const getMakeClient = (getToken: GetToken) => {
  const asyncAuthLink = setContext((_request) => {
    const asyncGetToken = async () => {
      if (token != null) {
        return token;
      }
      token = await getToken();
      return token;
    };
    return asyncGetToken().then((token) => {
      if (token == null || token === '') {
        return {};
      }
      return {
        headers: {
          authorization: `Bearer ${token}`,
        },
      };
    });
  });

  const makeClient = () => {
    const httpLinkWithCookie = new HttpLink({
      uri: GRAPHQL_URL,
      credentials: 'include',
    });

    const httpLinkWithoutCookie = new HttpLink({
      uri: GRAPHQL_URL,
      credentials: 'include',
    });

    return new ApolloClient({
      cache: new NextSSRInMemoryCache(),
      link:
        typeof window === 'undefined'
          ? ApolloLink.from([
              new SSRMultipartLink({
                stripDefer: true,
              }),
              httpLinkWithCookie,
            ])
          : ApolloLink.from([asyncAuthLink, httpLinkWithoutCookie]),
    });
  };
  return makeClient;
};

function makeSuspenseCache() {
  return new SuspenseCache();
}

export const ApolloWrapper = ({ children }: React.PropsWithChildren) => {
  const { getToken } = useAuth();
  return (
    <ApolloNextAppProvider
      makeClient={getMakeClient(getToken)}
      makeSuspenseCache={makeSuspenseCache}
    >
      {children}
    </ApolloNextAppProvider>
  );
};
