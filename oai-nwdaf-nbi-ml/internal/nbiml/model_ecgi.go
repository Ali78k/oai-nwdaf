/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.0.0-alpha.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nbiml

// Ecgi - Contains the ECGI (E-UTRAN Cell Global Identity), as described in 3GPP 23.003
type Ecgi struct {
	PlmnId PlmnId `json:"plmnId"`

	// 28-bit string identifying an E-UTRA Cell Id as specified in clause 9.3.1.9 of  3GPP TS 38.413, in hexadecimal representation. Each character in the string shall take a  value of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most  significant character representing the 4 most significant bits of the Cell Id shall appear  first in the string, and the character representing the 4 least significant bit of the  Cell Id shall appear last in the string.
	EutraCellId string `json:"eutraCellId"`

	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).
	Nid string `json:"nid,omitempty"`
}

// AssertEcgiRequired checks if the required fields are not zero-ed
func AssertEcgiRequired(obj Ecgi) error {
	elements := map[string]interface{}{
		"plmnId":      obj.PlmnId,
		"eutraCellId": obj.EutraCellId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertPlmnIdRequired(obj.PlmnId); err != nil {
		return err
	}
	return nil
}

// AssertRecurseEcgiRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Ecgi (e.g. [][]Ecgi), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseEcgiRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aEcgi, ok := obj.(Ecgi)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertEcgiRequired(aEcgi)
	})
}