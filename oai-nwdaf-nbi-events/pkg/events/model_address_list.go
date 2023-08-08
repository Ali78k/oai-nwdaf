/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// AddressList - Represents a list of IPv4 and/or IPv6 addresses.
type AddressList struct {
	Ipv4Addrs []string `json:"ipv4Addrs,omitempty"`

	Ipv6Addrs []Ipv6Addr `json:"ipv6Addrs,omitempty"`
}

// AssertAddressListRequired checks if the required fields are not zero-ed
func AssertAddressListRequired(obj AddressList) error {
	for _, el := range obj.Ipv6Addrs {
		if err := AssertIpv6AddrRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseAddressListRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AddressList (e.g. [][]AddressList), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAddressListRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAddressList, ok := obj.(AddressList)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAddressListRequired(aAddressList)
	})
}
