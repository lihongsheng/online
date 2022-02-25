// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"online/server/web/gen/models"
)

// GetUserTagsHandlerFunc turns a function with the right signature into a get user tags handler
type GetUserTagsHandlerFunc func(GetUserTagsParams, *models.Principle) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserTagsHandlerFunc) Handle(params GetUserTagsParams, principal *models.Principle) middleware.Responder {
	return fn(params, principal)
}

// GetUserTagsHandler interface for that can handle valid get user tags params
type GetUserTagsHandler interface {
	Handle(GetUserTagsParams, *models.Principle) middleware.Responder
}

// NewGetUserTags creates a new http.Handler for the get user tags operation
func NewGetUserTags(ctx *middleware.Context, handler GetUserTagsHandler) *GetUserTags {
	return &GetUserTags{Context: ctx, Handler: handler}
}

/* GetUserTags swagger:route GET /user/tags getUserTags

GetUserTags get user tags API

*/
type GetUserTags struct {
	Context *middleware.Context
	Handler GetUserTagsHandler
}

func (o *GetUserTags) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUserTagsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
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