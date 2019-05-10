package beatinst

import (
	"github.com/urso/bundle_example/ctx"

	"github.com/elastic/beats/filebeat/input"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/publisher/pipeline"
)

type Instance struct {
	Inputs          inputsSubsystem
	Pipeline        pipelineSubsystem
	StackManagement stackmgmtSubsystem
}

type pipelineSubsystem interface {
	Load(ctx ctx.ProcessCtx, config pipeline.Config) (pipeline.Pipeline, error)
}

type stackmgmtSubsystem interface {
	// ... similar to beats/libbeat/idxmgmt?
}

// for loading inputs from config, via config reloading and auto-discovery
type inputsSubsystem interface {
	Load(ctx ctx.ProcessCtx, stack stackmgmt.Manager, name string, config *common.Config) (input.Input, error)
}

func Run(inst *beatinst) error {
	panic("...")
}
