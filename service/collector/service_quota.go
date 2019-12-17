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
		prometheus.BuildFQName("todo_operator", "todo", "info"),
		"Todo description of the todo operator todo metric",
		[]string{
			labelInstallation,
			labelClusterID,
		},
		nil,
	)
)

type ServiceCollectorConfig struct {
}

type ServiceCollector struct {
}

func NewTodo(config ServiceCollectorConfig) (*ServiceCollector, error) {
	r := &ServiceCollector{}

	return r, nil
}

func (r *ServiceCollector) Collect(ch chan<- prometheus.Metric) error {
	return nil
}

func (r *ServiceCollector) Describe(ch chan<- *prometheus.Desc) error {
	ch <- ScheduleDesc

	return nil
}
