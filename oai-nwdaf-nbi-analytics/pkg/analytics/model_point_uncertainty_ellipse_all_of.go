/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

type PointUncertaintyEllipseAllOf struct {
	Point GeographicalCoordinates `json:"point"`

	UncertaintyEllipse UncertaintyEllipse `json:"uncertaintyEllipse"`

	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}

// AssertPointUncertaintyEllipseAllOfRequired checks if the required fields are not zero-ed
func AssertPointUncertaintyEllipseAllOfRequired(obj PointUncertaintyEllipseAllOf) error {
	elements := map[string]interface{}{
		"point":              obj.Point,
		"uncertaintyEllipse": obj.UncertaintyEllipse,
		"confidence":         obj.Confidence,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertGeographicalCoordinatesRequired(obj.Point); err != nil {
		return err
	}
	if err := AssertUncertaintyEllipseRequired(obj.UncertaintyEllipse); err != nil {
		return err
	}
	return nil
}

// AssertRecursePointUncertaintyEllipseAllOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PointUncertaintyEllipseAllOf (e.g. [][]PointUncertaintyEllipseAllOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePointUncertaintyEllipseAllOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPointUncertaintyEllipseAllOf, ok := obj.(PointUncertaintyEllipseAllOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPointUncertaintyEllipseAllOfRequired(aPointUncertaintyEllipseAllOf)
	})
}
