/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// NnwdafEventsSubscriptionNotification - Represents an Individual NWDAF Event Subscription Notification resource.
type NnwdafEventsSubscriptionNotification struct {

	// Notifications about Individual Events
	EventNotifications []EventNotification `json:"eventNotifications"`

	// String identifying a subscription to the Nnwdaf_EventsSubscription Service
	SubscriptionId string `json:"subscriptionId"`
}

// AssertNnwdafEventsSubscriptionNotificationRequired checks if the required fields are not zero-ed
func AssertNnwdafEventsSubscriptionNotificationRequired(obj NnwdafEventsSubscriptionNotification) error {
	elements := map[string]interface{}{
		"eventNotifications": obj.EventNotifications,
		"subscriptionId":     obj.SubscriptionId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.EventNotifications {
		if err := AssertEventNotificationRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseNnwdafEventsSubscriptionNotificationRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of NnwdafEventsSubscriptionNotification (e.g. [][]NnwdafEventsSubscriptionNotification), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNnwdafEventsSubscriptionNotificationRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNnwdafEventsSubscriptionNotification, ok := obj.(NnwdafEventsSubscriptionNotification)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNnwdafEventsSubscriptionNotificationRequired(aNnwdafEventsSubscriptionNotification)
	})
}