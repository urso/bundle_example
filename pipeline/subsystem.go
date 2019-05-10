package pipeline

import (
	"github.com/eapache/queue"
	"github.com/urso/bundle_example/ctx"
	"github.com/urso/bundle_example/feature"
	"github.com/urso/bundle_example/output"
	"github.com/urso/bundle_example/processor"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/publisher/pipeline"
)

type Subsystem struct {
	Outputs outputLoader
	Queue   queueLoader
}

type loader interface {
	ConfigSchema() feature.Schema
	TestConfig(ctx ctx.ProcessCtx, name string, config *common.Config) error
}

type processorLoader interface {
	loader
	LoadProcessors(ctx ctx.ProcessCtx, configs []*common.Config) (*processor.Group, error)
}

type outputLoader interface {
	loader
	LoadOutput(ctx ctx.ProcessCtx, name string, config *common.Config) ([]output.Instance, error)
}

type queueLoader interface {
	loader
	LoadQueue(ctx ctx.ProcessCtx, name string, config *common.Config) (queue.Queue, error)
}

func (s *Subsystem) Load(ctx ctx.ProcessCtx, config Config) (pipeline.Pipeline, error) {
	panic("...")
}
