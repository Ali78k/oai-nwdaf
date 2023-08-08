/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

type PointAltitudeAllOf struct {
	Point GeographicalCoordinates `json:"point"`

	// Indicates value of altitude.
	Altitude float64 `json:"altitude"`
}

// AssertPointAltitudeAllOfRequired checks if the required fields are not zero-ed
func AssertPointAltitudeAllOfRequired(obj PointAltitudeAllOf) error {
	elements := map[string]interface{}{
		"point":    obj.Point,
		"altitude": obj.Altitude,
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

// AssertRecursePointAltitudeAllOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PointAltitudeAllOf (e.g. [][]PointAltitudeAllOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePointAltitudeAllOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPointAltitudeAllOf, ok := obj.(PointAltitudeAllOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPointAltitudeAllOfRequired(aPointAltitudeAllOf)
	})
}
