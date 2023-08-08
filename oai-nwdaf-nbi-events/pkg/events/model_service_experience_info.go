/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// ServiceExperienceInfo - Represents service experience information.
type ServiceExperienceInfo struct {
	SvcExprc SvcExperience `json:"svcExprc"`

	// string with format 'float' as defined in OpenAPI.
	SvcExprcVariance float32 `json:"svcExprcVariance,omitempty"`

	Supis []string `json:"supis,omitempty"`

	Snssai Snssai `json:"snssai,omitempty"`

	// String providing an application identifier.
	AppId string `json:"appId,omitempty"`

	SrvExpcType ServiceExperienceType `json:"srvExpcType,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003; it shall contain either a DNN Network Identifier, or a full DNN with both the Network Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots (e.g. \"Label1.Label2.Label3\").
	Dnn string `json:"dnn,omitempty"`

	NetworkArea NetworkAreaInfo `json:"networkArea,omitempty"`

	// Contains the Identifier of the selected Network Slice instance
	NsiId string `json:"nsiId,omitempty"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	Ratio int32 `json:"ratio,omitempty"`

	RatType RatType `json:"ratType,omitempty"`

	// Integer value indicating the ARFCN applicable for a downlink, uplink or bi-directional (TDD) NR global frequency raster, as definition of \"ARFCN-ValueNR\" IE in clause 6.3.2 of 3GPP TS 38.331.
	Frequency int32 `json:"frequency,omitempty"`
}

// AssertServiceExperienceInfoRequired checks if the required fields are not zero-ed
func AssertServiceExperienceInfoRequired(obj ServiceExperienceInfo) error {
	elements := map[string]interface{}{
		"svcExprc": obj.SvcExprc,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertSvcExperienceRequired(obj.SvcExprc); err != nil {
		return err
	}
	if err := AssertSnssaiRequired(obj.Snssai); err != nil {
		return err
	}
	if err := AssertServiceExperienceTypeRequired(obj.SrvExpcType); err != nil {
		return err
	}
	if err := AssertNetworkAreaInfoRequired(obj.NetworkArea); err != nil {
		return err
	}
	if err := AssertRatTypeRequired(obj.RatType); err != nil {
		return err
	}
	return nil
}

// AssertRecurseServiceExperienceInfoRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ServiceExperienceInfo (e.g. [][]ServiceExperienceInfo), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseServiceExperienceInfoRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aServiceExperienceInfo, ok := obj.(ServiceExperienceInfo)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertServiceExperienceInfoRequired(aServiceExperienceInfo)
	})
}
