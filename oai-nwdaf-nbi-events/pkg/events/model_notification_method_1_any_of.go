/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

type NotificationMethod1AnyOf string

// List of NotificationMethod1AnyOf
const (
	NOTIFICATIONMETHOD1ANYOF_PERIODIC           NotificationMethod1AnyOf = "PERIODIC"
	NOTIFICATIONMETHOD1ANYOF_ONE_TIME           NotificationMethod1AnyOf = "ONE_TIME"
	NOTIFICATIONMETHOD1ANYOF_ON_EVENT_DETECTION NotificationMethod1AnyOf = "ON_EVENT_DETECTION"
)

// AssertNotificationMethod1AnyOfRequired checks if the required fields are not zero-ed
func AssertNotificationMethod1AnyOfRequired(obj NotificationMethod1AnyOf) error {
	return nil
}

// AssertRecurseNotificationMethod1AnyOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of NotificationMethod1AnyOf (e.g. [][]NotificationMethod1AnyOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNotificationMethod1AnyOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNotificationMethod1AnyOf, ok := obj.(NotificationMethod1AnyOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNotificationMethod1AnyOfRequired(aNotificationMethod1AnyOf)
	})
}
