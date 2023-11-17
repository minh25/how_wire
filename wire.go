//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitializeEvent(string) (Event, error) {
	wire.Build(NewEvent, NewGreater, NewMessage)
	return Event{}, nil
}
