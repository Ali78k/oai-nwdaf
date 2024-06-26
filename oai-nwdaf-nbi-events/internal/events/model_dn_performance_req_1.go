/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// DnPerformanceReq1 - Represents other DN performance analytics requirements.
type DnPerformanceReq1 struct {
	DnPerfOrderCriter DnPerfOrderingCriterion `json:"dnPerfOrderCriter,omitempty"`

	Order MatchingDirection `json:"order,omitempty"`

	ReportThresholds []ThresholdLevel1 `json:"reportThresholds,omitempty"`
}

// AssertDnPerformanceReq1Required checks if the required fields are not zero-ed
func AssertDnPerformanceReq1Required(obj DnPerformanceReq1) error {
	if err := AssertDnPerfOrderingCriterionRequired(obj.DnPerfOrderCriter); err != nil {
		return err
	}
	if err := AssertMatchingDirectionRequired(obj.Order); err != nil {
		return err
	}
	for _, el := range obj.ReportThresholds {
		if err := AssertThresholdLevel1Required(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseDnPerformanceReq1Required recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of DnPerformanceReq1 (e.g. [][]DnPerformanceReq1), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseDnPerformanceReq1Required(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aDnPerformanceReq1, ok := obj.(DnPerformanceReq1)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertDnPerformanceReq1Required(aDnPerformanceReq1)
	})
}
