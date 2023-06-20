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
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

var SQLLITE_EXECUTABLE = "/usr/bin/sqlite3"
var SQLLITE_DATABASE = "/etc/pihole/gravity.db"
var PIHOLE_EXECUTABLE = "/usr/local/bin/pihole"

// DefaultAPIService is a service that implements the logic for the DefaultAPIServicer
// This service should implement the business logic for every endpoint for the DefaultAPI API.
// Include any external packages or services that will be required by this service.
type DefaultAPIService struct {
}

// NewDefaultAPIService creates a default api service
func NewDefaultAPIService() DefaultAPIServicer {
	return &DefaultAPIService{}
}

// StatusActionPost
func (s *DefaultAPIService) StatusActionPost(ctx context.Context, action string) (ImplResponse, error) {
	action = strings.ToLower(action)
	var command string
	switch action {
	case "restart":
		command = "restartdns"
	case "enable":
		command = "enable"
	case "disable":
		command = "disable"
	default:
		return Response(400, "Valid options are `restartdns`, `enable`, and `disable`"), errors.New("Invalid Request")
	}
	cmd := exec.Command(PIHOLE_EXECUTABLE, command)
	output, err := cmd.Output()
	if err != nil {
		var errb bytes.Buffer
		cmd.Stderr = &errb
		log.Print(cmd.Args)
		log.Print(string(output))
		log.Print(errb.String())
		return Response(500, err.Error()), err
	}

	return Response(200, string(output)), nil
}

// StatusGet -
func (s *DefaultAPIService) StatusGet(ctx context.Context) (ImplResponse, error) {

	cmd := exec.Command(PIHOLE_EXECUTABLE, "status")
	output, err := cmd.Output()
	if err != nil {
		var errb bytes.Buffer
		cmd.Stderr = &errb
		log.Print(cmd.Args)
		log.Print(string(output))
		log.Print(errb.String())
		return Response(500, err.Error()), err
	}

	result := strings.TrimSpace(string(output))

	status := Status{
		Listening: strings.Contains(result, "[✓] FTL is listening on port 53"),
		Blocking:  strings.Contains(result, "[✓] Pi-hole blocking is enabled"),
		Ipv4: UdpTcp{
			Tcp: strings.Contains(result, "[✓] TCP (IPv4)"),
			Udp: strings.Contains(result, "[✓] UDP (IPv4)"),
		},

		Ipv6: UdpTcp{
			Tcp: strings.Contains(result, "[✓] TCP (IPv6)"),
			Udp: strings.Contains(result, "[✓] UDP (IPv6)"),
		},
	}

	if status.Listening {
		if status.Blocking {
			return Response(200, status), nil
		} else {
			return Response(201, status), nil
		}
	} else {
		return Response(500, status), nil
	}
}

// GravityGet -
func (s *DefaultAPIService) GravityGet(ctx context.Context) (ImplResponse, error) {

	cmd := exec.Command(SQLLITE_EXECUTABLE, SQLLITE_DATABASE, "SELECT id, address, comment FROM adlist;")
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

	cmd := exec.Command(SQLLITE_EXECUTABLE, SQLLITE_DATABASE, fmt.Sprintf("DELETE FROM adlist WHERE id = %v", id))
	output, err := cmd.Output()
	if err != nil {
		var errb bytes.Buffer
		cmd.Stderr = &errb
		log.Print(cmd.Args)
		log.Print(string(output))
		log.Print(errb.String())
		return Response(500, err.Error()), err
	}

	return Response(200, "Domain deleted"), nil
}

// GravityPatch -
func (s *DefaultAPIService) GravityPatch(ctx context.Context) (ImplResponse, error) {
	cmd := exec.Command("/usr/local/bin/pihole", "-g")
	output, err := cmd.Output()
	if err != nil {
		var errb bytes.Buffer
		cmd.Stderr = &errb
		log.Print(cmd.Args)
		log.Print(string(output))
		log.Print(errb.String())
		return Response(500, err.Error()), err
	}

	return Response(200, "Gravity Updated"), nil
}

// GravityPost -
func (s *DefaultAPIService) GravityPost(ctx context.Context, gravityObj GravityObj) (ImplResponse, error) {

	cmd := exec.Command(SQLLITE_EXECUTABLE,
		SQLLITE_DATABASE,
		fmt.Sprintf("SELECT id, address, comment FROM adlist WHERE address = '%v';", gravityObj.Address),
	)
	output, err := cmd.Output()
	if err != nil {
		var errb bytes.Buffer
		cmd.Stderr = &errb
		log.Print(cmd.Args)
		log.Print(string(output))
		log.Print(errb.String())
		return Response(500, err.Error()), err
	}

	result := strings.TrimSpace(string(output))
	if len(result) > 0 {
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
			gravity := GravityObj{
				Id:      id,
				Address: address,
				Comment: comment,
			}
			if len(lines) > 1 {
				return Response(500, nil), fmt.Errorf("more than one result found for '%v'", address)
			} else if len(lines) == 1 {
				if address == gravityObj.Address && comment == gravityObj.Comment {
					return Response(200, gravity), nil
				} else {
					updatecmd := exec.Command(SQLLITE_EXECUTABLE,
						SQLLITE_DATABASE,
						fmt.Sprintf("UPDATE adlist SET address = '%v', comment = '%v' WHERE id = %v", gravity.Address, gravity.Comment, id),
					)
					updateOutput, err := updatecmd.Output()
					if err != nil {
						var errb bytes.Buffer
						updatecmd.Stderr = &errb
						log.Print(updatecmd.Args)
						log.Print(string(updateOutput))
						log.Print(errb.String())
						return Response(500, err.Error()), err
					}
					gravity.Address = gravityObj.Address
					gravity.Comment = gravityObj.Comment
					gravity.Id = id
					return Response(201, gravity), nil
				}
			}
		}

	} else {
		insertcmd := exec.Command(SQLLITE_EXECUTABLE,
			SQLLITE_DATABASE,
			fmt.Sprintf("INSERT INTO adlist (address, comment) VALUES ('%v','%v')", gravityObj.Address, gravityObj.Comment),
		)
		insertOutput, err := insertcmd.Output()
		if err != nil {
			var errb bytes.Buffer
			insertcmd.Stderr = &errb
			log.Print(insertcmd.Args)
			log.Print(string(insertOutput))
			log.Print(errb.String())
			return Response(500, err.Error()), err
		}

		checkcmd := exec.Command(SQLLITE_EXECUTABLE,
			SQLLITE_DATABASE,
			fmt.Sprintf("SELECT id, address, comment FROM adlist WHERE address = '%v';", gravityObj.Address),
		)
		checkOutput, err := checkcmd.Output()
		if err != nil {
			var errb bytes.Buffer
			checkcmd.Stderr = &errb
			log.Print(checkcmd.Args)
			log.Print(string(checkOutput))
			log.Print(errb.String())
			return Response(500, err.Error()), err
		}
		return Response(200, gravityObj), nil
	}
	return Response(500, "Unknown Error"), errors.New("reached end of method unexpectedly")
}
