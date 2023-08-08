/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// RedundantTransmissionExpInfo - The redundant transmission experience related information. When subscribed event is \"RED_TRANS_EXP\", the \"redTransInfos\" attribute shall be included.
type RedundantTransmissionExpInfo struct {
	SpatialValidCon NetworkAreaInfo `json:"spatialValidCon,omitempty"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003; it shall contain either a DNN Network Identifier, or a full DNN with both the Network Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots (e.g. \"Label1.Label2.Label3\").
	Dnn string `json:"dnn,omitempty"`

	RedTransExps []RedundantTransmissionExpPerTs `json:"redTransExps"`
}

// AssertRedundantTransmissionExpInfoRequired checks if the required fields are not zero-ed
func AssertRedundantTransmissionExpInfoRequired(obj RedundantTransmissionExpInfo) error {
	elements := map[string]interface{}{
		"redTransExps": obj.RedTransExps,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertNetworkAreaInfoRequired(obj.SpatialValidCon); err != nil {
		return err
	}
	for _, el := range obj.RedTransExps {
		if err := AssertRedundantTransmissionExpPerTsRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseRedundantTransmissionExpInfoRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of RedundantTransmissionExpInfo (e.g. [][]RedundantTransmissionExpInfo), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseRedundantTransmissionExpInfoRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aRedundantTransmissionExpInfo, ok := obj.(RedundantTransmissionExpInfo)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertRedundantTransmissionExpInfoRequired(aRedundantTransmissionExpInfo)
	})
}
