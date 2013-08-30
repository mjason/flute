package main

import (
	"flute"
)

type Users struct {
	flute.Controller
}

func (u *Users) Index() {
	u.RenderText(200, "mj")
}

func (u *Users) Show() {
	u.RenderText(200, "mj"+u.GetVars("id"))
}

func main() {
	flute.HandleFunc("/", "GET", func(c *flute.Context) {
		c.RenderText(200, "hello word")
	})
	var users Users
	flute.Resources("users", &users)
	flute.Start(":1234")
}
