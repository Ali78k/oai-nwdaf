/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

type Local3dPointUncertaintyEllipsoidAllOf struct {
	LocalOrigin LocalOrigin `json:"localOrigin"`

	Point RelativeCartesianLocation `json:"point"`

	UncertaintyEllipsoid UncertaintyEllipsoid `json:"uncertaintyEllipsoid"`

	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}

// AssertLocal3dPointUncertaintyEllipsoidAllOfRequired checks if the required fields are not zero-ed
func AssertLocal3dPointUncertaintyEllipsoidAllOfRequired(obj Local3dPointUncertaintyEllipsoidAllOf) error {
	elements := map[string]interface{}{
		"localOrigin":          obj.LocalOrigin,
		"point":                obj.Point,
		"uncertaintyEllipsoid": obj.UncertaintyEllipsoid,
		"confidence":           obj.Confidence,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertLocalOriginRequired(obj.LocalOrigin); err != nil {
		return err
	}
	if err := AssertRelativeCartesianLocationRequired(obj.Point); err != nil {
		return err
	}
	if err := AssertUncertaintyEllipsoidRequired(obj.UncertaintyEllipsoid); err != nil {
		return err
	}
	return nil
}

// AssertRecurseLocal3dPointUncertaintyEllipsoidAllOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Local3dPointUncertaintyEllipsoidAllOf (e.g. [][]Local3dPointUncertaintyEllipsoidAllOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseLocal3dPointUncertaintyEllipsoidAllOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aLocal3dPointUncertaintyEllipsoidAllOf, ok := obj.(Local3dPointUncertaintyEllipsoidAllOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertLocal3dPointUncertaintyEllipsoidAllOfRequired(aLocal3dPointUncertaintyEllipsoidAllOf)
	})
}
