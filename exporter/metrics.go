package exporter

import "github.com/prometheus/client_golang/prometheus"

// AddMetrics - Add's all of the metrics to a map of strings, returns the map.
func AddMetrics() map[string]*prometheus.Desc {

	APIMetrics := make(map[string]*prometheus.Desc)

	APIMetrics["OpenIssues"] = prometheus.NewDesc(
		prometheus.BuildFQName("github", "repo", "open_issues"),
		"Total number of open issues for given repository",
		[]string{"repo", "user"}, nil,
	)
	APIMetrics["PullRequestCount"] = prometheus.NewDesc(
		prometheus.BuildFQName("github", "repo", "pull_request_count"),
		"Total number of pull requests for given repository",
		[]string{"repo"}, nil,
	)

	return APIMetrics
}

// processMetrics - processes the response data and sets the metrics using it as a source
func (e *Exporter) processMetrics(data []*Datum, ch chan<- prometheus.Metric) error {

	// APIMetrics - range through the data slice
	for _, x := range data {
		prCount := 0
		for range x.Pulls {
			prCount += 1
		}
		// issueCount = x.OpenIssue - prCount
		ch <- prometheus.MustNewConstMetric(e.APIMetrics["OpenIssues"], prometheus.GaugeValue, (x.OpenIssues - float64(prCount)), x.Name, x.Owner.Login)

		// prCount
		ch <- prometheus.MustNewConstMetric(e.APIMetrics["PullRequestCount"], prometheus.GaugeValue, float64(prCount), x.Name)
	}

	return nil
}
