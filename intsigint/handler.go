package main

import "context"

type Handler interface {
	Do(context.Context) error
}

type HandleFunc func(context.Context) error

func (fn HandleFunc) Do(ctx context.Context) error {
	return fn(ctx)
}

type InterceptFunc func(Handler) Handler

func run(h Handler, ctx context.Context) error {
	return h.Do(ctx)
}
