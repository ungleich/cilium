// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package recorder

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetRecorderHandlerFunc turns a function with the right signature into a get recorder handler
type GetRecorderHandlerFunc func(GetRecorderParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRecorderHandlerFunc) Handle(params GetRecorderParams) middleware.Responder {
	return fn(params)
}

// GetRecorderHandler interface for that can handle valid get recorder params
type GetRecorderHandler interface {
	Handle(GetRecorderParams) middleware.Responder
}

// NewGetRecorder creates a new http.Handler for the get recorder operation
func NewGetRecorder(ctx *middleware.Context, handler GetRecorderHandler) *GetRecorder {
	return &GetRecorder{Context: ctx, Handler: handler}
}

/*
GetRecorder swagger:route GET /recorder recorder getRecorder

Retrieve list of all recorders
*/
type GetRecorder struct {
	Context *middleware.Context
	Handler GetRecorderHandler
}

func (o *GetRecorder) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRecorderParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
