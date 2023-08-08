/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

// SmcceInfo - Represents the Session Management congestion control experience information.
type SmcceInfo struct {

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003; it shall contain either a DNN Network Identifier, or a full DNN with both the Network Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots (e.g. \"Label1.Label2.Label3\").
	Dnn string `json:"dnn,omitempty"`

	Snssai Snssai `json:"snssai,omitempty"`

	SmcceUeList []SmcceUeList `json:"smcceUeList,omitempty"`
}

// AssertSmcceInfoRequired checks if the required fields are not zero-ed
func AssertSmcceInfoRequired(obj SmcceInfo) error {
	if err := AssertSnssaiRequired(obj.Snssai); err != nil {
		return err
	}
	for _, el := range obj.SmcceUeList {
		if err := AssertSmcceUeListRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseSmcceInfoRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of SmcceInfo (e.g. [][]SmcceInfo), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseSmcceInfoRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aSmcceInfo, ok := obj.(SmcceInfo)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertSmcceInfoRequired(aSmcceInfo)
	})
}
