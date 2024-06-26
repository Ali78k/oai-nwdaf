/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// IpEthFlowDescription1 - Contains the description of an Uplink and/or Downlink Ethernet flow.
type IpEthFlowDescription1 struct {

	// Defines a packet filter of an IP flow.
	IpTrafficFilter string `json:"ipTrafficFilter,omitempty"`

	EthTrafficFilter EthFlowDescription `json:"ethTrafficFilter,omitempty"`
}

// AssertIpEthFlowDescription1Required checks if the required fields are not zero-ed
func AssertIpEthFlowDescription1Required(obj IpEthFlowDescription1) error {
	if err := AssertEthFlowDescriptionRequired(obj.EthTrafficFilter); err != nil {
		return err
	}
	return nil
}

// AssertRecurseIpEthFlowDescription1Required recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of IpEthFlowDescription1 (e.g. [][]IpEthFlowDescription1), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseIpEthFlowDescription1Required(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aIpEthFlowDescription1, ok := obj.(IpEthFlowDescription1)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertIpEthFlowDescription1Required(aIpEthFlowDescription1)
	})
}
