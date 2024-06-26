/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// DispersionCollection - Dispersion collection per UE location or per slice.
type DispersionCollection struct {
	UeLoc UserLocation `json:"ueLoc,omitempty"`

	Snssai Snssai `json:"snssai,omitempty"`

	Supis []string `json:"supis,omitempty"`

	Gpsis []string `json:"gpsis,omitempty"`

	AppVolumes []ApplicationVolume `json:"appVolumes,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	DisperAmount int32 `json:"disperAmount,omitempty"`

	DisperClass DispersionClass `json:"disperClass,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	UsageRank int32 `json:"usageRank,omitempty"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	PercentileRank int32 `json:"percentileRank,omitempty"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	UeRatio int32 `json:"ueRatio,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty"`
}

// AssertDispersionCollectionRequired checks if the required fields are not zero-ed
func AssertDispersionCollectionRequired(obj DispersionCollection) error {
	if err := AssertUserLocationRequired(obj.UeLoc); err != nil {
		return err
	}
	if err := AssertSnssaiRequired(obj.Snssai); err != nil {
		return err
	}
	for _, el := range obj.AppVolumes {
		if err := AssertApplicationVolumeRequired(el); err != nil {
			return err
		}
	}
	if err := AssertDispersionClassRequired(obj.DisperClass); err != nil {
		return err
	}
	return nil
}

// AssertRecurseDispersionCollectionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of DispersionCollection (e.g. [][]DispersionCollection), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseDispersionCollectionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aDispersionCollection, ok := obj.(DispersionCollection)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertDispersionCollectionRequired(aDispersionCollection)
	})
}
