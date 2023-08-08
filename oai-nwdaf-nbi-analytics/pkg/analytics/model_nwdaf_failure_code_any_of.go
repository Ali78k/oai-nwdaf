/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

type NwdafFailureCodeAnyOf string

// List of NwdafFailureCodeAnyOf
const (
	NWDAFFAILURECODEANYOF_UNAVAILABLE_DATA                     NwdafFailureCodeAnyOf = "UNAVAILABLE_DATA"
	NWDAFFAILURECODEANYOF_BOTH_STAT_PRED_NOT_ALLOWED           NwdafFailureCodeAnyOf = "BOTH_STAT_PRED_NOT_ALLOWED"
	NWDAFFAILURECODEANYOF_UNSATISFIED_REQUESTED_ANALYTICS_TIME NwdafFailureCodeAnyOf = "UNSATISFIED_REQUESTED_ANALYTICS_TIME"
	NWDAFFAILURECODEANYOF_OTHER                                NwdafFailureCodeAnyOf = "OTHER"
)

// AssertNwdafFailureCodeAnyOfRequired checks if the required fields are not zero-ed
func AssertNwdafFailureCodeAnyOfRequired(obj NwdafFailureCodeAnyOf) error {
	return nil
}

// AssertRecurseNwdafFailureCodeAnyOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of NwdafFailureCodeAnyOf (e.g. [][]NwdafFailureCodeAnyOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNwdafFailureCodeAnyOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNwdafFailureCodeAnyOf, ok := obj.(NwdafFailureCodeAnyOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNwdafFailureCodeAnyOfRequired(aNwdafFailureCodeAnyOf)
	})
}
