package metrics

func GetSystemMetrics() (SystemMetrics, error) {
	return FetchMetricsFromDB()
}
