/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

// RetainabilityThreshold - Represents a QoS flow retainability threshold.
type RetainabilityThreshold struct {

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	RelFlowNum int32 `json:"relFlowNum,omitempty"`

	RelTimeUnit TimeUnit `json:"relTimeUnit,omitempty"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	RelFlowRatio int32 `json:"relFlowRatio,omitempty"`
}

// AssertRetainabilityThresholdRequired checks if the required fields are not zero-ed
func AssertRetainabilityThresholdRequired(obj RetainabilityThreshold) error {
	if err := AssertTimeUnitRequired(obj.RelTimeUnit); err != nil {
		return err
	}
	return nil
}

// AssertRecurseRetainabilityThresholdRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of RetainabilityThreshold (e.g. [][]RetainabilityThreshold), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseRetainabilityThresholdRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aRetainabilityThreshold, ok := obj.(RetainabilityThreshold)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertRetainabilityThresholdRequired(aRetainabilityThreshold)
	})
}
