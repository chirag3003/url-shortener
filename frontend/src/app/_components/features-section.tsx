import { AuthModal } from "@/components/auth-modal";
import { Button } from "@/components/ui/button";

const features = [
  {
    title: "Secure & Compliant",
    description:
      "Built with enterprise-grade security. Support for SSO (SAML/OIDC), end-to-end encryption, and comprehensive audit logs.",
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
        <rect width="18" height="11" x="3" y="11" rx="2" ry="2" />
        <path d="M7 11V7a5 5 0 0 1 10 0v4" />
      </svg>
    ),
  },
  {
    title: "Team Collaboration",
    description:
      "Manage permissions with Role-Based Access Control (RBAC). Organize links by workspaces, teams, or campaigns seamlessly.",
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
        <path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2" />
        <circle cx="9" cy="7" r="4" />
        <path d="M22 21v-2a4 4 0 0 0-3-3.87" />
        <path d="M16 3.13a4 4 0 0 1 0 7.75" />
      </svg>
    ),
  },
  {
    title: "Developer-First API",
    description:
      "Integrate link generation programmatically into your existing infrastructure. Backed by a 99.99% uptime SLA.",
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
        <polyline points="16 18 22 12 16 6" />
        <polyline points="8 6 2 12 8 18" />
      </svg>
    ),
  },
  {
    title: "Actionable Analytics",
    description:
      "Make data-driven decisions with rich, privacy-respecting insights into geographic, device, and referrer data.",
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
];

export function FeaturesSection() {
  return (
    <>
      {/* Features Section */}
      <section className="py-20 sm:py-28 bg-muted/10">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-3xl font-bold tracking-tight sm:text-4xl">
              Engineered for{" "}
              <span className="text-primary">Enterprise Scale</span>
            </h2>
            <p className="mt-4 max-w-2xl mx-auto text-lg text-muted-foreground">
              Robust infrastructure, granular controls, and seamless integration
              capabilities to power your organization's link management.
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
              Ready to scale your link infrastructure?
            </h2>
            <p className="mt-4 text-lg text-muted-foreground">
              Contact our sales team or start exploring the API today.
            </p>
            <div className="mt-8 flex justify-center gap-4">
              <AuthModal defaultTab="register">
                <Button
                  size="lg"
                  className="rounded-xl px-8 font-semibold text-base"
                >
                  Get Started
                </Button>
              </AuthModal>
              <Button
                variant="outline"
                size="lg"
                className="rounded-xl px-8 font-semibold text-base"
              >
                Contact Sales
              </Button>
            </div>
          </div>
        </div>
      </section>
    </>
  );
}
