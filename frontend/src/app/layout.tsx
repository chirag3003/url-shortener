import type { Metadata } from "next";
import { Geist_Mono, Inter } from "next/font/google";
import "./globals.css";
import { ThemeProvider } from "@/components/theme-provider";
import { Toaster } from "@/components/ui/sonner";
import { TooltipProvider } from "@/components/ui/tooltip";
import { cn } from "@/lib/utils";
import { QueryProvider } from "@/providers/query-provider";

const inter = Inter({
  subsets: ["latin"],
  variable: "--font-sans",
  display: "swap",
});

const geistMono = Geist_Mono({
  variable: "--font-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Snip.ly — Shorten, Share, Track",
  description:
    "A blazing-fast URL shortener with powerful analytics. Create short links, custom aliases, and track every click in real time.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html
      lang="en"
      className={cn(inter.variable, geistMono.variable)}
      suppressHydrationWarning
    >
      <body className="font-sans antialiased">
        <QueryProvider>
          <ThemeProvider defaultTheme="dark">
            <TooltipProvider delayDuration={300}>
              <main>{children}</main>
              <Toaster richColors position="bottom-right" />
            </TooltipProvider>
          </ThemeProvider>
        </QueryProvider>
      </body>
    </html>
  );
}
