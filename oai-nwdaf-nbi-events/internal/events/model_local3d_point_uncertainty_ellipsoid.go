/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// Local3dPointUncertaintyEllipsoid - Local 3D point with uncertainty ellipsoid
type Local3dPointUncertaintyEllipsoid struct {
	GadShape

	LocalOrigin LocalOrigin `json:"localOrigin"`

	Point RelativeCartesianLocation `json:"point"`

	UncertaintyEllipsoid UncertaintyEllipsoid `json:"uncertaintyEllipsoid"`

	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}

// AssertLocal3dPointUncertaintyEllipsoidRequired checks if the required fields are not zero-ed
func AssertLocal3dPointUncertaintyEllipsoidRequired(obj Local3dPointUncertaintyEllipsoid) error {
	elements := map[string]interface{}{
		"localOrigin":          obj.LocalOrigin,
		"point":                obj.Point,
		"uncertaintyEllipsoid": obj.UncertaintyEllipsoid,
		"confidence":           obj.Confidence,
		"shape":                obj.Shape,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertGadShapeRequired(obj.GadShape); err != nil {
		return err
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

// AssertRecurseLocal3dPointUncertaintyEllipsoidRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Local3dPointUncertaintyEllipsoid (e.g. [][]Local3dPointUncertaintyEllipsoid), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseLocal3dPointUncertaintyEllipsoidRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aLocal3dPointUncertaintyEllipsoid, ok := obj.(Local3dPointUncertaintyEllipsoid)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertLocal3dPointUncertaintyEllipsoidRequired(aLocal3dPointUncertaintyEllipsoid)
	})
}