/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// NsiIdInfo1 - Represents the S-NSSAI and the optionally associated Network Slice Instance(s).
type NsiIdInfo1 struct {
	Snssai Snssai `json:"snssai"`

	NsiIds []string `json:"nsiIds,omitempty"`
}

// AssertNsiIdInfo1Required checks if the required fields are not zero-ed
func AssertNsiIdInfo1Required(obj NsiIdInfo1) error {
	elements := map[string]interface{}{
		"snssai": obj.Snssai,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertSnssaiRequired(obj.Snssai); err != nil {
		return err
	}
	return nil
}

// AssertRecurseNsiIdInfo1Required recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of NsiIdInfo1 (e.g. [][]NsiIdInfo1), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNsiIdInfo1Required(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNsiIdInfo1, ok := obj.(NsiIdInfo1)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNsiIdInfo1Required(aNsiIdInfo1)
	})
}
