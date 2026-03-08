import { z } from "zod/v4";

export const analyticsSummarySchema = z.object({
  totalClicks: z.number(),
  uniqueVisitors: z.number(),
  clicksLast24h: z.number(),
  clicksLast7d: z.number(),
});

export const analyticsPointSchema = z.object({
  bucket: z.string(),
  clicks: z.number(),
});

export const breakdownItemSchema = z.object({
  key: z.string(),
  count: z.number(),
});

export type AnalyticsSummary = z.infer<typeof analyticsSummarySchema>;
export type AnalyticsPoint = z.infer<typeof analyticsPointSchema>;
export type BreakdownItem = z.infer<typeof breakdownItemSchema>;
