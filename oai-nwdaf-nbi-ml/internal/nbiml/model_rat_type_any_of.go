/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.0.0-alpha.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nbiml

type RatTypeAnyOf string

// List of RatTypeAnyOf
const (
	RATTYPEANYOF_NR             RatTypeAnyOf = "NR"
	RATTYPEANYOF_EUTRA          RatTypeAnyOf = "EUTRA"
	RATTYPEANYOF_WLAN           RatTypeAnyOf = "WLAN"
	RATTYPEANYOF_VIRTUAL        RatTypeAnyOf = "VIRTUAL"
	RATTYPEANYOF_NBIOT          RatTypeAnyOf = "NBIOT"
	RATTYPEANYOF_WIRELINE       RatTypeAnyOf = "WIRELINE"
	RATTYPEANYOF_WIRELINE_CABLE RatTypeAnyOf = "WIRELINE_CABLE"
	RATTYPEANYOF_WIRELINE_BBF   RatTypeAnyOf = "WIRELINE_BBF"
	RATTYPEANYOF_LTE_M          RatTypeAnyOf = "LTE-M"
	RATTYPEANYOF_NR_U           RatTypeAnyOf = "NR_U"
	RATTYPEANYOF_EUTRA_U        RatTypeAnyOf = "EUTRA_U"
	RATTYPEANYOF_TRUSTED_N3_GA  RatTypeAnyOf = "TRUSTED_N3GA"
	RATTYPEANYOF_TRUSTED_WLAN   RatTypeAnyOf = "TRUSTED_WLAN"
	RATTYPEANYOF_UTRA           RatTypeAnyOf = "UTRA"
	RATTYPEANYOF_GERA           RatTypeAnyOf = "GERA"
	RATTYPEANYOF_NR_LEO         RatTypeAnyOf = "NR_LEO"
	RATTYPEANYOF_NR_MEO         RatTypeAnyOf = "NR_MEO"
	RATTYPEANYOF_NR_GEO         RatTypeAnyOf = "NR_GEO"
	RATTYPEANYOF_NR_OTHER_SAT   RatTypeAnyOf = "NR_OTHER_SAT"
	RATTYPEANYOF_NR_REDCAP      RatTypeAnyOf = "NR_REDCAP"
)

// AssertRatTypeAnyOfRequired checks if the required fields are not zero-ed
func AssertRatTypeAnyOfRequired(obj RatTypeAnyOf) error {
	return nil
}

// AssertRecurseRatTypeAnyOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of RatTypeAnyOf (e.g. [][]RatTypeAnyOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseRatTypeAnyOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aRatTypeAnyOf, ok := obj.(RatTypeAnyOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertRatTypeAnyOfRequired(aRatTypeAnyOf)
	})
}
