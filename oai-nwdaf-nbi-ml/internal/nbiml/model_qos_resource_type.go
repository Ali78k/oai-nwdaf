/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.0.0-alpha.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nbiml

// QosResourceType - The enumeration QosResourceType indicates whether a QoS Flow is non-GBR, delay critical GBR, or non-delay critical GBR (see clauses 5.7.3.4 and 5.7.3.5 of 3GPP TS 23.501). It shall comply with the provisions defined in table 5.5.3.6-1.
type QosResourceType struct {
}

// AssertQosResourceTypeRequired checks if the required fields are not zero-ed
func AssertQosResourceTypeRequired(obj QosResourceType) error {
	return nil
}

// AssertRecurseQosResourceTypeRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of QosResourceType (e.g. [][]QosResourceType), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseQosResourceTypeRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aQosResourceType, ok := obj.(QosResourceType)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertQosResourceTypeRequired(aQosResourceType)
	})
}