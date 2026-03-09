import Link from "next/link";
import { AuthModal } from "@/components/auth-modal";
import { QuickShortener } from "@/components/quick-shortener";
import { Button } from "@/components/ui/button";

export function HeroSection() {
  return (
    <section className="relative overflow-hidden">
      {/* Background gradient blobs */}
      <div className="absolute inset-0 -z-10">
        <div className="absolute top-0 left-1/4 h-[600px] w-[600px] rounded-full bg-primary/10 blur-[120px]" />
        <div className="absolute bottom-0 right-1/4 h-[500px] w-[500px] rounded-full bg-accent/10 blur-[120px]" />
        <div className="absolute top-1/2 left-1/2 h-[400px] w-[400px] -translate-x-1/2 -translate-y-1/2 rounded-full bg-secondary/8 blur-[100px]" />
      </div>

      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="flex flex-col items-center pt-24 pb-20 sm:pt-32 sm:pb-28">
          {/* Badge */}
          <div className="mb-8 inline-flex items-center gap-2 rounded-full border bg-card/80 px-4 py-1.5 text-sm font-medium text-muted-foreground backdrop-blur-sm animate-in fade-in slide-in-from-bottom-3 duration-700">
            <span className="relative flex h-2 w-2">
              <span className="absolute inline-flex h-full w-full animate-ping rounded-full bg-blue-400 opacity-75" />
              <span className="relative inline-flex h-2 w-2 rounded-full bg-blue-500" />
            </span>
            Open Source Project
          </div>

          {/* Headline */}
          <h1 className="max-w-4xl text-center text-4xl font-bold tracking-tight sm:text-5xl md:text-6xl lg:text-7xl animate-in fade-in slide-in-from-bottom-4 duration-700 delay-100">
            A High-Performance{" "}
            <span className="bg-linear-to-r from-primary via-accent to-secondary bg-clip-text text-transparent">
              URL Shortener
            </span>
          </h1>

          {/* Subheading */}
          <p className="mt-6 max-w-2xl text-center text-lg text-muted-foreground sm:text-xl animate-in fade-in slide-in-from-bottom-4 duration-700 delay-200">
            A full-stack, distributed architecture built to demonstrate high
            throughput. Powered by Go microservices, Redis Streams, Next.js, and
            a custom Hyperflake ID generator.
          </p>

          {/* Auth CTA */}
          <div className="mt-10 flex flex-col sm:flex-row items-center gap-4 animate-in fade-in slide-in-from-bottom-4 duration-700 delay-300">
            <Link
              href="https://github.com/chirag3003/url-shortener"
              target="_blank"
              rel="noopener noreferrer"
            >
              <Button
                size="lg"
                className="rounded-xl font-semibold px-8 hover:-translate-y-0.5 transition-transform gap-2"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="20"
                  height="20"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  strokeWidth="2"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                >
                  <path d="M15 22v-4a4.8 4.8 0 0 0-1-3.5c3 0 6-2 6-5.5.08-1.25-.27-2.48-1-3.5.28-1.15.28-2.35 0-3.5 0 0-1 0-3 1.5-2.64-.5-5.36-.5-8 0C6 2 5 2 5 2c-.3 1.15-.3 2.35 0 3.5A5.403 5.403 0 0 0 4 9c0 3.5 3 5.5 6 5.5-.39.49-.68 1.05-.85 1.65-.17.6-.22 1.23-.15 1.85v4" />
                  <path d="M9 18c-4.51 2-5-2-7-2" />
                </svg>
                View Source Code
              </Button>
            </Link>
            {/* <Link href="/dashboard">
              <Button
                variant="outline"
                size="lg"
                className="rounded-xl font-medium px-8 hover:-translate-y-0.5 transition-transform"
              >
                Try the Dashboard Demo
              </Button>
            </Link> */}
          </div>

          {/* Quick Shortener (Interactive Demo) */}
          <div className="mt-16 w-full max-w-2xl animate-in fade-in slide-in-from-bottom-4 duration-700 delay-500">
            <div className="text-center mb-4 text-sm font-medium text-muted-foreground uppercase tracking-wide">
              Try the API Sandbox
            </div>
            <div className="p-1 rounded-2xl bg-linear-to-r from-primary/10 via-accent/10 to-secondary/10 shadow-lg border bg-card/50 backdrop-blur-sm">
              <QuickShortener />
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
