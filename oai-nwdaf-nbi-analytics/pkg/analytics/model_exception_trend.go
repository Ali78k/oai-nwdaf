/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

// ExceptionTrend - Possible values are   - UP: Up trend of the exception level.   - DOWN: Down trend of the exception level.   - UNKNOW: Unknown trend of the exception level.   - STABLE: Stable trend of the exception level.
type ExceptionTrend struct {
}

// AssertExceptionTrendRequired checks if the required fields are not zero-ed
func AssertExceptionTrendRequired(obj ExceptionTrend) error {
	return nil
}

// AssertRecurseExceptionTrendRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ExceptionTrend (e.g. [][]ExceptionTrend), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseExceptionTrendRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aExceptionTrend, ok := obj.(ExceptionTrend)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertExceptionTrendRequired(aExceptionTrend)
	})
}
