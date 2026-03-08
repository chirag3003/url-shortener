import type {
  AnalyticsPoint,
  AnalyticsSummary,
  BreakdownItem,
} from "@/lib/validators/analytics";
import type { LinkResponse } from "@/lib/validators/link";
import type { UserResponse } from "@/lib/validators/user";

// ── Mock User ────────────────────────────────────────────────────────────

export const mockUser: UserResponse = {
  id: "7892143600128",
  name: "Chirag Bhalotia",
  email: "chirag@example.com",
};

// ── Mock Links ───────────────────────────────────────────────────────────

export const mockLinks: LinkResponse[] = [
  {
    id: "1001",
    userId: "7892143600128",
    longUrl:
      "https://www.figma.com/design/abcdef1234/url-shortener-design-system?node-id=0-1&t=xyz",
    shortCode: "figma-ds",
    shortUrl: "http://localhost:5000/figma-ds",
    redirectType: 302,
    isActive: true,
    createdAt: "2026-03-01T10:30:00Z",
    updatedAt: "2026-03-01T10:30:00Z",
  },
  {
    id: "1002",
    userId: "7892143600128",
    longUrl:
      "https://docs.google.com/document/d/1a2b3c4d5e6f/edit?usp=sharing&tab=overview",
    shortCode: "gdocs-rfc",
    shortUrl: "http://localhost:5000/gdocs-rfc",
    redirectType: 301,
    isActive: true,
    createdAt: "2026-02-28T14:00:00Z",
    updatedAt: "2026-03-05T09:15:00Z",
  },
  {
    id: "1003",
    userId: "7892143600128",
    longUrl: "https://github.com/chirag3003/go-backend-template/pull/42",
    shortCode: "pr42",
    shortUrl: "http://localhost:5000/pr42",
    redirectType: 302,
    isActive: true,
    createdAt: "2026-02-25T08:45:00Z",
    updatedAt: "2026-02-25T08:45:00Z",
  },
  {
    id: "1004",
    userId: "7892143600128",
    longUrl:
      "https://vercel.com/chirag3003/url-shortener/deployments?status=READY",
    shortCode: "deploys",
    shortUrl: "http://localhost:5000/deploys",
    redirectType: 302,
    expiresAt: "2026-04-01T00:00:00Z",
    isActive: true,
    createdAt: "2026-02-20T17:30:00Z",
    updatedAt: "2026-02-20T17:30:00Z",
  },
  {
    id: "1005",
    userId: "7892143600128",
    longUrl:
      "https://medium.com/@chirag3003/building-a-production-url-shortener-with-go-fiber",
    shortCode: "blog-go",
    shortUrl: "http://localhost:5000/blog-go",
    redirectType: 301,
    isActive: false,
    createdAt: "2026-02-15T12:00:00Z",
    updatedAt: "2026-03-02T16:40:00Z",
  },
  {
    id: "1006",
    userId: "7892143600128",
    longUrl: "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
    shortCode: "yt-demo",
    shortUrl: "http://localhost:5000/yt-demo",
    redirectType: 302,
    isActive: true,
    createdAt: "2026-02-10T09:00:00Z",
    updatedAt: "2026-02-10T09:00:00Z",
  },
];

// ── Click Counts (keyed by link ID) ──────────────────────────────────────

export const mockClickCounts: Record<string, number> = {
  "1001": 1247,
  "1002": 834,
  "1003": 456,
  "1004": 2901,
  "1005": 123,
  "1006": 15802,
};

// ── Mock Analytics Summary ───────────────────────────────────────────────

export const mockAnalyticsSummary: AnalyticsSummary = {
  totalClicks: 15802,
  uniqueVisitors: 9431,
  clicksLast24h: 342,
  clicksLast7d: 2184,
};

// ── Mock Time-Series Data ────────────────────────────────────────────────

function generateTimeSeries(days: number, scale: number): AnalyticsPoint[] {
  const points: AnalyticsPoint[] = [];
  const now = new Date();
  for (let i = days - 1; i >= 0; i--) {
    const date = new Date(now);
    date.setDate(date.getDate() - i);
    points.push({
      bucket: date.toISOString(),
      clicks: Math.floor(Math.random() * scale + scale * 0.3),
    });
  }
  return points;
}

export const mockTimeSeries30d: AnalyticsPoint[] = generateTimeSeries(30, 120);
export const mockTimeSeries7d: AnalyticsPoint[] = generateTimeSeries(7, 300);
export const mockTimeSeries24h: AnalyticsPoint[] = Array.from(
  { length: 24 },
  (_, i) => {
    const date = new Date();
    date.setHours(date.getHours() - (23 - i), 0, 0, 0);
    return {
      bucket: date.toISOString(),
      clicks: Math.floor(Math.random() * 50 + 5),
    };
  },
);

// ── Mock Breakdowns ──────────────────────────────────────────────────────

export const mockReferrers: BreakdownItem[] = [
  { key: "twitter.com", count: 4210 },
  { key: "google.com", count: 3890 },
  { key: "Direct", count: 3100 },
  { key: "reddit.com", count: 1820 },
  { key: "linkedin.com", count: 1450 },
  { key: "github.com", count: 920 },
  { key: "facebook.com", count: 412 },
];

export const mockDevices: BreakdownItem[] = [
  { key: "Desktop", count: 9200 },
  { key: "Mobile", count: 5100 },
  { key: "Tablet", count: 1502 },
];

export const mockBrowsers: BreakdownItem[] = [
  { key: "Chrome", count: 8400 },
  { key: "Safari", count: 3200 },
  { key: "Firefox", count: 2100 },
  { key: "Edge", count: 1400 },
  { key: "Other", count: 702 },
];

export const mockGeography: BreakdownItem[] = [
  { key: "United States", count: 5400 },
  { key: "India", count: 3800 },
  { key: "United Kingdom", count: 1900 },
  { key: "Germany", count: 1200 },
  { key: "Canada", count: 980 },
  { key: "France", count: 760 },
  { key: "Japan", count: 540 },
  { key: "Australia", count: 420 },
  { key: "Brazil", count: 380 },
  { key: "Netherlands", count: 222 },
];
