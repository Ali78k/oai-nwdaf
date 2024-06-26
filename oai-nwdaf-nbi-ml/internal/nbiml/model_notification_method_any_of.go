/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.0.0-alpha.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nbiml

type NotificationMethodAnyOf string

// List of NotificationMethodAnyOf
const (
	NOTIFICATIONMETHODANYOF_PERIODIC           NotificationMethodAnyOf = "PERIODIC"
	NOTIFICATIONMETHODANYOF_ONE_TIME           NotificationMethodAnyOf = "ONE_TIME"
	NOTIFICATIONMETHODANYOF_ON_EVENT_DETECTION NotificationMethodAnyOf = "ON_EVENT_DETECTION"
)

// AssertNotificationMethodAnyOfRequired checks if the required fields are not zero-ed
func AssertNotificationMethodAnyOfRequired(obj NotificationMethodAnyOf) error {
	return nil
}

// AssertRecurseNotificationMethodAnyOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of NotificationMethodAnyOf (e.g. [][]NotificationMethodAnyOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNotificationMethodAnyOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNotificationMethodAnyOf, ok := obj.(NotificationMethodAnyOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNotificationMethodAnyOfRequired(aNotificationMethodAnyOf)
	})
}
