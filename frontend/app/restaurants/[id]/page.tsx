import { getClient } from '@/lib/client';
import { gql } from '@apollo/client';
import { PageProps } from '@/.next/types/app/layout';

const query = gql`
  query getRestaurant($id: ID!) {
    restaurant(id: $id) {
      id
      name
      description
      phoneNumber
    }
  }
`;

export default async function Page({ params }: PageProps) {
  const { id } = params;
  console.log({ params });
  const client = getClient();
  const { data } = await client.query({ query, variables: { id } });
  const restaurant = data.restaurant;
  console.log({ restaurant });
  return (
    <main className="container">
      <div>
        <h1>{restaurant.name}</h1>
      </div>
      <div>
        <p>{restaurant.description}</p>
        <p>{restaurant.phoneNumber}</p>
      </div>
    </main>
  );
}
