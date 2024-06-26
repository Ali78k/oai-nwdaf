/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

// DispersionOrderingCriterion - Possible values are   - TIME_SLOT_START: Indicates the order of time slot start.   - DISPERSION: Indicates the order of data/transaction dispersion.   - CLASSIFICATION: Indicates the order of data/transaction classification.   - RANKING: Indicates the order of data/transaction ranking.   - PERCENTILE_RANKING: Indicates the order of data/transaction percentile ranking.
type DispersionOrderingCriterion struct {
}

// AssertDispersionOrderingCriterionRequired checks if the required fields are not zero-ed
func AssertDispersionOrderingCriterionRequired(obj DispersionOrderingCriterion) error {
	return nil
}

// AssertRecurseDispersionOrderingCriterionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of DispersionOrderingCriterion (e.g. [][]DispersionOrderingCriterion), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseDispersionOrderingCriterionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aDispersionOrderingCriterion, ok := obj.(DispersionOrderingCriterion)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertDispersionOrderingCriterionRequired(aDispersionOrderingCriterion)
	})
}
