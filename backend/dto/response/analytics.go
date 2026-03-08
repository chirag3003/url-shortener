package response

// AnalyticsSummaryResponse provides top-card analytics metrics.
type AnalyticsSummaryResponse struct {
	TotalClicks           int64   `json:"totalClicks"`
	TotalClicksChange     float64 `json:"totalClicksChange"`
	UniqueVisitors        int64   `json:"uniqueVisitors"`
	UniqueVisitorsChange  float64 `json:"uniqueVisitorsChange"`
	ClicksLast24H         int64   `json:"clicksLast24h"`
	ClicksLast24HChange   float64 `json:"clicksLast24hChange"`
	ClicksLast7Days       int64   `json:"clicksLast7d"`
	ClicksLast7DaysChange float64 `json:"clicksLast7dChange"`
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
