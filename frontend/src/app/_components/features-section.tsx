import { AuthModal } from "@/components/auth-modal";
import { Button } from "@/components/ui/button";

const features = [
  {
    title: "Real-Time Analytics",
    description:
      "Track every click with detailed breakdowns by geography, device, browser, and referrer.",
    icon: (
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="24"
        height="24"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        strokeWidth="1.5"
        strokeLinecap="round"
        strokeLinejoin="round"
      >
        <path d="M3 3v16a2 2 0 0 0 2 2h16" />
        <path d="m19 9-5 5-4-4-3 3" />
      </svg>
    ),
  },
  {
    title: "Custom Aliases",
    description:
      "Create branded, memorable short links with your own custom aliases instead of random codes.",
    icon: (
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="24"
        height="24"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        strokeWidth="1.5"
        strokeLinecap="round"
        strokeLinejoin="round"
      >
        <path d="m16 3 5 5-5 5" />
        <path d="M21 8H8a4 4 0 1 0 0 8h1" />
      </svg>
    ),
  },
  {
    title: "Link Expiry",
    description:
      "Set self-destruct dates on your links for time-sensitive campaigns and promotions.",
    icon: (
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="24"
        height="24"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        strokeWidth="1.5"
        strokeLinecap="round"
        strokeLinejoin="round"
      >
        <circle cx="12" cy="12" r="10" />
        <polyline points="12 6 12 12 16 14" />
      </svg>
    ),
  },
];

const stats = [
  { value: "10M+", label: "Links Created" },
  { value: "250M+", label: "Clicks Tracked" },
  { value: "99.9%", label: "Uptime" },
  { value: "<10ms", label: "Redirect Latency" },
];

export function FeaturesSection() {
  return (
    <>
      {/* Stats Section */}
      <section className="border-y bg-muted/30">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8 py-12 sm:py-16">
          <div className="grid grid-cols-2 gap-8 md:grid-cols-4">
            {stats.map((stat) => (
              <div key={stat.label} className="text-center">
                <p className="text-3xl font-bold tracking-tight text-primary sm:text-4xl">
                  {stat.value}
                </p>
                <p className="mt-1 text-sm text-muted-foreground">
                  {stat.label}
                </p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-20 sm:py-28">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-3xl font-bold tracking-tight sm:text-4xl">
              Everything you need to{" "}
              <span className="text-primary">own your links</span>
            </h2>
            <p className="mt-4 max-w-2xl mx-auto text-lg text-muted-foreground">
              From quick one-off links to full analytics dashboards, Snip.ly
              gives you the tools to make every link count.
            </p>
          </div>

          <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-4">
            {features.map((feature) => (
              <div
                key={feature.title}
                className="group relative rounded-2xl border bg-card p-6 transition-all duration-300 hover:shadow-lg hover:shadow-primary/5 hover:-translate-y-1"
              >
                <div className="mb-4 inline-flex h-12 w-12 items-center justify-center rounded-xl bg-primary/10 text-primary transition-colors group-hover:bg-primary group-hover:text-primary-foreground">
                  {feature.icon}
                </div>
                <h3 className="mb-2 text-lg font-semibold">{feature.title}</h3>
                <p className="text-sm leading-relaxed text-muted-foreground">
                  {feature.description}
                </p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="relative overflow-hidden border-t">
        <div className="absolute inset-0 -z-10">
          <div className="absolute top-0 right-0 h-[400px] w-[400px] rounded-full bg-primary/8 blur-[100px]" />
          <div className="absolute bottom-0 left-0 h-[300px] w-[300px] rounded-full bg-accent/8 blur-[80px]" />
        </div>
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8 py-20 sm:py-28">
          <div className="text-center">
            <h2 className="text-3xl font-bold tracking-tight sm:text-4xl">
              Ready to supercharge your links?
            </h2>
            <p className="mt-4 text-lg text-muted-foreground">
              Join thousands of developers and marketers who trust Snip.ly.
            </p>
            <div className="mt-8 flex justify-center gap-4">
              <AuthModal defaultTab="register">
                <Button
                  size="lg"
                  className="rounded-xl px-8 font-semibold text-base"
                >
                  Get Started Free
                </Button>
              </AuthModal>
            </div>
          </div>
        </div>
      </section>
    </>
  );
}
