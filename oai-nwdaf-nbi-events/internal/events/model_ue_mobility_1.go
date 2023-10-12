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

// UeMobility1 - Represents UE mobility information.
type UeMobility1 struct {

	// string with format 'date-time' as defined in OpenAPI.
	Ts time.Time `json:"ts,omitempty"`

	RecurringTime ScheduledCommunicationTime `json:"recurringTime,omitempty"`

	// indicating a time in seconds.
	Duration int32 `json:"duration"`

	// string with format 'float' as defined in OpenAPI.
	DurationVariance float32 `json:"durationVariance,omitempty"`

	LocInfos []LocationInfo1 `json:"locInfos"`
}

// AssertUeMobility1Required checks if the required fields are not zero-ed
func AssertUeMobility1Required(obj UeMobility1) error {
	elements := map[string]interface{}{
		"duration": obj.Duration,
		"locInfos": obj.LocInfos,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertScheduledCommunicationTimeRequired(obj.RecurringTime); err != nil {
		return err
	}
	for _, el := range obj.LocInfos {
		if err := AssertLocationInfo1Required(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseUeMobility1Required recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of UeMobility1 (e.g. [][]UeMobility1), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUeMobility1Required(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUeMobility1, ok := obj.(UeMobility1)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUeMobility1Required(aUeMobility1)
	})
}