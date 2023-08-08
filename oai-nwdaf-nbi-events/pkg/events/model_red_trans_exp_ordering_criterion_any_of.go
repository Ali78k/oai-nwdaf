/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

type RedTransExpOrderingCriterionAnyOf string

// List of RedTransExpOrderingCriterionAnyOf
const (
	REDTRANSEXPORDERINGCRITERIONANYOF_TIME_SLOT_START RedTransExpOrderingCriterionAnyOf = "TIME_SLOT_START"
	REDTRANSEXPORDERINGCRITERIONANYOF_RED_TRANS_EXP   RedTransExpOrderingCriterionAnyOf = "RED_TRANS_EXP"
)

// AssertRedTransExpOrderingCriterionAnyOfRequired checks if the required fields are not zero-ed
func AssertRedTransExpOrderingCriterionAnyOfRequired(obj RedTransExpOrderingCriterionAnyOf) error {
	return nil
}

// AssertRecurseRedTransExpOrderingCriterionAnyOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of RedTransExpOrderingCriterionAnyOf (e.g. [][]RedTransExpOrderingCriterionAnyOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseRedTransExpOrderingCriterionAnyOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aRedTransExpOrderingCriterionAnyOf, ok := obj.(RedTransExpOrderingCriterionAnyOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertRedTransExpOrderingCriterionAnyOfRequired(aRedTransExpOrderingCriterionAnyOf)
	})
}
