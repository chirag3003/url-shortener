"use client";

import {
  mockAnalyticsSummary,
  mockLinks,
  mockTimeSeries7d,
  mockTimeSeries24h,
  mockTimeSeries30d,
} from "@/lib/mock-data";
import { Breakdowns } from "./_components/breakdowns";
import { ClicksChart } from "./_components/clicks-chart";
import { SummaryCards } from "./_components/summary-cards";

export default function AnalyticsPage() {
  const timeSeriesData = {
    "24h": mockTimeSeries24h,
    "7d": mockTimeSeries7d,
    "30d": mockTimeSeries30d,
  };

  // Find the link for this analytics page (using first link as mock)
  const link = mockLinks[0];
  const summary = mockAnalyticsSummary;

  const summaryCards = [
    {
      title: "Total Clicks",
      value: summary.totalClicks.toLocaleString(),
      change: "+12%",
      positive: true,
    },
    {
      title: "Unique Visitors",
      value: summary.uniqueVisitors.toLocaleString(),
      change: "+8%",
      positive: true,
    },
    {
      title: "Last 24 Hours",
      value: summary.clicksLast24h.toLocaleString(),
      change: "-3%",
      positive: false,
    },
    {
      title: "Last 7 Days",
      value: summary.clicksLast7d.toLocaleString(),
      change: "+15%",
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
        {link && (
          <div className="mt-1 flex items-center gap-2 text-muted-foreground">
            <span className="font-mono text-primary font-semibold">
              /{link.shortCode}
            </span>
            <span>→</span>
            <span className="truncate text-sm max-w-md">{link.longUrl}</span>
          </div>
        )}
      </div>

      <SummaryCards cards={summaryCards} />
      <ClicksChart timeSeriesData={timeSeriesData} />
      <Breakdowns />
    </div>
  );
}
