'use client';

import { useSuspenseQuery } from '@apollo/experimental-nextjs-app-support/ssr';

import { gql } from '@apollo/client';
import { RestaurantsQuery, RestaurantsQueryVariables } from '@/graphql/graphql';
import Link from 'next/link';

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
  const { data } = useSuspenseQuery<
    RestaurantsQuery,
    RestaurantsQueryVariables
  >(query);

  const restaurants = data.restaurants;
  return (
    <main className="container">
      <div>
        <ul>
          {restaurants.map((restaurant) => {
            return (
              <li key={restaurant.id}>
                <Link href={`/restaurants/${restaurant.id}`}>
                  {restaurant.name}
                </Link>
              </li>
            );
          })}
        </ul>
      </div>
    </main>
  );
}
