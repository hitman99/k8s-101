package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("public", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/", "static/index.html")
	Files("/styles/*filepath", "static/styles")
	Files("/images/*filepath", "static/images")
})
