/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

// AddrFqdn - IP address and/or FQDN.
type AddrFqdn struct {
	IpAddr IpAddr `json:"ipAddr,omitempty"`

	// Indicates an FQDN.
	Fqdn string `json:"fqdn,omitempty"`
}

// AssertAddrFqdnRequired checks if the required fields are not zero-ed
func AssertAddrFqdnRequired(obj AddrFqdn) error {
	if err := AssertIpAddrRequired(obj.IpAddr); err != nil {
		return err
	}
	return nil
}

// AssertRecurseAddrFqdnRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AddrFqdn (e.g. [][]AddrFqdn), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAddrFqdnRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAddrFqdn, ok := obj.(AddrFqdn)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAddrFqdnRequired(aAddrFqdn)
	})
}
