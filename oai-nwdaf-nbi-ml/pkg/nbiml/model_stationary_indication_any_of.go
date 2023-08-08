/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.0.0-alpha.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nbiml

type StationaryIndicationAnyOf string

// List of StationaryIndicationAnyOf
const (
	STATIONARYINDICATIONANYOF_STATIONARY StationaryIndicationAnyOf = "STATIONARY"
	STATIONARYINDICATIONANYOF_MOBILE     StationaryIndicationAnyOf = "MOBILE"
)

// AssertStationaryIndicationAnyOfRequired checks if the required fields are not zero-ed
func AssertStationaryIndicationAnyOfRequired(obj StationaryIndicationAnyOf) error {
	return nil
}

// AssertRecurseStationaryIndicationAnyOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of StationaryIndicationAnyOf (e.g. [][]StationaryIndicationAnyOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseStationaryIndicationAnyOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aStationaryIndicationAnyOf, ok := obj.(StationaryIndicationAnyOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertStationaryIndicationAnyOfRequired(aStationaryIndicationAnyOf)
	})
}
