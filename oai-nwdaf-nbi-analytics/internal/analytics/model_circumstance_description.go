/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

import (
	"time"
)

// CircumstanceDescription - Contains the description of a circumstance.
type CircumstanceDescription struct {

	// string with format 'float' as defined in OpenAPI.
	Freq float32 `json:"freq,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	Tm time.Time `json:"tm,omitempty"`

	LocArea NetworkAreaInfo `json:"locArea,omitempty"`

	// Unsigned integer identifying a volume in units of bytes.
	Vol int64 `json:"vol,omitempty"`
}

// AssertCircumstanceDescriptionRequired checks if the required fields are not zero-ed
func AssertCircumstanceDescriptionRequired(obj CircumstanceDescription) error {
	if err := AssertNetworkAreaInfoRequired(obj.LocArea); err != nil {
		return err
	}
	return nil
}

// AssertRecurseCircumstanceDescriptionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CircumstanceDescription (e.g. [][]CircumstanceDescription), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCircumstanceDescriptionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCircumstanceDescription, ok := obj.(CircumstanceDescription)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCircumstanceDescriptionRequired(aCircumstanceDescription)
	})
}