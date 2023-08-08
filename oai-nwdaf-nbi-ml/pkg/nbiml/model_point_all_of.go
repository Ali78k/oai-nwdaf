/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.0.0-alpha.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nbiml

type PointAllOf struct {
	Point GeographicalCoordinates `json:"point"`
}

// AssertPointAllOfRequired checks if the required fields are not zero-ed
func AssertPointAllOfRequired(obj PointAllOf) error {
	elements := map[string]interface{}{
		"point": obj.Point,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertGeographicalCoordinatesRequired(obj.Point); err != nil {
		return err
	}
	return nil
}

// AssertRecursePointAllOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PointAllOf (e.g. [][]PointAllOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePointAllOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPointAllOf, ok := obj.(PointAllOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPointAllOfRequired(aPointAllOf)
	})
}
