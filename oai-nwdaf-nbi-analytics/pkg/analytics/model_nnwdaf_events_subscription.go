/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

// NnwdafEventsSubscription - Represents an Individual NWDAF Event Subscription resource.
type NnwdafEventsSubscription struct {

	// Subscribed events
	EventSubscriptions []EventSubscription `json:"eventSubscriptions"`

	EvtReq ReportingInformation `json:"evtReq,omitempty"`

	// String providing an URI formatted according to RFC 3986
	NotificationURI string `json:"notificationURI,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	EventNotifications []EventNotification `json:"eventNotifications,omitempty"`

	FailEventReports []FailureEventInfo `json:"failEventReports,omitempty"`

	PrevSub SpecificAnalyticsSubscription1 `json:"prevSub,omitempty"`

	ConsNfInfo ConsumerNfInformation `json:"consNfInfo,omitempty"`
}

// AssertNnwdafEventsSubscriptionRequired checks if the required fields are not zero-ed
func AssertNnwdafEventsSubscriptionRequired(obj NnwdafEventsSubscription) error {
	elements := map[string]interface{}{
		"eventSubscriptions": obj.EventSubscriptions,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.EventSubscriptions {
		if err := AssertEventSubscriptionRequired(el); err != nil {
			return err
		}
	}
	if err := AssertReportingInformationRequired(obj.EvtReq); err != nil {
		return err
	}
	for _, el := range obj.EventNotifications {
		if err := AssertEventNotificationRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.FailEventReports {
		if err := AssertFailureEventInfoRequired(el); err != nil {
			return err
		}
	}
	if err := AssertSpecificAnalyticsSubscription1Required(obj.PrevSub); err != nil {
		return err
	}
	if err := AssertConsumerNfInformationRequired(obj.ConsNfInfo); err != nil {
		return err
	}
	return nil
}

// AssertRecurseNnwdafEventsSubscriptionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of NnwdafEventsSubscription (e.g. [][]NnwdafEventsSubscription), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNnwdafEventsSubscriptionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNnwdafEventsSubscription, ok := obj.(NnwdafEventsSubscription)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNnwdafEventsSubscriptionRequired(aNnwdafEventsSubscription)
	})
}
