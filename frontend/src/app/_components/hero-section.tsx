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
              <span className="absolute inline-flex h-full w-full animate-ping rounded-full bg-green-400 opacity-75" />
              <span className="relative inline-flex h-2 w-2 rounded-full bg-green-500" />
            </span>
            Enterprise-grade reliability and security
          </div>

          {/* Headline */}
          <h1 className="max-w-4xl text-center text-4xl font-bold tracking-tight sm:text-5xl md:text-6xl lg:text-7xl animate-in fade-in slide-in-from-bottom-4 duration-700 delay-100">
            The Enterprise Link{" "}
            <span className="bg-linear-to-r from-primary via-accent to-secondary bg-clip-text text-transparent">
              Management Platform
            </span>
          </h1>

          {/* Subheading */}
          <p className="mt-6 max-w-2xl text-center text-lg text-muted-foreground sm:text-xl animate-in fade-in slide-in-from-bottom-4 duration-700 delay-200">
            Secure, scalable, and compliant link routing for modern teams.
            Control your brand's digital touchpoints with granular analytics and robust developer APIs.
          </p>

          {/* Auth CTA */}
          <div className="mt-10 flex flex-col sm:flex-row items-center gap-4 animate-in fade-in slide-in-from-bottom-4 duration-700 delay-300">
            <AuthModal defaultTab="register">
              <Button
                size="lg"
                className="rounded-xl font-semibold px-8 hover:-translate-y-0.5 transition-transform"
              >
                Get Started
              </Button>
            </AuthModal>
            <Link href="/dashboard">
              <Button
                variant="outline"
                size="lg"
                className="rounded-xl font-medium px-8 hover:-translate-y-0.5 transition-transform"
              >
                View Documentation
              </Button>
            </Link>
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
