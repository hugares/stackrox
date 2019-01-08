package all

import (
	"sync"

	"github.com/stackrox/rox/central/sensor/service/pipeline"
)

var (
	once sync.Once

	factory pipeline.Factory
)

// Singleton provides the factory that creates pipelines per cluster.
func Singleton() pipeline.Factory {
	once.Do(func() {
		factory = NewFactory()
	})
	return factory
}
