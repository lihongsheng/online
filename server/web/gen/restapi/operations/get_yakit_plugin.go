// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"online/server/web/gen/models"
)

// GetYakitPluginHandlerFunc turns a function with the right signature into a get yakit plugin handler
type GetYakitPluginHandlerFunc func(GetYakitPluginParams, *models.Principle) middleware.Responder

// Handle executing the request and returning a response
func (fn GetYakitPluginHandlerFunc) Handle(params GetYakitPluginParams, principal *models.Principle) middleware.Responder {
	return fn(params, principal)
}

// GetYakitPluginHandler interface for that can handle valid get yakit plugin params
type GetYakitPluginHandler interface {
	Handle(GetYakitPluginParams, *models.Principle) middleware.Responder
}

// NewGetYakitPlugin creates a new http.Handler for the get yakit plugin operation
func NewGetYakitPlugin(ctx *middleware.Context, handler GetYakitPluginHandler) *GetYakitPlugin {
	return &GetYakitPlugin{Context: ctx, Handler: handler}
}

/* GetYakitPlugin swagger:route GET /yakit/plugin getYakitPlugin

GetYakitPlugin get yakit plugin API

*/
type GetYakitPlugin struct {
	Context *middleware.Context
	Handler GetYakitPluginHandler
}

func (o *GetYakitPlugin) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetYakitPluginParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principle
	if uprinc != nil {
		principal = uprinc.(*models.Principle) // this is really a models.Principle, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
