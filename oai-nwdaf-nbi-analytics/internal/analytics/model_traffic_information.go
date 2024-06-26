/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

// TrafficInformation - Traffic information including UL/DL data rate and/or Traffic volume.
type TrafficInformation struct {

	// String representing a bit rate prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	UplinkRate string `json:"uplinkRate,omitempty"`

	// String representing a bit rate prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	DownlinkRate string `json:"downlinkRate,omitempty"`

	// Unsigned integer identifying a volume in units of bytes.
	UplinkVolume int64 `json:"uplinkVolume,omitempty"`

	// Unsigned integer identifying a volume in units of bytes.
	DownlinkVolume int64 `json:"downlinkVolume,omitempty"`

	// Unsigned integer identifying a volume in units of bytes.
	TotalVolume int64 `json:"totalVolume,omitempty"`
}

// AssertTrafficInformationRequired checks if the required fields are not zero-ed
func AssertTrafficInformationRequired(obj TrafficInformation) error {
	return nil
}

// AssertRecurseTrafficInformationRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TrafficInformation (e.g. [][]TrafficInformation), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTrafficInformationRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTrafficInformation, ok := obj.(TrafficInformation)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTrafficInformationRequired(aTrafficInformation)
	})
}
