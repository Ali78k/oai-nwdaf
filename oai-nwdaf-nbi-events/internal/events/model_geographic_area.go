/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// GeographicArea - Geographic area specified by different shape.
type GeographicArea struct {
	Shape SupportedGadShapes `json:"shape"`

	Point GeographicalCoordinates `json:"point"`

	// Indicates value of uncertainty.
	Uncertainty float32 `json:"uncertainty"`

	UncertaintyEllipse UncertaintyEllipse `json:"uncertaintyEllipse"`

	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`

	// List of points.
	PointList []GeographicalCoordinates `json:"pointList"`

	// Indicates value of altitude.
	Altitude float64 `json:"altitude"`

	// Indicates value of uncertainty.
	UncertaintyAltitude float32 `json:"uncertaintyAltitude"`

	// Indicates value of the inner radius.
	InnerRadius int32 `json:"innerRadius"`

	// Indicates value of uncertainty.
	UncertaintyRadius float32 `json:"uncertaintyRadius"`

	// Indicates value of angle.
	OffsetAngle int32 `json:"offsetAngle"`

	// Indicates value of angle.
	IncludedAngle int32 `json:"includedAngle"`
}

// AssertGeographicAreaRequired checks if the required fields are not zero-ed
func AssertGeographicAreaRequired(obj GeographicArea) error {
	elements := map[string]interface{}{
		"shape":               obj.Shape,
		"point":               obj.Point,
		"uncertainty":         obj.Uncertainty,
		"uncertaintyEllipse":  obj.UncertaintyEllipse,
		"confidence":          obj.Confidence,
		"pointList":           obj.PointList,
		"altitude":            obj.Altitude,
		"uncertaintyAltitude": obj.UncertaintyAltitude,
		"innerRadius":         obj.InnerRadius,
		"uncertaintyRadius":   obj.UncertaintyRadius,
		"offsetAngle":         obj.OffsetAngle,
		"includedAngle":       obj.IncludedAngle,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertSupportedGadShapesRequired(obj.Shape); err != nil {
		return err
	}
	if err := AssertGeographicalCoordinatesRequired(obj.Point); err != nil {
		return err
	}
	if err := AssertUncertaintyEllipseRequired(obj.UncertaintyEllipse); err != nil {
		return err
	}
	for _, el := range obj.PointList {
		if err := AssertGeographicalCoordinatesRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseGeographicAreaRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of GeographicArea (e.g. [][]GeographicArea), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseGeographicAreaRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aGeographicArea, ok := obj.(GeographicArea)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertGeographicAreaRequired(aGeographicArea)
	})
}
