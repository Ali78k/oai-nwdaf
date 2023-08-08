/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

type UmtTime struct {

	// String with format partial-time or full-time as defined in clause 5.6 of IETF RFC 3339. Examples, 20:15:00, 20:15:00-08:00 (for 8 hours behind UTC).
	TimeOfDay string `json:"timeOfDay"`

	// integer between and including 1 and 7 denoting a weekday. 1 shall indicate Monday, and the subsequent weekdays  shall be indicated with the next higher numbers. 7 shall indicate Sunday.
	DayOfWeek int32 `json:"dayOfWeek"`
}

// AssertUmtTimeRequired checks if the required fields are not zero-ed
func AssertUmtTimeRequired(obj UmtTime) error {
	elements := map[string]interface{}{
		"timeOfDay": obj.TimeOfDay,
		"dayOfWeek": obj.DayOfWeek,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseUmtTimeRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of UmtTime (e.g. [][]UmtTime), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUmtTimeRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUmtTime, ok := obj.(UmtTime)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUmtTimeRequired(aUmtTime)
	})
}
