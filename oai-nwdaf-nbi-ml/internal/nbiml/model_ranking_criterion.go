/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.0.0-alpha.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nbiml

// RankingCriterion - Indicates the usage ranking criterion between the high, medium and low usage UE.
type RankingCriterion struct {

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	HighBase int32 `json:"highBase"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	LowBase int32 `json:"lowBase,omitempty"`
}

// AssertRankingCriterionRequired checks if the required fields are not zero-ed
func AssertRankingCriterionRequired(obj RankingCriterion) error {
	elements := map[string]interface{}{
		"highBase": obj.HighBase,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseRankingCriterionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of RankingCriterion (e.g. [][]RankingCriterion), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseRankingCriterionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aRankingCriterion, ok := obj.(RankingCriterion)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertRankingCriterionRequired(aRankingCriterion)
	})
}