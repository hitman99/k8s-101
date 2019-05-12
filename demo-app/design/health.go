package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("health", func() {

	BasePath("/")

	Action("health", func() {
		Routing(
			GET("/health"),
		)
		Description("Perform health check.")
		Response(OK, HealthMedia)
	})
})

var HealthMedia = MediaType("application/json", func() {
	Description("Health Check")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("status", String, "system status")
		Attribute("code", Integer, "status code")
	})
	Required("status", "code")
	View("default", func() {
		Attribute("status", String)
		Attribute("code", Integer)
	})
})
