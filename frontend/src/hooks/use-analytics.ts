import { useQueries, useQuery } from "@tanstack/react-query";
import {
  type AnalyticsKind,
  analyticsService,
} from "@/services/analytics.service";

export function useAnalyticsSummary(id: string) {
  return useQuery({
    queryKey: ["analytics", "summary", id],
    queryFn: () => analyticsService.getSummary(id),
    enabled: !!id,
  });
}

export function useAnalyticsTimeSeries(
  id: string,
  range: "24h" | "7d" | "30d" = "7d",
) {
  return useQuery({
    queryKey: ["analytics", "timeseries", id, range],
    queryFn: () => analyticsService.getTimeSeries(id, { range }),
    enabled: !!id,
  });
}

export function useAnalyticsBreakdown(id: string, kind: AnalyticsKind) {
  return useQuery({
    queryKey: ["analytics", "breakdown", id, kind],
    queryFn: () => analyticsService.getBreakdown(id, kind),
    enabled: !!id,
  });
}

export function useFullLinkAnalytics(id: string) {
  const kinds: AnalyticsKind[] = [
    "referrers",
    "devices",
    "browsers",
    "geography",
  ];

  return useQueries({
    queries: kinds.map((kind) => ({
      queryKey: ["analytics", "breakdown", id, kind],
      queryFn: () => analyticsService.getBreakdown(id, kind),
      enabled: !!id,
    })),
  });
}
