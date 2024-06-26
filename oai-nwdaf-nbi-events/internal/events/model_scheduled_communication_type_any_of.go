/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

type ScheduledCommunicationTypeAnyOf string

// List of ScheduledCommunicationTypeAnyOf
const (
	SCHEDULEDCOMMUNICATIONTYPEANYOF_DOWNLINK_ONLY ScheduledCommunicationTypeAnyOf = "DOWNLINK_ONLY"
	SCHEDULEDCOMMUNICATIONTYPEANYOF_UPLINK_ONLY   ScheduledCommunicationTypeAnyOf = "UPLINK_ONLY"
	SCHEDULEDCOMMUNICATIONTYPEANYOF_BIDIRECTIONAL ScheduledCommunicationTypeAnyOf = "BIDIRECTIONAL"
)

// AssertScheduledCommunicationTypeAnyOfRequired checks if the required fields are not zero-ed
func AssertScheduledCommunicationTypeAnyOfRequired(obj ScheduledCommunicationTypeAnyOf) error {
	return nil
}

// AssertRecurseScheduledCommunicationTypeAnyOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ScheduledCommunicationTypeAnyOf (e.g. [][]ScheduledCommunicationTypeAnyOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseScheduledCommunicationTypeAnyOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aScheduledCommunicationTypeAnyOf, ok := obj.(ScheduledCommunicationTypeAnyOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertScheduledCommunicationTypeAnyOfRequired(aScheduledCommunicationTypeAnyOf)
	})
}
