package metrics

import (
	"reginald_backend/internal/audit"
	"time"
)

type LatencyPoint struct {
	Timestamp time.Time `json:"timestamp"`
	Latency   float64   `json:"latency"`
}

type ReliabilityMetrics struct {
	Maturity       float64 `json:"maturity"`
	FaultTolerance float64 `json:"fault_tolerance"`
	Recoverability float64 `json:"recoverability"`
}

type PerformanceMetrics struct {
	ResourceUtilization float64 `json:"resource_utilization"`
	Capacity            float64 `json:"capacity"`
}

type SystemMetrics struct {
	TotalUsers      int                `json:"total_users"`
	TotalTasks      int                `json:"total_tasks"`
	TotalRequests   int                `json:"total_requests"`
	PendingTasks    int                `json:"pending_tasks"`
	PendingRequests int                `json:"pending_requests"`
	AvgLatency      float64            `json:"avg_latency"`
	Availability    float64            `json:"availability"`
	Reliability     ReliabilityMetrics `json:"reliability"`
	Performance     PerformanceMetrics `json:"performance"`
	LatencyHistory  []LatencyPoint     `json:"latency_history"`
	RecentLogs      []audit.AuditLog   `json:"recent_logs"`
	SecurityAlerts  []audit.AuditLog   `json:"security_alerts"`
}
