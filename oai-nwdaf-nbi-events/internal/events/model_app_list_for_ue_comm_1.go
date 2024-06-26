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

// AppListForUeComm1 - Represents the analytics of the application list used by UE.
type AppListForUeComm1 struct {

	// String providing an application identifier.
	AppId string `json:"appId,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	StartTime time.Time `json:"startTime,omitempty"`

	// indicating a time in seconds.
	AppDur int32 `json:"appDur,omitempty"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	OccurRatio int32 `json:"occurRatio,omitempty"`

	SpatialValidity NetworkAreaInfo `json:"spatialValidity,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty"`
}

// AssertAppListForUeComm1Required checks if the required fields are not zero-ed
func AssertAppListForUeComm1Required(obj AppListForUeComm1) error {
	if err := AssertNetworkAreaInfoRequired(obj.SpatialValidity); err != nil {
		return err
	}
	return nil
}

// AssertRecurseAppListForUeComm1Required recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AppListForUeComm1 (e.g. [][]AppListForUeComm1), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAppListForUeComm1Required(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAppListForUeComm1, ok := obj.(AppListForUeComm1)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAppListForUeComm1Required(aAppListForUeComm1)
	})
}
