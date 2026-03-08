import { DonutChart } from "@/components/charts/donut-chart";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import {
  mockBrowsers,
  mockDevices,
  mockGeography,
  mockReferrers,
} from "@/lib/mock-data";

// Country flag emoji helper
const countryFlags: Record<string, string> = {
  "United States": "🇺🇸",
  India: "🇮🇳",
  "United Kingdom": "🇬🇧",
  Germany: "🇩🇪",
  Canada: "🇨🇦",
  France: "🇫🇷",
  Japan: "🇯🇵",
  Australia: "🇦🇺",
  Brazil: "🇧🇷",
  Netherlands: "🇳🇱",
};

const referrerFavicons: Record<string, string> = {
  "twitter.com": "🐦",
  "google.com": "🔍",
  Direct: "🔗",
  "reddit.com": "🤖",
  "linkedin.com": "💼",
  "github.com": "🐙",
  "facebook.com": "👤",
};

export function Breakdowns() {
  return (
    <div className="grid gap-6 lg:grid-cols-2">
      {/* Referrers */}
      <Card>
        <CardHeader>
          <CardTitle className="text-base">Top Referrers</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="space-y-3">
            {mockReferrers.map((item) => {
              const maxCount = mockReferrers[0].count;
              const pct = (item.count / maxCount) * 100;
              return (
                <div key={item.key} className="space-y-1">
                  <div className="flex items-center justify-between text-sm">
                    <span className="flex items-center gap-2">
                      <span>{referrerFavicons[item.key] ?? "🌐"}</span>
                      <span>{item.key}</span>
                    </span>
                    <span className="font-semibold tabular-nums">
                      {item.count.toLocaleString()}
                    </span>
                  </div>
                  <div className="h-1.5 w-full rounded-full bg-muted overflow-hidden">
                    <div
                      className="h-full rounded-full bg-primary/70 transition-all duration-500"
                      style={{ width: `${pct}%` }}
                    />
                  </div>
                </div>
              );
            })}
          </div>
        </CardContent>
      </Card>

      {/* Geography */}
      <Card>
        <CardHeader>
          <CardTitle className="text-base">Geography</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="space-y-3">
            {mockGeography.map((item) => {
              const maxCount = mockGeography[0].count;
              const pct = (item.count / maxCount) * 100;
              return (
                <div key={item.key} className="space-y-1">
                  <div className="flex items-center justify-between text-sm">
                    <span className="flex items-center gap-2">
                      <span>{countryFlags[item.key] ?? "🌍"}</span>
                      <span>{item.key}</span>
                    </span>
                    <span className="font-semibold tabular-nums">
                      {item.count.toLocaleString()}
                    </span>
                  </div>
                  <div className="h-1.5 w-full rounded-full bg-muted overflow-hidden">
                    <div
                      className="h-full rounded-full bg-accent/70 transition-all duration-500"
                      style={{ width: `${pct}%` }}
                    />
                  </div>
                </div>
              );
            })}
          </div>
        </CardContent>
      </Card>

      {/* Devices */}
      <Card>
        <CardHeader>
          <CardTitle className="text-base">Devices</CardTitle>
        </CardHeader>
        <CardContent>
          <DonutChart data={mockDevices} />
        </CardContent>
      </Card>

      {/* Browsers */}
      <Card>
        <CardHeader>
          <CardTitle className="text-base">Browsers</CardTitle>
        </CardHeader>
        <CardContent>
          <DonutChart data={mockBrowsers} />
        </CardContent>
      </Card>
    </div>
  );
}
