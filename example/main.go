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

type AuthUser struct {
	flute.Middleware
}

func (auth_user *AuthUser) Before() bool {
	auth_user.RenderText(401, "没有权限")
	return false
}

func main() {
	e := flute.LoadConfig("./app_config.conf")
	if e != nil {
		panic(e)
	}
	flute.HandleFunc("/", "GET", func(c *flute.Context) {
		c.RenderText(200, "hello word")
	})
	var users Users
	flute.Resources("users", &users)
	e = flute.Start()
	if e != nil {
		panic(e)
	}
}
