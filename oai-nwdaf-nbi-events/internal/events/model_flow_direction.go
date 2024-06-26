/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// FlowDirection - Possible values are - DOWNLINK: The corresponding filter applies for traffic to the UE. - UPLINK: The corresponding filter applies for traffic from the UE. - BIDIRECTIONAL: The corresponding filter applies for traffic both to and from the UE. - UNSPECIFIED: The corresponding filter applies for traffic to the UE (downlink), but has no specific direction declared. The service data flow detection shall apply the filter for uplink traffic as if the filter was bidirectional. The PCF shall not use the value UNSPECIFIED in filters created by the network in NW-initiated procedures. The PCF shall only include the value UNSPECIFIED in filters in UE-initiated procedures if the same value is received from the SMF.
type FlowDirection struct {
}

// AssertFlowDirectionRequired checks if the required fields are not zero-ed
func AssertFlowDirectionRequired(obj FlowDirection) error {
	return nil
}

// AssertRecurseFlowDirectionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of FlowDirection (e.g. [][]FlowDirection), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseFlowDirectionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aFlowDirection, ok := obj.(FlowDirection)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertFlowDirectionRequired(aFlowDirection)
	})
}
