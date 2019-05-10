package ctx

import "github.com/urso/bundle_example/logs"

type ProcessCtx struct {
	Done <-chan struct{}
	Log  *logs.Logger

	// ...
}
