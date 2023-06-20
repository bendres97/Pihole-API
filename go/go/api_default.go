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
		"GravityPatch": Route{
			strings.ToUpper("Patch"),
			"/gravity",
			c.GravityPatch,
		},
		"GravityPost": Route{
			strings.ToUpper("Post"),
			"/gravity",
			c.GravityPost,
		},
		"RecordsDelete": Route{
			strings.ToUpper("Delete"),
			"/records",
			c.RecordsDelete,
		},
		"RecordsGet": Route{
			strings.ToUpper("Get"),
			"/records",
			c.RecordsGet,
		},
		"RecordsPost": Route{
			strings.ToUpper("Post"),
			"/records",
			c.RecordsPost,
		},
		"StatusActionPost": Route{
			strings.ToUpper("Post"),
			"/status/{action}",
			c.StatusActionPost,
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

// GravityPatch - 
func (c *DefaultAPIController) GravityPatch(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GravityPatch(r.Context())
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

// RecordsDelete - 
func (c *DefaultAPIController) RecordsDelete(w http.ResponseWriter, r *http.Request) {
	deleteRecordParam := DeleteRecord{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&deleteRecordParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertDeleteRecordRequired(deleteRecordParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertDeleteRecordConstraints(deleteRecordParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.RecordsDelete(r.Context(), deleteRecordParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// RecordsGet - 
func (c *DefaultAPIController) RecordsGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.RecordsGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// RecordsPost - 
func (c *DefaultAPIController) RecordsPost(w http.ResponseWriter, r *http.Request) {
	recordParam := Record{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&recordParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertRecordRequired(recordParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertRecordConstraints(recordParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.RecordsPost(r.Context(), recordParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// StatusActionPost - 
func (c *DefaultAPIController) StatusActionPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	actionParam := params["action"]
	result, err := c.service.StatusActionPost(r.Context(), actionParam)
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
