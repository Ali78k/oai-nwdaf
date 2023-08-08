/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

import (
	"time"
)

// QosSustainabilityInfo - Represents the QoS Sustainability information.
type QosSustainabilityInfo struct {
	AreaInfo NetworkAreaInfo `json:"areaInfo,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	StartTs time.Time `json:"startTs,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	EndTs time.Time `json:"endTs,omitempty"`

	QosFlowRetThd RetainabilityThreshold `json:"qosFlowRetThd,omitempty"`

	// String representing a bit rate prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	RanUeThrouThd string `json:"ranUeThrouThd,omitempty"`

	Snssai Snssai `json:"snssai,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty"`
}

// AssertQosSustainabilityInfoRequired checks if the required fields are not zero-ed
func AssertQosSustainabilityInfoRequired(obj QosSustainabilityInfo) error {
	if err := AssertNetworkAreaInfoRequired(obj.AreaInfo); err != nil {
		return err
	}
	if err := AssertRetainabilityThresholdRequired(obj.QosFlowRetThd); err != nil {
		return err
	}
	if err := AssertSnssaiRequired(obj.Snssai); err != nil {
		return err
	}
	return nil
}

// AssertRecurseQosSustainabilityInfoRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of QosSustainabilityInfo (e.g. [][]QosSustainabilityInfo), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseQosSustainabilityInfoRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aQosSustainabilityInfo, ok := obj.(QosSustainabilityInfo)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertQosSustainabilityInfoRequired(aQosSustainabilityInfo)
	})
}
