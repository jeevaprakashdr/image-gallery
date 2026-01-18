import type { Metadata } from "next";
import "./globals.css";


export const metadata: Metadata = {
  title: "Image Gallery",
  description: "",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <main className="m-10 max-w-4xl mx-auto">
          {children}
        </main>
      </body>
    </html>
  );
}
