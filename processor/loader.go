package processor

import (
	"github.com/urso/bundle_example/ctx"
	"github.com/urso/bundle_example/feature"
	"github.com/urso/bundle_example/logs"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
)

type Loader []ProcessorFeature

type ProcessorFeature interface {
	Info() feature.Info
	TestConfig(Context, *common.Config) error
	Create(Context, *common.Config) (Processor, error)
}

type Group struct {
	processors []Processor
}

type Processor interface {
	ProcessEvent(event beat.Event) (beat.Event, error)
}

type Context struct {
	Done       <-chan struct{} // for processor with aync processes/caches
	Log        *logs.Logger
	Loader     Loader
	Oberserver Observer
}

type Observer interface {
	// methods to oberserve state update and collect metrics
}

func (l Loader) LoadProcessors(ctx ctx.ProcessCtx, configs []*common.Config) (*ProcessorGroup, error) {
	processors := make([]Processor, 0, len(configs))
	ctxChannels := make([]chan struct{}, 0, len(configs))

	closeContexts := func() {
		for i := range contexts {
			close(ctxChannels[i])
		}
	}

	ok := false
	defer func() {
		if !ok {
			closeContexts()
		}
	}()

	for i, config := range configs {
		done := make(chan struct{})

		var ob Oberserver
		ctx := convContext(ctx, l, ob)
		ctx.Done = done

		processor, err := l.load(ctx, config)
		if err != nil {
			return nil, err
		}

		processors = append(processors, processor)
		ctxChannels = append(ctxChannels, done)
	}

	ok = true
	go func() { // start routine for propagating shutdown to processors
		<-ctx.Done
		closeContexts()
	}()

	return Group{processors}, nil
}

func (l Loader) TestConfig(ctx ctx.ProcessCtx, name string, config *common.Config) error {
	f := l.Find(name)
	if name == nil {
		// return error ...
	}

	var ob Observer // init
	f.TestConfig(convContext(ctx, ob), config)
}

func convContext(ctx ctx.ProcessCtx, loader Loader, ob Observer) Context {
	return Context{
		Done:       ctx.Done,
		Log:        ctx.Log,
		Loader:     loader,
		Oberserver: ob,
	}
}
