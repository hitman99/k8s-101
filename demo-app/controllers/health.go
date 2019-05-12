package controllers

import (
	"github.com/goadesign/goa"
	"github.com/hitman99/k8s-101/demo-app/app"
)

// HealthController implements the health resource.
type HealthController struct {
	*goa.Controller
}

// NewHealthController creates a health controller.
func NewHealthController(service *goa.Service) *HealthController {
	return &HealthController{Controller: service.NewController("HealthController")}
}

// Health runs the health action.
func (c *HealthController) Health(ctx *app.HealthHealthContext) error {
	// HealthController_Health: start_implement

	// Put your logic here

	res := &app.JSON{
		Code:   200,
		Status: "OK",
	}
	return ctx.OK(res)
	// HealthController_Health: end_implement
}
