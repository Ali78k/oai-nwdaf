/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// AbnormalBehaviour1 - Represents the abnormal behaviour information.
type AbnormalBehaviour1 struct {
	Supis []string `json:"supis,omitempty"`

	Excep Exception1 `json:"excep"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003; it shall contain either a DNN Network Identifier, or a full DNN with both the Network Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots (e.g. \"Label1.Label2.Label3\").
	Dnn string `json:"dnn,omitempty"`

	Snssai Snssai `json:"snssai,omitempty"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	Ratio int32 `json:"ratio,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty"`

	AddtMeasInfo AdditionalMeasurement1 `json:"addtMeasInfo,omitempty"`
}

// AssertAbnormalBehaviour1Required checks if the required fields are not zero-ed
func AssertAbnormalBehaviour1Required(obj AbnormalBehaviour1) error {
	elements := map[string]interface{}{
		"excep": obj.Excep,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertException1Required(obj.Excep); err != nil {
		return err
	}
	if err := AssertSnssaiRequired(obj.Snssai); err != nil {
		return err
	}
	if err := AssertAdditionalMeasurement1Required(obj.AddtMeasInfo); err != nil {
		return err
	}
	return nil
}

// AssertRecurseAbnormalBehaviour1Required recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AbnormalBehaviour1 (e.g. [][]AbnormalBehaviour1), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAbnormalBehaviour1Required(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAbnormalBehaviour1, ok := obj.(AbnormalBehaviour1)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAbnormalBehaviour1Required(aAbnormalBehaviour1)
	})
}
