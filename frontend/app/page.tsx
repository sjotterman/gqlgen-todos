import { getClient } from '@/lib/client';
import { gql } from '@apollo/client';
import { UserButton } from '@clerk/nextjs';
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
  const restaurants = data.restaurants;
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <UserButton afterSignOutUrl="/" />
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
