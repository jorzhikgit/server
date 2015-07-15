package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRegisterWithHub(t *testing.T) {
	Convey("Given a new hub and user", t, func() {
		h := NewHub()
		u := NewUser(FakeConnection{}, 0, Player{})

		go h.Run()

		Convey("Pass the user to the registration channel", func() {
			h.register <- u
		})
	})
}
