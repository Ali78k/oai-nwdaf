/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.0.0-alpha.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nbiml

// NfType - NF types known to NRF
type NfType struct {
}

// AssertNfTypeRequired checks if the required fields are not zero-ed
func AssertNfTypeRequired(obj NfType) error {
	return nil
}

// AssertRecurseNfTypeRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of NfType (e.g. [][]NfType), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNfTypeRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNfType, ok := obj.(NfType)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNfTypeRequired(aNfType)
	})
}
