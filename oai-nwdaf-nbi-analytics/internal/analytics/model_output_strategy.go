/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

// OutputStrategy - Possible values are   - BINARY: Indicates that the analytics shall only be reported when the requested level of accuracy is reached within a cycle of periodic notification.   - GRADIENT: Indicates that the analytics shall be reported according with the periodicity irrespective of whether the requested level of accuracy has been reached or not.
type OutputStrategy struct {
}

// AssertOutputStrategyRequired checks if the required fields are not zero-ed
func AssertOutputStrategyRequired(obj OutputStrategy) error {
	return nil
}

// AssertRecurseOutputStrategyRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of OutputStrategy (e.g. [][]OutputStrategy), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseOutputStrategyRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aOutputStrategy, ok := obj.(OutputStrategy)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertOutputStrategyRequired(aOutputStrategy)
	})
}
