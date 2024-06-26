/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.0.0-alpha.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nbiml

// LocalOrigin - Indicates a Local origin in a reference system
type LocalOrigin struct {
	CoordinateId string `json:"coordinateId,omitempty"`

	Point GeographicalCoordinates `json:"point,omitempty"`
}

// AssertLocalOriginRequired checks if the required fields are not zero-ed
func AssertLocalOriginRequired(obj LocalOrigin) error {
	if err := AssertGeographicalCoordinatesRequired(obj.Point); err != nil {
		return err
	}
	return nil
}

// AssertRecurseLocalOriginRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of LocalOrigin (e.g. [][]LocalOrigin), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseLocalOriginRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aLocalOrigin, ok := obj.(LocalOrigin)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertLocalOriginRequired(aLocalOrigin)
	})
}
