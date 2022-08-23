package main

import (
	"context"
)

type app struct {
	ctx   context.Context
	binds *binds
}

func newApp() *app {
	return &app{
		ctx:   nil,
		binds: newBinds(),
	}
}

func (a *app) setContext(ctx context.Context) {
	a.ctx = ctx
	a.binds.setContext(ctx)
}

func (a *app) startup(ctx context.Context) {
	a.setContext(ctx)
}
