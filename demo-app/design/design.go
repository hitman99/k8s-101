package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("demo", func() {
	Title("demo-app")
	Contact(func() {
		Name("Tomas A.")
		Email("tomas@adomavicius.com")
	})
	Description("Kubernetes 101 Demo application")
	Scheme("http")
	Host("localhost")
})
