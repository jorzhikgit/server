package main

import (
	"github.com/boyle-bingo/server/events"
	//"github.com/smartystreets/goconvey/convey"
)

type FakeConnection struct{}

func (fc *FakeConnection) Read() error                    { return nil }
func (fc *FakeConnection) Write(Event events.Event) error { return nil }
func (fc *FakeConnection) Close() error                   { return nil }
