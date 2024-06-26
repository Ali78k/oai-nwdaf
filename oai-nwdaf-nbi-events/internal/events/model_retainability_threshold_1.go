/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// RetainabilityThreshold1 - Represents a QoS flow retainability threshold.
type RetainabilityThreshold1 struct {

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	RelFlowNum int32 `json:"relFlowNum,omitempty"`

	RelTimeUnit TimeUnit `json:"relTimeUnit,omitempty"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	RelFlowRatio int32 `json:"relFlowRatio,omitempty"`
}

// AssertRetainabilityThreshold1Required checks if the required fields are not zero-ed
func AssertRetainabilityThreshold1Required(obj RetainabilityThreshold1) error {
	if err := AssertTimeUnitRequired(obj.RelTimeUnit); err != nil {
		return err
	}
	return nil
}

// AssertRecurseRetainabilityThreshold1Required recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of RetainabilityThreshold1 (e.g. [][]RetainabilityThreshold1), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseRetainabilityThreshold1Required(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aRetainabilityThreshold1, ok := obj.(RetainabilityThreshold1)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertRetainabilityThreshold1Required(aRetainabilityThreshold1)
	})
}
