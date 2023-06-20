/*
 * PiHole Management API
 *
 * A management webservice to be run locally on a PiHole instance
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
)



// DefaultAPIRouter defines the required methods for binding the api requests to a responses for the DefaultAPI
// The DefaultAPIRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultAPIServicer to perform the required actions, then write the service results to the http response.
type DefaultAPIRouter interface { 
	GravityGet(http.ResponseWriter, *http.Request)
	GravityIdDelete(http.ResponseWriter, *http.Request)
	GravityPatch(http.ResponseWriter, *http.Request)
	GravityPost(http.ResponseWriter, *http.Request)
	StatusActionPost(http.ResponseWriter, *http.Request)
	StatusGet(http.ResponseWriter, *http.Request)
}


// DefaultAPIServicer defines the api actions for the DefaultAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DefaultAPIServicer interface { 
	GravityGet(context.Context) (ImplResponse, error)
	GravityIdDelete(context.Context, int32) (ImplResponse, error)
	GravityPatch(context.Context) (ImplResponse, error)
	GravityPost(context.Context, GravityObj) (ImplResponse, error)
	StatusActionPost(context.Context, string) (ImplResponse, error)
	StatusGet(context.Context) (ImplResponse, error)
}
