/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// IpAddr - Contains an IP adresse.
type IpAddr struct {

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	Ipv4Addr string `json:"ipv4Addr,omitempty"`

	Ipv6Addr Ipv6Addr `json:"ipv6Addr,omitempty"`

	Ipv6Prefix Ipv6Prefix `json:"ipv6Prefix,omitempty"`
}

// AssertIpAddrRequired checks if the required fields are not zero-ed
func AssertIpAddrRequired(obj IpAddr) error {
	if err := AssertIpv6AddrRequired(obj.Ipv6Addr); err != nil {
		return err
	}
	if err := AssertIpv6PrefixRequired(obj.Ipv6Prefix); err != nil {
		return err
	}
	return nil
}

// AssertRecurseIpAddrRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of IpAddr (e.g. [][]IpAddr), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseIpAddrRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aIpAddr, ok := obj.(IpAddr)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertIpAddrRequired(aIpAddr)
	})
}
