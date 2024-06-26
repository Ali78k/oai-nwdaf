/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// NnwdafEventsSubscription1 - Represents an Individual NWDAF Event Subscription resource.
type NnwdafEventsSubscription1 struct {

	// Subscribed events
	EventSubscriptions []EventSubscription1 `json:"eventSubscriptions"`

	EvtReq ReportingInformation `json:"evtReq,omitempty"`

	// String providing an URI formatted according to RFC 3986
	NotificationURI string `json:"notificationURI,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	EventNotifications []EventNotification1 `json:"eventNotifications,omitempty"`

	FailEventReports []FailureEventInfo1 `json:"failEventReports,omitempty"`

	PrevSub SpecificAnalyticsSubscription `json:"prevSub,omitempty"`

	ConsNfInfo ConsumerNfInformation1 `json:"consNfInfo,omitempty"`
}

// AssertNnwdafEventsSubscription1Required checks if the required fields are not zero-ed
func AssertNnwdafEventsSubscription1Required(obj NnwdafEventsSubscription1) error {
	elements := map[string]interface{}{
		"eventSubscriptions": obj.EventSubscriptions,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.EventSubscriptions {
		if err := AssertEventSubscription1Required(el); err != nil {
			return err
		}
	}
	if err := AssertReportingInformationRequired(obj.EvtReq); err != nil {
		return err
	}
	for _, el := range obj.EventNotifications {
		if err := AssertEventNotification1Required(el); err != nil {
			return err
		}
	}
	for _, el := range obj.FailEventReports {
		if err := AssertFailureEventInfo1Required(el); err != nil {
			return err
		}
	}
	if err := AssertSpecificAnalyticsSubscriptionRequired(obj.PrevSub); err != nil {
		return err
	}
	if err := AssertConsumerNfInformation1Required(obj.ConsNfInfo); err != nil {
		return err
	}
	return nil
}

// AssertRecurseNnwdafEventsSubscription1Required recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of NnwdafEventsSubscription1 (e.g. [][]NnwdafEventsSubscription1), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNnwdafEventsSubscription1Required(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNnwdafEventsSubscription1, ok := obj.(NnwdafEventsSubscription1)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNnwdafEventsSubscription1Required(aNnwdafEventsSubscription1)
	})
}
