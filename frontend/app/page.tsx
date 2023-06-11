import { getClient } from '@/lib/client';
import { gql } from '@apollo/client';
import { RestaurantsQuery, RestaurantsQueryVariables } from '@/graphql/graphql';

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
  console.log({ data });
  const restaurants = data.restaurants;
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div>
        <h1 className="text-4xl font-bold">Hello, world</h1>
      </div>
      <div>
        {restaurants.map((restaurant) => {
          return <div key={restaurant.id}>{restaurant.name}</div>;
        })}
      </div>
    </main>
  );
}
