/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

type AdrfDataTypeAnyOf string

// List of AdrfDataTypeAnyOf
const (
	ADRFDATATYPEANYOF_ANALYTICS AdrfDataTypeAnyOf = "HISTORICAL_ANALYTICS"
	ADRFDATATYPEANYOF_DATA      AdrfDataTypeAnyOf = "HISTORICAL_DATA"
)

// AssertAdrfDataTypeAnyOfRequired checks if the required fields are not zero-ed
func AssertAdrfDataTypeAnyOfRequired(obj AdrfDataTypeAnyOf) error {
	return nil
}

// AssertRecurseAdrfDataTypeAnyOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AdrfDataTypeAnyOf (e.g. [][]AdrfDataTypeAnyOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAdrfDataTypeAnyOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAdrfDataTypeAnyOf, ok := obj.(AdrfDataTypeAnyOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAdrfDataTypeAnyOfRequired(aAdrfDataTypeAnyOf)
	})
}