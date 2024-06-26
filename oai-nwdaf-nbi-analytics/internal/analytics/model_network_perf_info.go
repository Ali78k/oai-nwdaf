/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

// NetworkPerfInfo - Represents the network performance information.
type NetworkPerfInfo struct {
	NetworkArea NetworkAreaInfo `json:"networkArea"`

	NwPerfType NetworkPerfTypeAnyOf `json:"nwPerfType"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	RelativeRatio *int32 `json:"relativeRatio,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	AbsoluteNum *int32 `json:"absoluteNum,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence *int32 `json:"confidence,omitempty"`
}

// AssertNetworkPerfInfoRequired checks if the required fields are not zero-ed
func AssertNetworkPerfInfoRequired(obj NetworkPerfInfo) error {
	elements := map[string]interface{}{
		"networkArea": obj.NetworkArea,
		"nwPerfType":  obj.NwPerfType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertNetworkAreaInfoRequired(obj.NetworkArea); err != nil {
		return err
	}
	if err := AssertNetworkPerfTypeAnyOfRequired(obj.NwPerfType); err != nil {
		return err
	}
	return nil
}

// AssertRecurseNetworkPerfInfoRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of NetworkPerfInfo (e.g. [][]NetworkPerfInfo), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNetworkPerfInfoRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNetworkPerfInfo, ok := obj.(NetworkPerfInfo)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNetworkPerfInfoRequired(aNetworkPerfInfo)
	})
}
