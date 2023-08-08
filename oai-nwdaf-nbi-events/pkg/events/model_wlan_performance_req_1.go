/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// WlanPerformanceReq1 - Represents other WLAN performance analytics requirements.
type WlanPerformanceReq1 struct {
	SsIds []string `json:"ssIds,omitempty"`

	BssIds []string `json:"bssIds,omitempty"`

	WlanOrderCriter WlanOrderingCriterion `json:"wlanOrderCriter,omitempty"`

	Order MatchingDirection `json:"order,omitempty"`
}

// AssertWlanPerformanceReq1Required checks if the required fields are not zero-ed
func AssertWlanPerformanceReq1Required(obj WlanPerformanceReq1) error {
	if err := AssertWlanOrderingCriterionRequired(obj.WlanOrderCriter); err != nil {
		return err
	}
	if err := AssertMatchingDirectionRequired(obj.Order); err != nil {
		return err
	}
	return nil
}

// AssertRecurseWlanPerformanceReq1Required recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of WlanPerformanceReq1 (e.g. [][]WlanPerformanceReq1), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseWlanPerformanceReq1Required(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aWlanPerformanceReq1, ok := obj.(WlanPerformanceReq1)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertWlanPerformanceReq1Required(aWlanPerformanceReq1)
	})
}
