/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// TimeUnit - Possible values are - MINUTE: Time unit is per minute. - HOUR: Time unit is per hour. - DAY: Time unit is per day.
type TimeUnit struct {
}

// AssertTimeUnitRequired checks if the required fields are not zero-ed
func AssertTimeUnitRequired(obj TimeUnit) error {
	return nil
}

// AssertRecurseTimeUnitRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TimeUnit (e.g. [][]TimeUnit), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTimeUnitRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTimeUnit, ok := obj.(TimeUnit)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTimeUnitRequired(aTimeUnit)
	})
}