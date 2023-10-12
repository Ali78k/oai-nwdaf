/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// StationaryIndication - Possible values are: - STATIONARY: Identifies the UE is stationary - MOBILE: Identifies the UE is mobile
type StationaryIndication struct {
}

// AssertStationaryIndicationRequired checks if the required fields are not zero-ed
func AssertStationaryIndicationRequired(obj StationaryIndication) error {
	return nil
}

// AssertRecurseStationaryIndicationRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of StationaryIndication (e.g. [][]StationaryIndication), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseStationaryIndicationRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aStationaryIndication, ok := obj.(StationaryIndication)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertStationaryIndicationRequired(aStationaryIndication)
	})
}