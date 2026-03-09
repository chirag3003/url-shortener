"use client";

import { useParams } from "next/navigation";
import { useLink } from "@/hooks/use-links";
import {
  useAnalyticsSummary,
  useAnalyticsTimeSeries,
} from "@/hooks/use-analytics";
import { Breakdowns } from "./_components/breakdowns";
import { ClicksChart } from "./_components/clicks-chart";
import { SummaryCards } from "./_components/summary-cards";
import { Skeleton } from "@/components/ui/skeleton";

export default function AnalyticsPage() {
  const { id } = useParams<{ id: string }>();

  const { data: link, isLoading: isLoadingLink } = useLink(id);
  const { data: summary, isLoading: isLoadingSummary } =
    useAnalyticsSummary(id);

  const { data: ts24h, isLoading: isTs24h } = useAnalyticsTimeSeries(id, "24h");
  const { data: ts7d, isLoading: isTs7d } = useAnalyticsTimeSeries(id, "7d");
  const { data: ts30d, isLoading: isTs30d } = useAnalyticsTimeSeries(id, "30d");

  const timeSeriesData = {
    "24h": ts24h ?? [],
    "7d": ts7d ?? [],
    "30d": ts30d ?? [],
  };

  const summaryCards = [
    {
      title: "Total Clicks",
      value: summary?.totalClicks.toLocaleString() ?? "...",
      change: "All time",
      positive: true,
    },
    {
      title: "Unique Visitors",
      value: summary?.uniqueVisitors.toLocaleString() ?? "...",
      change: "Unique IPs",
      positive: true,
    },
    {
      title: "Last 24 Hours",
      value: summary?.clicksLast24h.toLocaleString() ?? "...",
      change: "Recent",
      positive: true,
    },
    {
      title: "Last 7 Days",
      value: summary?.clicksLast7d.toLocaleString() ?? "...",
      change: "Weekly",
      positive: true,
    },
  ];

  return (
    <div className="space-y-8">
      {/* Header */}
      <div>
        <h1 className="text-2xl font-bold tracking-tight sm:text-3xl">
          Analytics
        </h1>
        {isLoadingLink ? (
          <Skeleton className="mt-2 h-6 w-64" />
        ) : link ? (
          <div className="mt-1 flex items-center gap-2 text-muted-foreground">
            <span className="font-mono text-primary font-semibold">
              /{link.shortCode}
            </span>
            <span>→</span>
            <span className="truncate text-sm max-w-md">{link.longUrl}</span>
          </div>
        ) : (
          <div className="mt-1 text-sm text-destructive">Link not found</div>
        )}
      </div>

      <SummaryCards cards={summaryCards} />
      <ClicksChart timeSeriesData={timeSeriesData} />
      <Breakdowns id={id} />
    </div>
  );
}
