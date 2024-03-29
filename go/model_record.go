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



type Record struct {

	// CNAME or A
	Type string `json:"type"`

	// The domain to resolve
	Domain string `json:"domain"`

	// The destination IP or DNS record
	Destination string `json:"destination"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *Record) UnmarshalJSON(data []byte) error {

	type Alias Record // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertRecordRequired checks if the required fields are not zero-ed
func AssertRecordRequired(obj Record) error {
	elements := map[string]interface{}{
		"type": obj.Type,
		"domain": obj.Domain,
		"destination": obj.Destination,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecordConstraints checks if the values respects the defined constraints
func AssertRecordConstraints(obj Record) error {
	return nil
}
