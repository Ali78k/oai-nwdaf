/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

type ExceptionTrendAnyOf string

// List of ExceptionTrendAnyOf
const (
	EXCEPTIONTRENDANYOF_UP     ExceptionTrendAnyOf = "UP"
	EXCEPTIONTRENDANYOF_DOWN   ExceptionTrendAnyOf = "DOWN"
	EXCEPTIONTRENDANYOF_UNKNOW ExceptionTrendAnyOf = "UNKNOW"
	EXCEPTIONTRENDANYOF_STABLE ExceptionTrendAnyOf = "STABLE"
)

// AssertExceptionTrendAnyOfRequired checks if the required fields are not zero-ed
func AssertExceptionTrendAnyOfRequired(obj ExceptionTrendAnyOf) error {
	return nil
}

// AssertRecurseExceptionTrendAnyOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ExceptionTrendAnyOf (e.g. [][]ExceptionTrendAnyOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseExceptionTrendAnyOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aExceptionTrendAnyOf, ok := obj.(ExceptionTrendAnyOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertExceptionTrendAnyOfRequired(aExceptionTrendAnyOf)
	})
}
