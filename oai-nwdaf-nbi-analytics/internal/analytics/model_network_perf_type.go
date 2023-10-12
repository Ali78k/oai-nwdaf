/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

// NetworkPerfType - Possible values are   - GNB_ACTIVE_RATIO: Indicates that the network performance requirement is gNodeB active (i.e. up and running) rate. Indicates the ratio of gNB active (i.e. up and running) number to the total number of gNB   - GNB_COMPUTING_USAGE: Indicates gNodeB computing resource usage.   - GNB_MEMORY_USAGE: Indicates gNodeB memory usage.   - GNB_DISK_USAGE: Indicates gNodeB disk usage.   - NUM_OF_UE: Indicates number of UEs.   - SESS_SUCC_RATIO: Indicates ratio of successful setup of PDU sessions to total PDU session setup attempts.   - HO_SUCC_RATIO: Indicates Ratio of successful handovers to the total handover attempts.
type NetworkPerfType struct {
}

// AssertNetworkPerfTypeRequired checks if the required fields are not zero-ed
func AssertNetworkPerfTypeRequired(obj NetworkPerfType) error {
	return nil
}

// AssertRecurseNetworkPerfTypeRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of NetworkPerfType (e.g. [][]NetworkPerfType), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNetworkPerfTypeRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNetworkPerfType, ok := obj.(NetworkPerfType)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNetworkPerfTypeRequired(aNetworkPerfType)
	})
}