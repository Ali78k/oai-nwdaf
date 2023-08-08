/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

// LocationInfo - Represents UE location information.
type LocationInfo struct {
	Loc UserLocation `json:"loc"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	Ratio int32 `json:"ratio,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty"`
}

// AssertLocationInfoRequired checks if the required fields are not zero-ed
func AssertLocationInfoRequired(obj LocationInfo) error {
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

// AssertRecurseLocationInfoRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of LocationInfo (e.g. [][]LocationInfo), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseLocationInfoRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aLocationInfo, ok := obj.(LocationInfo)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertLocationInfoRequired(aLocationInfo)
	})
}
