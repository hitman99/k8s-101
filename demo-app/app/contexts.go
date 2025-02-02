// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "demo": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/hitman99/k8s-101/demo-app/design
// --out=$(GOPATH)/src/github.com/hitman99/k8s-101/demo-app
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// HealthHealthContext provides the health health action context.
type HealthHealthContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewHealthHealthContext parses the incoming request URL and body, performs validations and creates the
// context used by the health controller health action.
func NewHealthHealthContext(ctx context.Context, r *http.Request, service *goa.Service) (*HealthHealthContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := HealthHealthContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *HealthHealthContext) OK(r *JSON) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}
