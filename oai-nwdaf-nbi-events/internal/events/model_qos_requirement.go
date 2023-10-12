/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// QosRequirement - Represents the QoS requirements.
type QosRequirement struct {

	// Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255
	Var5qi int32 `json:"5qi,omitempty"`

	// String representing a bit rate prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	GfbrUl string `json:"gfbrUl,omitempty"`

	// String representing a bit rate prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	GfbrDl string `json:"gfbrDl,omitempty"`

	ResType QosResourceType `json:"resType,omitempty"`

	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	Pdb int32 `json:"pdb,omitempty"`

	// String representing Packet Error Rate (see clause 5.7.3.5 and 5.7.4 of 3GPP TS 23.501, expressed as a \"scalar x 10-k\" where the scalar and the exponent k are each encoded as one decimal digit.
	Per string `json:"per,omitempty"`
}

// AssertQosRequirementRequired checks if the required fields are not zero-ed
func AssertQosRequirementRequired(obj QosRequirement) error {
	if err := AssertQosResourceTypeRequired(obj.ResType); err != nil {
		return err
	}
	return nil
}

// AssertRecurseQosRequirementRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of QosRequirement (e.g. [][]QosRequirement), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseQosRequirementRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aQosRequirement, ok := obj.(QosRequirement)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertQosRequirementRequired(aQosRequirement)
	})
}