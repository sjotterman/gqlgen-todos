import { getClient } from '@/lib/client';
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

export default async function Home() {
  const client = getClient();
  const { data } = await client.query<
    RestaurantsQuery,
    RestaurantsQueryVariables
  >({ query });
  const restaurants = data.restaurants;
  return (
    <main className="container">
      <div>
        <h1>Hello, world</h1>
      </div>
      <Link href="/restaurants">
        <p>Restaurants</p>
      </Link>
      <div>
        {restaurants.map((restaurant) => {
          return <div key={restaurant.id}>{restaurant.name}</div>;
        })}
      </div>
    </main>
  );
}
