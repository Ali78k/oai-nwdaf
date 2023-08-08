/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

type NwdafEventAnyOf string

// List of NwdafEventAnyOf
const (
	NWDAFEVENTANYOF_SLICE_LOAD_LEVEL     NwdafEventAnyOf = "SLICE_LOAD_LEVEL"
	NWDAFEVENTANYOF_NETWORK_PERFORMANCE  NwdafEventAnyOf = "NETWORK_PERFORMANCE"
	NWDAFEVENTANYOF_NF_LOAD              NwdafEventAnyOf = "NF_LOAD"
	NWDAFEVENTANYOF_SERVICE_EXPERIENCE   NwdafEventAnyOf = "SERVICE_EXPERIENCE"
	NWDAFEVENTANYOF_UE_MOBILITY          NwdafEventAnyOf = "UE_MOBILITY"
	NWDAFEVENTANYOF_UE_COMMUNICATION     NwdafEventAnyOf = "UE_COMMUNICATION"
	NWDAFEVENTANYOF_QOS_SUSTAINABILITY   NwdafEventAnyOf = "QOS_SUSTAINABILITY"
	NWDAFEVENTANYOF_ABNORMAL_BEHAVIOUR   NwdafEventAnyOf = "ABNORMAL_BEHAVIOUR"
	NWDAFEVENTANYOF_USER_DATA_CONGESTION NwdafEventAnyOf = "USER_DATA_CONGESTION"
	NWDAFEVENTANYOF_NSI_LOAD_LEVEL       NwdafEventAnyOf = "NSI_LOAD_LEVEL"
	NWDAFEVENTANYOF_DN_PERFORMANCE       NwdafEventAnyOf = "DN_PERFORMANCE"
	NWDAFEVENTANYOF_DISPERSION           NwdafEventAnyOf = "DISPERSION"
	NWDAFEVENTANYOF_RED_TRANS_EXP        NwdafEventAnyOf = "RED_TRANS_EXP"
	NWDAFEVENTANYOF_WLAN_PERFORMANCE     NwdafEventAnyOf = "WLAN_PERFORMANCE"
)

// AssertNwdafEventAnyOfRequired checks if the required fields are not zero-ed
func AssertNwdafEventAnyOfRequired(obj NwdafEventAnyOf) error {
	return nil
}

// AssertRecurseNwdafEventAnyOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of NwdafEventAnyOf (e.g. [][]NwdafEventAnyOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNwdafEventAnyOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNwdafEventAnyOf, ok := obj.(NwdafEventAnyOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNwdafEventAnyOfRequired(aNwdafEventAnyOf)
	})
}
