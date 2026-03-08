package response

// AnalyticsSummaryResponse provides top-card analytics metrics.
type AnalyticsSummaryResponse struct {
	TotalClicks     int64 `json:"totalClicks"`
	UniqueVisitors  int64 `json:"uniqueVisitors"`
	ClicksLast24H   int64 `json:"clicksLast24h"`
	ClicksLast7Days int64 `json:"clicksLast7d"`
}

// AnalyticsPoint represents one time bucket.
type AnalyticsPoint struct {
	Bucket string `json:"bucket"`
	Clicks int64  `json:"clicks"`
}

// BreakdownItem is used for referrer/device/browser/country responses.
type BreakdownItem struct {
	Key   string `json:"key"`
	Count int64  `json:"count"`
}
