/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

// NumberAverage - Represents average and variance information.
type NumberAverage struct {
	Number int32 `json:"number"`

	// string with format 'float' as defined in OpenAPI.
	Variance float32 `json:"variance"`
}

// AssertNumberAverageRequired checks if the required fields are not zero-ed
func AssertNumberAverageRequired(obj NumberAverage) error {
	elements := map[string]interface{}{
		"number":   obj.Number,
		"variance": obj.Variance,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseNumberAverageRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of NumberAverage (e.g. [][]NumberAverage), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNumberAverageRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNumberAverage, ok := obj.(NumberAverage)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNumberAverageRequired(aNumberAverage)
	})
}