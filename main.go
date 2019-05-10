package main

import (
	"fmt"

	"github.com/urso/bundle_example/beatinst"
	"github.com/urso/bundle_example/output"
	"github.com/urso/bundle_example/output/console"
	"github.com/urso/bundle_example/output/elasticsearch"
	"github.com/urso/bundle_example/pipeline"
	"github.com/urso/bundle_example/pipeline/queue"
	"github.com/urso/bundle_example/processor"

	...
)

func main() {

	processors := processor.Loader(
		dropevent.Processor,
		dropfields.Processor,
		dissect.Processor,
	)

	inst := beatinst.Instance{
		Inputs: inputs.MakeLoader(processors,
			container.Input,
			docker.Input,
			syslog.Input,
			logs.Input,
		),
		StackManagement: idxmgmt.NewSupport(
			...
		),
		Pipeline: &pipeline.Subsystem{
			Processors: processors,
			Queue: queue.Loader(
				memqueue.Queue,
			),
			Outputs: output.Loader(
				console.Output,
				elasticsearch.Output,
			),
		},
	}

	err := beatinst.Run(inst)
	if err != nil {
		...
	}
}
