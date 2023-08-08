/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

type ExceptionIdAnyOf string

// List of ExceptionIdAnyOf
const (
	EXCEPTIONIDANYOF_UNEXPECTED_UE_LOCATION         ExceptionIdAnyOf = "UNEXPECTED_UE_LOCATION"
	EXCEPTIONIDANYOF_UNEXPECTED_LONG_LIVE_FLOW      ExceptionIdAnyOf = "UNEXPECTED_LONG_LIVE_FLOW"
	EXCEPTIONIDANYOF_UNEXPECTED_LARGE_RATE_FLOW     ExceptionIdAnyOf = "UNEXPECTED_LARGE_RATE_FLOW"
	EXCEPTIONIDANYOF_UNEXPECTED_WAKEUP              ExceptionIdAnyOf = "UNEXPECTED_WAKEUP"
	EXCEPTIONIDANYOF_SUSPICION_OF_DDOS_ATTACK       ExceptionIdAnyOf = "SUSPICION_OF_DDOS_ATTACK"
	EXCEPTIONIDANYOF_WRONG_DESTINATION_ADDRESS      ExceptionIdAnyOf = "WRONG_DESTINATION_ADDRESS"
	EXCEPTIONIDANYOF_TOO_FREQUENT_SERVICE_ACCESS    ExceptionIdAnyOf = "TOO_FREQUENT_SERVICE_ACCESS"
	EXCEPTIONIDANYOF_UNEXPECTED_RADIO_LINK_FAILURES ExceptionIdAnyOf = "UNEXPECTED_RADIO_LINK_FAILURES"
	EXCEPTIONIDANYOF_PING_PONG_ACROSS_CELLS         ExceptionIdAnyOf = "PING_PONG_ACROSS_CELLS"
)

// AssertExceptionIdAnyOfRequired checks if the required fields are not zero-ed
func AssertExceptionIdAnyOfRequired(obj ExceptionIdAnyOf) error {
	return nil
}

// AssertRecurseExceptionIdAnyOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ExceptionIdAnyOf (e.g. [][]ExceptionIdAnyOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseExceptionIdAnyOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aExceptionIdAnyOf, ok := obj.(ExceptionIdAnyOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertExceptionIdAnyOfRequired(aExceptionIdAnyOf)
	})
}
