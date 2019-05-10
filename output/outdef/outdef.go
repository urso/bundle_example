package outdef

import (
	"github.com/urso/bundle_example/feature"
	"github.com/urso/bundle_example/output"

	"github.com/elastic/beats/libbeat/common"
)

type Feature struct {
	Info feature.Info
	Test func(*output.Context, *common.Config) error
	New  func(*output.Context, *common.Config) ([]Instance, error)
}

func (f *Feature) Info() feature.Info { return f.Info }
func (f *Feature) TestConfig(ctx *output.Context, cfg *common.Config) error {
	return f.Test(ctx, cfg)
}
func (f *Feature) Create(ctx *output.Context, cfg *common.Config) error {
	return f.New(ctx, cfg)
}
