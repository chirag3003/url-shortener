import apiClient from "@/lib/api/client";
import type {
  AnalyticsPoint,
  AnalyticsSummary,
  BreakdownItem,
} from "@/lib/validators/analytics";

export type AnalyticsKind = "referrers" | "devices" | "browsers" | "geography";

export const analyticsService = {
  async getSummary(id: string): Promise<AnalyticsSummary> {
    const response = await apiClient.get<AnalyticsSummary>(
      `/api/v1/links/${id}/analytics/summary`
    );
    return response.data;
  },

  async getTimeSeries(
    id: string,
    query?: { range?: "24h" | "7d" | "30d" }
  ): Promise<AnalyticsPoint[]> {
    const response = await apiClient.get<AnalyticsPoint[]>(
      `/api/v1/links/${id}/analytics/timeseries`,
      { params: query }
    );
    return response.data;
  },

  async getBreakdown(
    id: string,
    kind: AnalyticsKind
  ): Promise<BreakdownItem[]> {
    const response = await apiClient.get<BreakdownItem[]>(
      `/api/v1/links/${id}/analytics/${kind}`
    );
    return response.data;
  },
};
