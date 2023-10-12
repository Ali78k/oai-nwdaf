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

// EventNotification1 - Represents a notification on events that occurred.
type EventNotification1 struct {
	Event NwdafEvent `json:"event"`

	// string with format 'date-time' as defined in OpenAPI.
	Start time.Time `json:"start,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	Expiry time.Time `json:"expiry,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	TimeStampGen time.Time `json:"timeStampGen,omitempty"`

	FailNotifyCode NwdafFailureCode `json:"failNotifyCode,omitempty"`

	// indicating a time in seconds.
	RvWaitTime int32 `json:"rvWaitTime,omitempty"`

	AnaMetaInfo AnalyticsMetadataInfo1 `json:"anaMetaInfo,omitempty"`

	NfLoadLevelInfos []NfLoadLevelInformation1 `json:"nfLoadLevelInfos,omitempty"`

	NsiLoadLevelInfos []NsiLoadLevelInfo1 `json:"nsiLoadLevelInfos,omitempty"`

	SliceLoadLevelInfo SliceLoadLevelInformation1 `json:"sliceLoadLevelInfo,omitempty"`

	SvcExps []ServiceExperienceInfo1 `json:"svcExps,omitempty"`

	QosSustainInfos []QosSustainabilityInfo1 `json:"qosSustainInfos,omitempty"`

	UeComms []UeCommunication1 `json:"ueComms,omitempty"`

	UeMobs []UeMobility1 `json:"ueMobs,omitempty"`

	UserDataCongInfos []UserDataCongestionInfo1 `json:"userDataCongInfos,omitempty"`

	AbnorBehavrs []AbnormalBehaviour1 `json:"abnorBehavrs,omitempty"`

	NwPerfs []NetworkPerfInfo1 `json:"nwPerfs,omitempty"`

	DnPerfInfos []DnPerfInfo1 `json:"dnPerfInfos,omitempty"`

	DisperInfos []DispersionInfo1 `json:"disperInfos,omitempty"`

	RedTransInfos []RedundantTransmissionExpInfo1 `json:"redTransInfos,omitempty"`

	WlanInfos []WlanPerformanceInfo1 `json:"wlanInfos,omitempty"`
}

// AssertEventNotification1Required checks if the required fields are not zero-ed
func AssertEventNotification1Required(obj EventNotification1) error {
	elements := map[string]interface{}{
		"event": obj.Event,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertNwdafEventRequired(obj.Event); err != nil {
		return err
	}
	if err := AssertNwdafFailureCodeRequired(obj.FailNotifyCode); err != nil {
		return err
	}
	if err := AssertAnalyticsMetadataInfo1Required(obj.AnaMetaInfo); err != nil {
		return err
	}
	for _, el := range obj.NfLoadLevelInfos {
		if err := AssertNfLoadLevelInformation1Required(el); err != nil {
			return err
		}
	}
	for _, el := range obj.NsiLoadLevelInfos {
		if err := AssertNsiLoadLevelInfo1Required(el); err != nil {
			return err
		}
	}
	if err := AssertSliceLoadLevelInformation1Required(obj.SliceLoadLevelInfo); err != nil {
		return err
	}
	for _, el := range obj.SvcExps {
		if err := AssertServiceExperienceInfo1Required(el); err != nil {
			return err
		}
	}
	for _, el := range obj.QosSustainInfos {
		if err := AssertQosSustainabilityInfo1Required(el); err != nil {
			return err
		}
	}
	for _, el := range obj.UeComms {
		if err := AssertUeCommunication1Required(el); err != nil {
			return err
		}
	}
	for _, el := range obj.UeMobs {
		if err := AssertUeMobility1Required(el); err != nil {
			return err
		}
	}
	for _, el := range obj.UserDataCongInfos {
		if err := AssertUserDataCongestionInfo1Required(el); err != nil {
			return err
		}
	}
	for _, el := range obj.AbnorBehavrs {
		if err := AssertAbnormalBehaviour1Required(el); err != nil {
			return err
		}
	}
	for _, el := range obj.NwPerfs {
		if err := AssertNetworkPerfInfo1Required(el); err != nil {
			return err
		}
	}
	for _, el := range obj.DnPerfInfos {
		if err := AssertDnPerfInfo1Required(el); err != nil {
			return err
		}
	}
	for _, el := range obj.DisperInfos {
		if err := AssertDispersionInfo1Required(el); err != nil {
			return err
		}
	}
	for _, el := range obj.RedTransInfos {
		if err := AssertRedundantTransmissionExpInfo1Required(el); err != nil {
			return err
		}
	}
	for _, el := range obj.WlanInfos {
		if err := AssertWlanPerformanceInfo1Required(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseEventNotification1Required recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of EventNotification1 (e.g. [][]EventNotification1), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseEventNotification1Required(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aEventNotification1, ok := obj.(EventNotification1)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertEventNotification1Required(aEventNotification1)
	})
}