package main

import (
	"changeme/constants"
	"changeme/state"
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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

func (a *app) onStartup(ctx context.Context) {
	a.setContext(ctx)
	state.Subscribe(a.onStateChanged)
	a.registerListeners()
}

func (a *app) onShutdown(ctx context.Context) {
	fmt.Println("App shutdown")
}

func (a *app) onStateChanged(newState *state.AppState) error {
	runtime.LogInfo(a.ctx, fmt.Sprint(newState))
	runtime.EventsEmit(a.ctx, constants.EventNameBackendStateChanged, newState)
	return nil
}

func (a *app) registerListeners() {
	runtime.EventsOn(a.ctx, constants.EventNameFrontendReady, func(optionalData ...interface{}) {
		runtime.EventsEmit(a.ctx, constants.EventNameBackendStateChanged, state.GetState())
	})
}
