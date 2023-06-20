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
)



type Status struct {

	Listening bool `json:"listening"`

	Ipv4 UdpTcp `json:"ipv4"`

	Ipv6 UdpTcp `json:"ipv6"`

	Blocking bool `json:"blocking"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *Status) UnmarshalJSON(data []byte) error {

	type Alias Status // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertStatusRequired checks if the required fields are not zero-ed
func AssertStatusRequired(obj Status) error {
	elements := map[string]interface{}{
		"listening": obj.Listening,
		"ipv4": obj.Ipv4,
		"ipv6": obj.Ipv6,
		"blocking": obj.Blocking,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertUdpTcpRequired(obj.Ipv4); err != nil {
		return err
	}
	if err := AssertUdpTcpRequired(obj.Ipv6); err != nil {
		return err
	}
	return nil
}

// AssertStatusConstraints checks if the values respects the defined constraints
func AssertStatusConstraints(obj Status) error {
	return nil
}