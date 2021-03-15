// Code generated by go-swagger; DO NOT EDIT.

// Copyright Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package alertgroup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAlertGroupsHandlerFunc turns a function with the right signature into a get alert groups handler
type GetAlertGroupsHandlerFunc func(GetAlertGroupsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAlertGroupsHandlerFunc) Handle(params GetAlertGroupsParams) middleware.Responder {
	return fn(params)
}

// GetAlertGroupsHandler interface for that can handle valid get alert groups params
type GetAlertGroupsHandler interface {
	Handle(GetAlertGroupsParams) middleware.Responder
}

// NewGetAlertGroups creates a new http.Handler for the get alert groups operation
func NewGetAlertGroups(ctx *middleware.Context, handler GetAlertGroupsHandler) *GetAlertGroups {
	return &GetAlertGroups{Context: ctx, Handler: handler}
}

/*GetAlertGroups swagger:route GET /alerts/groups alertgroup getAlertGroups

Get a list of alert groups

*/
type GetAlertGroups struct {
	Context *middleware.Context
	Handler GetAlertGroupsHandler
}

func (o *GetAlertGroups) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAlertGroupsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
