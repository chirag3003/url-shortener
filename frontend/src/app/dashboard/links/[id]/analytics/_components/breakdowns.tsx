import { DonutChart } from "@/components/charts/donut-chart";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { useFullLinkAnalytics } from "@/hooks/use-analytics";
import { Skeleton } from "@/components/ui/skeleton";

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

export function Breakdowns({ id }: { id: string }) {
  const results = useFullLinkAnalytics(id);

  const [referrers, devices, browsers, geography] = results.map(
    (r) => r.data ?? [],
  );
  const isLoading = results.some((r) => r.isLoading);

  const renderSkeleton = () => (
    <div className="space-y-4">
      {[1, 2, 3, 4].map((i) => (
        <div key={i} className="space-y-2">
          <Skeleton className="h-4 w-full" />
          <Skeleton className="h-2 w-full" />
        </div>
      ))}
    </div>
  );

  return (
    <div className="grid gap-6 lg:grid-cols-2">
      {/* Referrers */}
      <Card>
        <CardHeader>
          <CardTitle className="text-base">Top Referrers</CardTitle>
        </CardHeader>
        <CardContent>
          {isLoading ? (
            renderSkeleton()
          ) : referrers.length === 0 ? (
            <div className="py-8 text-center text-sm text-muted-foreground">
              No data yet
            </div>
          ) : (
            <div className="space-y-3">
              {referrers.map((item, idx) => {
                const maxCount = referrers[0].count;
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
          )}
        </CardContent>
      </Card>

      {/* Geography */}
      <Card>
        <CardHeader>
          <CardTitle className="text-base">Geography</CardTitle>
        </CardHeader>
        <CardContent>
          {isLoading ? (
            renderSkeleton()
          ) : geography.length === 0 ? (
            <div className="py-8 text-center text-sm text-muted-foreground">
              No data yet
            </div>
          ) : (
            <div className="space-y-3">
              {geography.map((item) => {
                const maxCount = geography[0].count;
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
          )}
        </CardContent>
      </Card>

      {/* Devices */}
      <Card>
        <CardHeader>
          <CardTitle className="text-base">Devices</CardTitle>
        </CardHeader>
        <CardContent>
          {isLoading ? (
            <div className="flex items-center justify-center h-[200px]">
              <Skeleton className="h-32 w-32 rounded-full" />
            </div>
          ) : devices.length === 0 ? (
            <div className="py-8 text-center text-sm text-muted-foreground">
              No data yet
            </div>
          ) : (
            <DonutChart data={devices} />
          )}
        </CardContent>
      </Card>

      {/* Browsers */}
      <Card>
        <CardHeader>
          <CardTitle className="text-base">Browsers</CardTitle>
        </CardHeader>
        <CardContent>
          {isLoading ? (
            <div className="flex items-center justify-center h-[200px]">
              <Skeleton className="h-32 w-32 rounded-full" />
            </div>
          ) : browsers.length === 0 ? (
            <div className="py-8 text-center text-sm text-muted-foreground">
              No data yet
            </div>
          ) : (
            <DonutChart data={browsers} />
          )}
        </CardContent>
      </Card>
    </div>
  );
}
