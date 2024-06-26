/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

type AnalyticsMetadataAnyOf string

// List of AnalyticsMetadataAnyOf
const (
	ANALYTICSMETADATAANYOF_NUM_OF_SAMPLES  AnalyticsMetadataAnyOf = "NUM_OF_SAMPLES"
	ANALYTICSMETADATAANYOF_DATA_WINDOW     AnalyticsMetadataAnyOf = "DATA_WINDOW"
	ANALYTICSMETADATAANYOF_DATA_STAT_PROPS AnalyticsMetadataAnyOf = "DATA_STAT_PROPS"
	ANALYTICSMETADATAANYOF_STRATEGY        AnalyticsMetadataAnyOf = "STRATEGY"
	ANALYTICSMETADATAANYOF_ACCURACY        AnalyticsMetadataAnyOf = "ACCURACY"
)

// AssertAnalyticsMetadataAnyOfRequired checks if the required fields are not zero-ed
func AssertAnalyticsMetadataAnyOfRequired(obj AnalyticsMetadataAnyOf) error {
	return nil
}

// AssertRecurseAnalyticsMetadataAnyOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AnalyticsMetadataAnyOf (e.g. [][]AnalyticsMetadataAnyOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAnalyticsMetadataAnyOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAnalyticsMetadataAnyOf, ok := obj.(AnalyticsMetadataAnyOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAnalyticsMetadataAnyOfRequired(aAnalyticsMetadataAnyOf)
	})
}
