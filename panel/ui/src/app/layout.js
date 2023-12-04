import './globals.scss';
import { Providers } from './providers';

export const metadata = {
  title: 'PlutoEngine Panel',
  description: 'PlutoEngine administrator panel',
};

export default function RootLayout({ children }) {
  return (
    <html lang='en'>
    <body>
    <Providers>
      {children}
    </Providers>
    </body>
    </html>
  );
}
