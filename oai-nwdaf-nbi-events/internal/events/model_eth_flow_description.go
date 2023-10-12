/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// EthFlowDescription - Identifies an Ethernet flow.
type EthFlowDescription struct {

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042
	DestMacAddr string `json:"destMacAddr,omitempty"`

	EthType string `json:"ethType"`

	// Defines a packet filter of an IP flow.
	FDesc string `json:"fDesc,omitempty"`

	FDir FlowDirection `json:"fDir,omitempty"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042
	SourceMacAddr string `json:"sourceMacAddr,omitempty"`

	VlanTags []string `json:"vlanTags,omitempty"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042
	SrcMacAddrEnd string `json:"srcMacAddrEnd,omitempty"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042
	DestMacAddrEnd string `json:"destMacAddrEnd,omitempty"`
}

// AssertEthFlowDescriptionRequired checks if the required fields are not zero-ed
func AssertEthFlowDescriptionRequired(obj EthFlowDescription) error {
	elements := map[string]interface{}{
		"ethType": obj.EthType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertFlowDirectionRequired(obj.FDir); err != nil {
		return err
	}
	return nil
}

// AssertRecurseEthFlowDescriptionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of EthFlowDescription (e.g. [][]EthFlowDescription), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseEthFlowDescriptionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aEthFlowDescription, ok := obj.(EthFlowDescription)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertEthFlowDescriptionRequired(aEthFlowDescription)
	})
}