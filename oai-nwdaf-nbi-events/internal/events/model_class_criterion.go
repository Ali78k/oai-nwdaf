/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// ClassCriterion - Indicates the dispersion class criterion for fixed, camper and/or traveller UE, and/or the top-heavy UE dispersion class criterion.
type ClassCriterion struct {
	DisperClass DispersionClass `json:"disperClass"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	ClassThreshold int32 `json:"classThreshold"`

	ThresMatch MatchingDirection `json:"thresMatch"`
}

// AssertClassCriterionRequired checks if the required fields are not zero-ed
func AssertClassCriterionRequired(obj ClassCriterion) error {
	elements := map[string]interface{}{
		"disperClass":    obj.DisperClass,
		"classThreshold": obj.ClassThreshold,
		"thresMatch":     obj.ThresMatch,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertDispersionClassRequired(obj.DisperClass); err != nil {
		return err
	}
	if err := AssertMatchingDirectionRequired(obj.ThresMatch); err != nil {
		return err
	}
	return nil
}

// AssertRecurseClassCriterionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ClassCriterion (e.g. [][]ClassCriterion), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseClassCriterionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aClassCriterion, ok := obj.(ClassCriterion)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertClassCriterionRequired(aClassCriterion)
	})
}