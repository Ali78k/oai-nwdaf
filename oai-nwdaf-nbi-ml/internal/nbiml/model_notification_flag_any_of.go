/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.0.0-alpha.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nbiml

type NotificationFlagAnyOf string

// List of NotificationFlagAnyOf
const (
	NOTIFICATIONFLAGANYOF_ACTIVATE   NotificationFlagAnyOf = "ACTIVATE"
	NOTIFICATIONFLAGANYOF_DEACTIVATE NotificationFlagAnyOf = "DEACTIVATE"
	NOTIFICATIONFLAGANYOF_RETRIEVAL  NotificationFlagAnyOf = "RETRIEVAL"
)

// AssertNotificationFlagAnyOfRequired checks if the required fields are not zero-ed
func AssertNotificationFlagAnyOfRequired(obj NotificationFlagAnyOf) error {
	return nil
}

// AssertRecurseNotificationFlagAnyOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of NotificationFlagAnyOf (e.g. [][]NotificationFlagAnyOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNotificationFlagAnyOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNotificationFlagAnyOf, ok := obj.(NotificationFlagAnyOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNotificationFlagAnyOfRequired(aNotificationFlagAnyOf)
	})
}