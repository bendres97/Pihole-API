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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// DefaultAPIController binds http requests to an api service and writes the service results to the http response
type DefaultAPIController struct {
	service DefaultAPIServicer
	errorHandler ErrorHandler
}

// DefaultAPIOption for how the controller is set up.
type DefaultAPIOption func(*DefaultAPIController)

// WithDefaultAPIErrorHandler inject ErrorHandler into controller
func WithDefaultAPIErrorHandler(h ErrorHandler) DefaultAPIOption {
	return func(c *DefaultAPIController) {
		c.errorHandler = h
	}
}

// NewDefaultAPIController creates a default api controller
func NewDefaultAPIController(s DefaultAPIServicer, opts ...DefaultAPIOption) Router {
	controller := &DefaultAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DefaultAPIController
func (c *DefaultAPIController) Routes() Routes {
	return Routes{
		"GravityGet": Route{
			strings.ToUpper("Get"),
			"/gravity",
			c.GravityGet,
		},
		"GravityIdDelete": Route{
			strings.ToUpper("Delete"),
			"/gravity/{id}",
			c.GravityIdDelete,
		},
		"GravityIdPatch": Route{
			strings.ToUpper("Patch"),
			"/gravity/{id}",
			c.GravityIdPatch,
		},
		"GravityPost": Route{
			strings.ToUpper("Post"),
			"/gravity",
			c.GravityPost,
		},
		"StatusGet": Route{
			strings.ToUpper("Get"),
			"/status",
			c.StatusGet,
		},
	}
}

// GravityGet - 
func (c *DefaultAPIController) GravityGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GravityGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GravityIdDelete - 
func (c *DefaultAPIController) GravityIdDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam, err := parseNumericParameter[int32](
		params["id"],
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GravityIdDelete(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GravityIdPatch - 
func (c *DefaultAPIController) GravityIdPatch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam, err := parseNumericParameter[int32](
		params["id"],
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	gravityObjParam := GravityObj{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&gravityObjParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertGravityObjRequired(gravityObjParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertGravityObjConstraints(gravityObjParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.GravityIdPatch(r.Context(), idParam, gravityObjParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GravityPost - 
func (c *DefaultAPIController) GravityPost(w http.ResponseWriter, r *http.Request) {
	gravityObjParam := GravityObj{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&gravityObjParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertGravityObjRequired(gravityObjParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertGravityObjConstraints(gravityObjParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.GravityPost(r.Context(), gravityObjParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// StatusGet - 
func (c *DefaultAPIController) StatusGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.StatusGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
