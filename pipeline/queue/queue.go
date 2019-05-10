package queue

import (
	"github.com/urso/bundle_example/ctx"
	"github.com/urso/bundle_example/feature"
	"github.com/urso/bundle_example/logs"

	"github.com/elastic/beats/libbeat/common"
)

type Loader []QueueFeature

type QueueFeature interface {
	Info() feature.Info
	TestConfig(Context, *common.Config) error
	Create(Context, *common.Config) (Queue, error)
}

type Context struct {
	Done       <-chan struct{}
	Log        *logs.Logger
	Oberserver Observer
}

type Observer interface {
	// methods to oberserve state update and collect metrics
}

type Queue interface {
	// broker queue methods...
}

func (l Loader) LoadQueue(ctx *ctx.ProcessCtx, name string, config *common.Config) (Queue, error) {
	f := l.Find(name)
	if f == nil {
		// return error ...
	}

	var ob Observer // init
	return f.Create(convContext(ctx, ob), config)
}

func (l Loader) TestConfig(ctx ctx.ProcessCtx, name string, config *common.Config) error {
	f := l.Find(name)
	if name == nil {
		// return error ...
	}

	var ob Observer // init
	f.TestConfig(convContext(ctx, ob), config)
}

func (l Loader) Find(name string) QueueFeature {
	for _, f := range l {
		info := f.Info()
		if info.Name == name {
			return f
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
