/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

// MatchingDirection - Possible values are   - ASCENDING: Threshold is crossed in ascending direction.   - DESCENDING: Threshold is crossed in descending direction.   - CROSSED: Threshold is crossed either in ascending or descending direction.
type MatchingDirection struct {
}

// AssertMatchingDirectionRequired checks if the required fields are not zero-ed
func AssertMatchingDirectionRequired(obj MatchingDirection) error {
	return nil
}

// AssertRecurseMatchingDirectionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of MatchingDirection (e.g. [][]MatchingDirection), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseMatchingDirectionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aMatchingDirection, ok := obj.(MatchingDirection)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertMatchingDirectionRequired(aMatchingDirection)
	})
}
