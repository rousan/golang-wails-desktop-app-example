package main

import "context"

type binds struct {
	ctx context.Context
}

func newBinds() *binds {
	return &binds{
		ctx: nil,
	}
}

func (b *binds) setContext(ctx context.Context) {
	b.ctx = ctx
}

func (b *binds) Sum(x, y int64) int64 {
	return x + y
}
