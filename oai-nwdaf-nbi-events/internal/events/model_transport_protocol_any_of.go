/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

type TransportProtocolAnyOf string

// List of TransportProtocolAnyOf
const (
	TRANSPORTPROTOCOLANYOF_UDP TransportProtocolAnyOf = "UDP"
	TRANSPORTPROTOCOLANYOF_TCP TransportProtocolAnyOf = "TCP"
)

// AssertTransportProtocolAnyOfRequired checks if the required fields are not zero-ed
func AssertTransportProtocolAnyOfRequired(obj TransportProtocolAnyOf) error {
	return nil
}

// AssertRecurseTransportProtocolAnyOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TransportProtocolAnyOf (e.g. [][]TransportProtocolAnyOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTransportProtocolAnyOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTransportProtocolAnyOf, ok := obj.(TransportProtocolAnyOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTransportProtocolAnyOfRequired(aTransportProtocolAnyOf)
	})
}
