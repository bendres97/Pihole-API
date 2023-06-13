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
	"bytes"
	"context"
	"errors"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

// DefaultAPIService is a service that implements the logic for the DefaultAPIServicer
// This service should implement the business logic for every endpoint for the DefaultAPI API.
// Include any external packages or services that will be required by this service.
type DefaultAPIService struct {
}

// NewDefaultAPIService creates a default api service
func NewDefaultAPIService() DefaultAPIServicer {
	return &DefaultAPIService{}
}

// GravityGet -
func (s *DefaultAPIService) GravityGet(ctx context.Context) (ImplResponse, error) {

	cmd := exec.Command("/usr/bin/sqlite3", "/etc/pihole/gravity.db", "SELECT id, address, comment FROM adlist;")
	output, err := cmd.Output()
	if err != nil {
		var errb bytes.Buffer
		cmd.Stderr = &errb
		log.Print(cmd.Args)
		log.Print(string(output))
		log.Print(errb.String())
		return Response(500, err.Error()), err
	}

	var gravity []GravityObj

	result := strings.TrimSpace(string(output))
	lines := strings.Split(result, "\n")
	for _, line := range lines {
		split := strings.Split(line, "|")
		id_val, err := strconv.Atoi(split[0])
		if err != nil {
			return Response(500, err.Error()), err
		}
		id := int32(id_val)
		address := split[1]
		comment := split[2]
		gravity = append(gravity, GravityObj{
			Id:      id,
			Address: address,
			Comment: comment,
		})
	}

	return Response(200, gravity), nil
}

// GravityIdDelete -
func (s *DefaultAPIService) GravityIdDelete(ctx context.Context, id int32) (ImplResponse, error) {
	// TODO - update GravityIdDelete with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, string{}) or use other options such as http.Ok ...
	// return Response(200, string{}), nil

	// TODO: Uncomment the next line to return response Response(404, GravityIdDelete404Response{}) or use other options such as http.Ok ...
	// return Response(404, GravityIdDelete404Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GravityIdDelete method not implemented")
}

// GravityIdPatch -
func (s *DefaultAPIService) GravityIdPatch(ctx context.Context, id int32, gravityObj GravityObj) (ImplResponse, error) {
	// TODO - update GravityIdPatch with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, GravityObj{}) or use other options such as http.Ok ...
	// return Response(200, GravityObj{}), nil

	// TODO: Uncomment the next line to return response Response(404, GravityIdDelete404Response{}) or use other options such as http.Ok ...
	// return Response(404, GravityIdDelete404Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GravityIdPatch method not implemented")
}

// GravityPost -
func (s *DefaultAPIService) GravityPost(ctx context.Context, gravityObj GravityObj) (ImplResponse, error) {
	// TODO - update GravityPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, GravityObj{}) or use other options such as http.Ok ...
	// return Response(200, GravityObj{}), nil

	// TODO: Uncomment the next line to return response Response(201, GravityObj{}) or use other options such as http.Ok ...
	// return Response(201, GravityObj{}), nil

	// TODO: Uncomment the next line to return response Response(202, GravityObj{}) or use other options such as http.Ok ...
	// return Response(202, GravityObj{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GravityPost method not implemented")
}