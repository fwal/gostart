package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"net/http"
)

func httpHandler() http.Handler {
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
