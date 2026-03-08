import { Navbar } from "@/components/navbar";
import { FeaturesSection } from "./_components/features-section";
import { HeroSection } from "./_components/hero-section";

// ── Page ─────────────────────────────────────────────────────────────────

export default function Home() {
  return (
    <div className="relative min-h-screen">
      <Navbar />
      <HeroSection />
      <FeaturesSection />

      {/* Footer */}
      <footer className="border-t bg-muted/20">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8 py-10">
          <div className="flex flex-col items-center justify-between gap-4 sm:flex-row">
            <div className="flex items-center gap-2">
              <div className="flex h-7 w-7 items-center justify-center rounded-md bg-primary">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="14"
                  height="14"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  strokeWidth="2.5"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  className="text-primary-foreground"
                >
                  <path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71" />
                  <path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71" />
                </svg>
              </div>
              <span className="text-sm font-semibold">Snip.ly</span>
            </div>
            <p className="text-sm text-muted-foreground">
              © {new Date().getFullYear()} Snip.ly. Built with ♥ by Chirag
              Bhalotia.
            </p>
          </div>
        </div>
      </footer>
    </div>
  );
}
