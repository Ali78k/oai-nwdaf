/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

import (
	"time"
)

// TimeWindow - Represents a time window identified by a start time and a stop time.
type TimeWindow struct {

	// string with format \"date-time\" as defined in OpenAPI.
	StartTime time.Time `json:"startTime"`

	// string with format \"date-time\" as defined in OpenAPI.
	StopTime time.Time `json:"stopTime"`
}

// AssertTimeWindowRequired checks if the required fields are not zero-ed
func AssertTimeWindowRequired(obj TimeWindow) error {
	elements := map[string]interface{}{
		"startTime": obj.StartTime,
		"stopTime":  obj.StopTime,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseTimeWindowRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TimeWindow (e.g. [][]TimeWindow), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTimeWindowRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTimeWindow, ok := obj.(TimeWindow)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTimeWindowRequired(aTimeWindow)
	})
}
