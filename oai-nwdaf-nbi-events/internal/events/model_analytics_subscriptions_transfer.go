/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// AnalyticsSubscriptionsTransfer - Contains information about a request to transfer analytics subscriptions.
type AnalyticsSubscriptionsTransfer struct {
	SubsTransInfos []SubscriptionTransferInfo `json:"subsTransInfos"`
}

// AssertAnalyticsSubscriptionsTransferRequired checks if the required fields are not zero-ed
func AssertAnalyticsSubscriptionsTransferRequired(obj AnalyticsSubscriptionsTransfer) error {
	elements := map[string]interface{}{
		"subsTransInfos": obj.SubsTransInfos,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.SubsTransInfos {
		if err := AssertSubscriptionTransferInfoRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseAnalyticsSubscriptionsTransferRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AnalyticsSubscriptionsTransfer (e.g. [][]AnalyticsSubscriptionsTransfer), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAnalyticsSubscriptionsTransferRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAnalyticsSubscriptionsTransfer, ok := obj.(AnalyticsSubscriptionsTransfer)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAnalyticsSubscriptionsTransferRequired(aAnalyticsSubscriptionsTransfer)
	})
}
