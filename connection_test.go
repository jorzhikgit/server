package main

import (
	"github.com/boyle-bingo/server/events"
	//"github.com/smartystreets/goconvey/convey"
)

type fakeConnection struct{}

func (fc *fakeConnection) Read() error                    { return nil }
func (fc *fakeConnection) Write(Event events.Event) error { return nil }
func (fc *fakeConnection) Close() error                   { return nil }
