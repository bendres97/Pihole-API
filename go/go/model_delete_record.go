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



type DeleteRecord struct {

	// CNAME or A
	Type string `json:"type"`

	// The domain to resolve
	Domain string `json:"domain"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *DeleteRecord) UnmarshalJSON(data []byte) error {

	type Alias DeleteRecord // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertDeleteRecordRequired checks if the required fields are not zero-ed
func AssertDeleteRecordRequired(obj DeleteRecord) error {
	elements := map[string]interface{}{
		"type": obj.Type,
		"domain": obj.Domain,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertDeleteRecordConstraints checks if the values respects the defined constraints
func AssertDeleteRecordConstraints(obj DeleteRecord) error {
	return nil
}
