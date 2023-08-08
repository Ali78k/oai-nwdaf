/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// LocationInfo1 - Represents UE location information.
type LocationInfo1 struct {
	Loc UserLocation `json:"loc"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	Ratio int32 `json:"ratio,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty"`
}

// AssertLocationInfo1Required checks if the required fields are not zero-ed
func AssertLocationInfo1Required(obj LocationInfo1) error {
	elements := map[string]interface{}{
		"loc": obj.Loc,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertUserLocationRequired(obj.Loc); err != nil {
		return err
	}
	return nil
}

// AssertRecurseLocationInfo1Required recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of LocationInfo1 (e.g. [][]LocationInfo1), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseLocationInfo1Required(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aLocationInfo1, ok := obj.(LocationInfo1)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertLocationInfo1Required(aLocationInfo1)
	})
}
