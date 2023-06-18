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
import { useEffect, useMemo, useState } from 'react';
import { useAuth } from '@clerk/nextjs';

/*
 * This client is used in the browser
 */

const GRAPHQL_URL =
  process.env.NEXT_PUBLIC_GRAPHQL_URL ?? 'http://localhost:8080/query';

const getMakeClient = (token: string) => {
  const makeClient = () => {
    const httpLink = new HttpLink({
      uri: GRAPHQL_URL,
      credentials: 'include',
      headers: {
        ...(token != null ? { authorization: `Bearer ${token}` } : {}),
      },
    });

    return new ApolloClient({
      cache: new NextSSRInMemoryCache(),
      link:
        typeof window === 'undefined'
          ? ApolloLink.from([
              new SSRMultipartLink({
                stripDefer: true,
              }),
              httpLink,
            ])
          : ApolloLink.from([httpLink]),
    });
  };
  return makeClient;
};

function makeSuspenseCache() {
  return new SuspenseCache();
}

export const ApolloWrapper = ({ children }: React.PropsWithChildren) => {
  const { getToken } = useAuth();
  const [token, setToken] = useState<string | null>();
  useEffect(() => {
    const asyncFunc = async () => {
      const newToken = await getToken();
      setToken(newToken);
    };
    asyncFunc().catch((e) => {
      console.error('error in async func', e);
    });
  });
  const client = useMemo(() => {
    if (token == null) {
      return null;
    }
    return getMakeClient(token);
  }, [token]);
  if (client == null) {
    return null;
  }
  return (
    <ApolloNextAppProvider
      makeClient={client}
      makeSuspenseCache={makeSuspenseCache}
    >
      {children}
    </ApolloNextAppProvider>
  );
};
