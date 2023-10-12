/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// PlmnId - When PlmnId needs to be converted to string (e.g. when used in maps as key), the string  shall be composed of three digits \"mcc\" followed by \"-\" and two or three digits \"mnc\".
type PlmnId struct {

	// Mobile Country Code part of the PLMN, comprising 3 digits, as defined in clause 9.3.3.5 of 3GPP TS 38.413.
	Mcc string `json:"mcc"`

	// Mobile Network Code part of the PLMN, comprising 2 or 3 digits, as defined in clause 9.3.3.5 of 3GPP TS 38.413.
	Mnc string `json:"mnc"`
}

// AssertPlmnIdRequired checks if the required fields are not zero-ed
func AssertPlmnIdRequired(obj PlmnId) error {
	elements := map[string]interface{}{
		"mcc": obj.Mcc,
		"mnc": obj.Mnc,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecursePlmnIdRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PlmnId (e.g. [][]PlmnId), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePlmnIdRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPlmnId, ok := obj.(PlmnId)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPlmnIdRequired(aPlmnId)
	})
}