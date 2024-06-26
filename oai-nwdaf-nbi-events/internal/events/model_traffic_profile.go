/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// TrafficProfile - Possible values are: - SINGLE_TRANS_UL: Uplink single packet transmission. - SINGLE_TRANS_DL: Downlink single packet transmission. - DUAL_TRANS_UL_FIRST: Dual packet transmission, firstly uplink packet transmission with subsequent downlink packet transmission. - DUAL_TRANS_DL_FIRST: Dual packet transmission, firstly downlink packet transmission with subsequent uplink packet transmission.
type TrafficProfile struct {
}

// AssertTrafficProfileRequired checks if the required fields are not zero-ed
func AssertTrafficProfileRequired(obj TrafficProfile) error {
	return nil
}

// AssertRecurseTrafficProfileRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TrafficProfile (e.g. [][]TrafficProfile), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTrafficProfileRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTrafficProfile, ok := obj.(TrafficProfile)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTrafficProfileRequired(aTrafficProfile)
	})
}
