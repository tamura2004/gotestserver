package main

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()
	m.RunOnAddr(":80")
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}
