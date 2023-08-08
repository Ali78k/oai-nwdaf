/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.0.0-alpha.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nbiml

// RelativeCartesianLocation - Relative Cartesian Location
type RelativeCartesianLocation struct {

	// string with format 'float' as defined in OpenAPI.
	X float32 `json:"x"`

	// string with format 'float' as defined in OpenAPI.
	Y float32 `json:"y"`

	// string with format 'float' as defined in OpenAPI.
	Z float32 `json:"z,omitempty"`
}

// AssertRelativeCartesianLocationRequired checks if the required fields are not zero-ed
func AssertRelativeCartesianLocationRequired(obj RelativeCartesianLocation) error {
	elements := map[string]interface{}{
		"x": obj.X,
		"y": obj.Y,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseRelativeCartesianLocationRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of RelativeCartesianLocation (e.g. [][]RelativeCartesianLocation), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseRelativeCartesianLocationRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aRelativeCartesianLocation, ok := obj.(RelativeCartesianLocation)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertRelativeCartesianLocationRequired(aRelativeCartesianLocation)
	})
}
