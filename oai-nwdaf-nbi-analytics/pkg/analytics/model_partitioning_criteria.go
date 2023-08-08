/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

// PartitioningCriteria - Possible values are: - \"TAC\": Type Allocation Code - \"SUBPLMN\": Subscriber PLMN ID - \"GEOAREA\": Geographical area, i.e. list(s) of TAI(s) - \"SNSSAI\": S-NSSAI - \"DNN\": DNN
type PartitioningCriteria struct {
}

// AssertPartitioningCriteriaRequired checks if the required fields are not zero-ed
func AssertPartitioningCriteriaRequired(obj PartitioningCriteria) error {
	return nil
}

// AssertRecursePartitioningCriteriaRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PartitioningCriteria (e.g. [][]PartitioningCriteria), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePartitioningCriteriaRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPartitioningCriteria, ok := obj.(PartitioningCriteria)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPartitioningCriteriaRequired(aPartitioningCriteria)
	})
}
