package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	labelInstallation = "installation"
	labelClusterID    = "cluster_id"
)

var (
	ScheduleDesc *prometheus.Desc = prometheus.NewDesc(
		prometheus.BuildFQName("service_quota_operator", "service_quota", "info"),
		"Todo description of the todo operator todo metric",
		[]string{
			labelInstallation,
			labelClusterID,
		},
		nil,
	)
)

type ServiceQuotaCollectorConfig struct {
}

type ServiceQuotaCollector struct {
}

func NewServiceQuota(config ServiceQuotaCollectorConfig) (*ServiceQuotaCollector, error) {
	r := &ServiceQuotaCollector{}

	return r, nil
}

func (r *ServiceQuotaCollector) Collect(ch chan<- prometheus.Metric) error {
	return nil
}

func (r *ServiceQuotaCollector) Describe(ch chan<- *prometheus.Desc) error {
	ch <- ScheduleDesc

	return nil
}
