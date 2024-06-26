/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

type NetworkPerfTypeAnyOf string

// List of NetworkPerfTypeAnyOf
const (
	NETWORKPERFTYPEANYOF_GNB_ACTIVE_RATIO    NetworkPerfTypeAnyOf = "GNB_ACTIVE_RATIO"
	NETWORKPERFTYPEANYOF_GNB_COMPUTING_USAGE NetworkPerfTypeAnyOf = "GNB_COMPUTING_USAGE"
	NETWORKPERFTYPEANYOF_GNB_MEMORY_USAGE    NetworkPerfTypeAnyOf = "GNB_MEMORY_USAGE"
	NETWORKPERFTYPEANYOF_GNB_DISK_USAGE      NetworkPerfTypeAnyOf = "GNB_DISK_USAGE"
	NETWORKPERFTYPEANYOF_NUM_OF_UE           NetworkPerfTypeAnyOf = "NUM_OF_UE"
	NETWORKPERFTYPEANYOF_SESS_SUCC_RATIO     NetworkPerfTypeAnyOf = "SESS_SUCC_RATIO"
	NETWORKPERFTYPEANYOF_HO_SUCC_RATIO       NetworkPerfTypeAnyOf = "HO_SUCC_RATIO"
)

// AssertNetworkPerfTypeAnyOfRequired checks if the required fields are not zero-ed
func AssertNetworkPerfTypeAnyOfRequired(obj NetworkPerfTypeAnyOf) error {
	return nil
}

// AssertRecurseNetworkPerfTypeAnyOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of NetworkPerfTypeAnyOf (e.g. [][]NetworkPerfTypeAnyOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNetworkPerfTypeAnyOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNetworkPerfTypeAnyOf, ok := obj.(NetworkPerfTypeAnyOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNetworkPerfTypeAnyOfRequired(aNetworkPerfTypeAnyOf)
	})
}
