/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

// PointAltitudeUncertainty - Ellipsoid point with altitude and uncertainty ellipsoid.
type PointAltitudeUncertainty struct {
	GadShape

	Point GeographicalCoordinates `json:"point"`

	// Indicates value of altitude.
	Altitude float64 `json:"altitude"`

	UncertaintyEllipse UncertaintyEllipse `json:"uncertaintyEllipse"`

	// Indicates value of uncertainty.
	UncertaintyAltitude float32 `json:"uncertaintyAltitude"`

	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}

// AssertPointAltitudeUncertaintyRequired checks if the required fields are not zero-ed
func AssertPointAltitudeUncertaintyRequired(obj PointAltitudeUncertainty) error {
	elements := map[string]interface{}{
		"point":               obj.Point,
		"altitude":            obj.Altitude,
		"uncertaintyEllipse":  obj.UncertaintyEllipse,
		"uncertaintyAltitude": obj.UncertaintyAltitude,
		"confidence":          obj.Confidence,
		"shape":               obj.Shape,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertGadShapeRequired(obj.GadShape); err != nil {
		return err
	}

	if err := AssertGeographicalCoordinatesRequired(obj.Point); err != nil {
		return err
	}
	if err := AssertUncertaintyEllipseRequired(obj.UncertaintyEllipse); err != nil {
		return err
	}
	return nil
}

// AssertRecursePointAltitudeUncertaintyRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PointAltitudeUncertainty (e.g. [][]PointAltitudeUncertainty), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePointAltitudeUncertaintyRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPointAltitudeUncertainty, ok := obj.(PointAltitudeUncertainty)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPointAltitudeUncertaintyRequired(aPointAltitudeUncertainty)
	})
}
