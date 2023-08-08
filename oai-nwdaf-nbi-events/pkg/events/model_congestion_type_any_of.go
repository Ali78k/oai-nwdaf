/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

type CongestionTypeAnyOf string

// List of CongestionTypeAnyOf
const (
	CONGESTIONTYPEANYOF_USER_PLANE             CongestionTypeAnyOf = "USER_PLANE"
	CONGESTIONTYPEANYOF_CONTROL_PLANE          CongestionTypeAnyOf = "CONTROL_PLANE"
	CONGESTIONTYPEANYOF_USER_AND_CONTROL_PLANE CongestionTypeAnyOf = "USER_AND_CONTROL_PLANE"
)

// AssertCongestionTypeAnyOfRequired checks if the required fields are not zero-ed
func AssertCongestionTypeAnyOfRequired(obj CongestionTypeAnyOf) error {
	return nil
}

// AssertRecurseCongestionTypeAnyOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CongestionTypeAnyOf (e.g. [][]CongestionTypeAnyOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCongestionTypeAnyOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCongestionTypeAnyOf, ok := obj.(CongestionTypeAnyOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCongestionTypeAnyOfRequired(aCongestionTypeAnyOf)
	})
}
