package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Directory:  "frontend/html",
		Extensions: []string{".html"}, // Specify extensions to load for templates.
	}))

	m.Use(martini.Static("frontend"))

	m.Get("/", func(r render.Render) {
		r.HTML(200, "App/home", map[string]interface{}{"title": "World of Chris", "active": "home"})
	})

	m.Get("/work", func(r render.Render) {
		r.HTML(200, "App/work", map[string]interface{}{"title": "Chris's work ", "active": "work"})
	})

	m.Get("/words", func(r render.Render) {
		r.HTML(200, "App/words", map[string]interface{}{"title": "Chris's words", "active": "words"})
	})

	m.Get("/links", func(r render.Render) {
		r.HTML(200, "App/links", map[string]interface{}{"title": "Chris's links", "active": "links"})
	})

	m.Get("/contact", func(r render.Render) {
		r.HTML(200, "App/contact", map[string]interface{}{"title": "Chris's info", "active": "contact"})
	})

	m.Get("/resume", func(r render.Render) {
		r.HTML(200, "App/contact", map[string]interface{}{"title": "Chris's info", "active": "work"})
	})

	m.RunOnAddr(":9001")
}
