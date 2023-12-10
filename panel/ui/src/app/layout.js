import './globals.scss';

export const metadata = {
  title: 'PlutoEngine Panel',
  description: 'PlutoEngine administrator panel',
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
