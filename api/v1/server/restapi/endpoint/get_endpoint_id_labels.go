// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetEndpointIDLabelsHandlerFunc turns a function with the right signature into a get endpoint ID labels handler
type GetEndpointIDLabelsHandlerFunc func(GetEndpointIDLabelsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetEndpointIDLabelsHandlerFunc) Handle(params GetEndpointIDLabelsParams) middleware.Responder {
	return fn(params)
}

// GetEndpointIDLabelsHandler interface for that can handle valid get endpoint ID labels params
type GetEndpointIDLabelsHandler interface {
	Handle(GetEndpointIDLabelsParams) middleware.Responder
}

// NewGetEndpointIDLabels creates a new http.Handler for the get endpoint ID labels operation
func NewGetEndpointIDLabels(ctx *middleware.Context, handler GetEndpointIDLabelsHandler) *GetEndpointIDLabels {
	return &GetEndpointIDLabels{Context: ctx, Handler: handler}
}

/*
GetEndpointIDLabels swagger:route GET /endpoint/{id}/labels endpoint getEndpointIdLabels

Retrieves the list of labels associated with an endpoint.
*/
type GetEndpointIDLabels struct {
	Context *middleware.Context
	Handler GetEndpointIDLabelsHandler
}

func (o *GetEndpointIDLabels) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetEndpointIDLabelsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
