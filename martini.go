package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

func Martini() *martini.ClassicMartini {
	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Directory:  "static/tmpl",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
	}))

	m.Get("/", func(r render.Render) {
		r.HTML(200, "hello", "World!")
	})

	return m
}
