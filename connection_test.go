package main

import (
//"github.com/smartystreets/goconvey/convey"
)

type FakeConnection struct{}

func (fc *FakeConnection) Read() error  { return nil }
func (fc *FakeConnection) Write() error { return nil }
