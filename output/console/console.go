package console

import (
	"github.com/urso/bundle_example/feature"
	"github.com/urso/bundle_example/output"
	"github.com/urso/bundle_example/output/outdef"

	"github.com/elastic/beats/libbeat/common"
)

var Output = outdef.Feature{
	Info: feature.Info{},
	Test: testOutput,
	New:  newOutput,
}

func testOutput(ctx *output.Context, cfg *common.Config) error {
	return nil
}

func newOutput(ctx *output.Config, cfg *common.Config) ([]output.Instance, error) {
	return nil
}
