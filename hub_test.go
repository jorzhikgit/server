package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRegisterWithHub(t *testing.T) {
	Convey("Given a new hub and user", t, func() {
		h := NewHub()
		fc := fakeConnection{}
		u := NewUser(&fc, 0, Player{})

		go h.Run()

		Convey("Pass the user to the registration channel", func() {
			h.register <- u

			Convey("User should be in the notInGame list", func() {
				userNotInGame, ok := h.notInGame[u]
				So(ok, ShouldEqual, true)
				So(userNotInGame, ShouldEqual, true)
			})
		})
	})
}
