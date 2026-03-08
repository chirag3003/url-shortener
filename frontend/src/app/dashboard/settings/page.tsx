"use client";

import { useState } from "react";
import { cn } from "@/lib/utils";
import { ProfileView } from "./_components/profile-view";
import { SecurityView } from "./_components/security-view";

// ── Settings Page ────────────────────────────────────────────────────────

const sidebarNavItems = [
  {
    title: "Profile",
    id: "profile",
  },
  {
    title: "Security",
    id: "security",
  },
];

export default function SettingsPage() {
  const [activeTab, setActiveTab] = useState("profile");

  return (
    <div className="space-y-6 pb-16 md:block">
      <div className="space-y-0.5">
        <h2 className="text-2xl font-bold tracking-tight">Settings</h2>
        <p className="text-muted-foreground">
          Manage your account settings and set e-mail preferences.
        </p>
      </div>
      <div className="h-px bg-border/60" />

      <div className="flex flex-col space-y-8 lg:flex-row lg:space-x-12 lg:space-y-0">
        <aside className="lg:w-1/5 shrink-0">
          <nav className="flex space-x-2 lg:flex-col lg:space-x-0 lg:space-y-1 overflow-x-auto pb-2 lg:pb-0">
            {sidebarNavItems.map((item) => (
              // biome-ignore lint/a11y/useButtonType: <later>
              <button
                key={item.id}
                onClick={() => setActiveTab(item.id)}
                className={cn(
                  "inline-flex items-center rounded-md px-3 py-2 text-sm font-medium transition-colors hover:bg-muted hover:text-foreground whitespace-nowrap",
                  activeTab === item.id
                    ? "bg-muted font-semibold text-foreground"
                    : "text-muted-foreground",
                )}
              >
                {item.title}
              </button>
            ))}
          </nav>
        </aside>

        <div className="flex-1 max-w-3xl">
          {activeTab === "profile" && <ProfileView />}
          {activeTab === "security" && <SecurityView />}
        </div>
      </div>
    </div>
  );
}
