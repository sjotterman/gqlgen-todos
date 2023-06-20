import './globals.css';
import Link from 'next/link';
import { UserButton } from '@clerk/nextjs';
import { ApolloWrapper } from '@/lib/apollo-provider';
import { ClerkProvider } from '@clerk/nextjs';


export const metadata = {
  title: 'Create Next App',
  description: 'Generated by create next app',
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <ClerkProvider>
      <html lang="en">
        <head>
          <link
            rel="stylesheet"
            href="https://cdn.jsdelivr.net/npm/@picocss/pico@1/css/pico.min.css"
          ></link>
        </head>
        <body>
          <div className="hero" data-theme="dark">
            <nav className="container-fluid">
              <ul>
                <li>
                  <Link className="contrast" href="/">
                    Home
                  </Link>
                </li>
              </ul>
              <ul dir="rtl">
                <li>
                  <UserButton afterSignOutUrl="/" />
                </li>
              </ul>
            </nav>
            <header className="container">
              <hgroup>
                <h1>Title</h1>
                <h2>Subtitle</h2>
              </hgroup>
            </header>
          </div>
          <ApolloWrapper>{children}</ApolloWrapper>
        </body>
      </html>
    </ClerkProvider>
  );
}
