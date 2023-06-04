'use client';

import { useSuspenseQuery } from '@apollo/experimental-nextjs-app-support/ssr';

import { gql } from '@apollo/client';

const query = gql`
  query restaurants {
    restaurants {
      id
      name
      description
      phoneNumber
    }
  }
`;

export default function Page() {
  const { data } = useSuspenseQuery(query);
  console.log({ data });

  return (
    <main>
      <div>
        {data.restaurants.map((restaurant) => {
          return <div key={restaurant.id}>{restaurant.name}</div>;
        })}
      </div>
    </main>
  );
}
