/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

// NotificationFlag - Possible values are: - ACTIVATE: The event notification is activated. - DEACTIVATE: The event notification is deactivated and shall be muted. The available event(s) shall be stored. - RETRIEVAL: The event notification shall be sent to the NF service consumer(s), after that, is muted again.
type NotificationFlag struct {
}

// AssertNotificationFlagRequired checks if the required fields are not zero-ed
func AssertNotificationFlagRequired(obj NotificationFlag) error {
	return nil
}

// AssertRecurseNotificationFlagRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of NotificationFlag (e.g. [][]NotificationFlag), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNotificationFlagRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNotificationFlag, ok := obj.(NotificationFlag)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNotificationFlagRequired(aNotificationFlag)
	})
}
