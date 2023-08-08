/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// DispersionRequirement - Represents the dispersion analytics requirements.
type DispersionRequirement struct {
	DisperType DispersionType `json:"disperType"`

	ClassCriters []ClassCriterion `json:"classCriters,omitempty"`

	RankCriters []RankingCriterion `json:"rankCriters,omitempty"`

	DispOrderCriter DispersionOrderingCriterion `json:"dispOrderCriter,omitempty"`

	Order MatchingDirection `json:"order,omitempty"`
}

// AssertDispersionRequirementRequired checks if the required fields are not zero-ed
func AssertDispersionRequirementRequired(obj DispersionRequirement) error {
	elements := map[string]interface{}{
		"disperType": obj.DisperType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertDispersionTypeRequired(obj.DisperType); err != nil {
		return err
	}
	for _, el := range obj.ClassCriters {
		if err := AssertClassCriterionRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.RankCriters {
		if err := AssertRankingCriterionRequired(el); err != nil {
			return err
		}
	}
	if err := AssertDispersionOrderingCriterionRequired(obj.DispOrderCriter); err != nil {
		return err
	}
	if err := AssertMatchingDirectionRequired(obj.Order); err != nil {
		return err
	}
	return nil
}

// AssertRecurseDispersionRequirementRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of DispersionRequirement (e.g. [][]DispersionRequirement), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseDispersionRequirementRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aDispersionRequirement, ok := obj.(DispersionRequirement)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertDispersionRequirementRequired(aDispersionRequirement)
	})
}
