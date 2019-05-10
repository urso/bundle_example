package output

import (
	"github.com/urso/bundle_example/ctx"
	"github.com/urso/bundle_example/feature"
	"github.com/urso/bundle_example/logs"

	"github.com/elastic/beats/libbeat/common"
)

type Loader []Output

type Output interface {
	Info() feature.Info
	TestConfig(*Context, *common.Config) error
	Create(*Context, *common.Config) ([]Instance, error)
}

type Instance interface {
	// ...
}

type Context struct {
	Done       <-chan struct{}
	Log        *logs.Logger
	Oberserver Observer
}

type Observer interface {
	// methods to oberserve state update and collect metrics
}

func (l Loader) LoadOutput(ctx *ctx.ProcessCtx, name string, config *common.Config) ([]Instance, error) {
	output := l.Find(name)
	if output == nil {
		// return error...
	}

	var ob Observer // init
	return output.Create(convContext(ctx, ob), config)
}

func (l Loader) ConfigSchema() feature.Schema {
	schemas := make([]feature.Schema, len(l))
	names := make([]string, len(l))
	for i, o := range l {
		info := o.Info()
		names[i] = info.Name
		schemas[i] = feature.Namespaced(info.Name, info.Schema)
	}

	return feature.OneOf(names, schemas)
}

func (l Loader) TestConfig(ctx ctx.ProcessCtx, name string, config *common.Config) error {
	output := l.Find(name)
	if name == nil {
		// return error ...
	}

	var ob Observer // init
	output.TestConfig(convContext(ctx, ob), config)
}

func (l Loader) Find(name string) Output {
	for _, o := range l {
		info := o.Info()
		if info.Name == name {
			return o
		}
	}
	return nil
}

func convContext(ctx ctx.ProcessCtx, ob Observer) Context {
	return Context{
		Done:       ctx.Done,
		Log:        ctx.Log,
		Oberserver: ob,
	}
}
