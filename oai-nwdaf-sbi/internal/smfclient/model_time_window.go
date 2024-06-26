/*
Nsmf_EventExposure

Session Management Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smfclient

import (
	"encoding/json"
	"time"
)

// TimeWindow Represents a time window identified by a start time and a stop time.
type TimeWindow struct {
	// string with format \"date-time\" as defined in OpenAPI.
	StartTime time.Time `json:"startTime"`
	// string with format \"date-time\" as defined in OpenAPI.
	StopTime time.Time `json:"stopTime"`
}

// NewTimeWindow instantiates a new TimeWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeWindow(startTime time.Time, stopTime time.Time) *TimeWindow {
	this := TimeWindow{}
	this.StartTime = startTime
	this.StopTime = stopTime
	return &this
}

// NewTimeWindowWithDefaults instantiates a new TimeWindow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeWindowWithDefaults() *TimeWindow {
	this := TimeWindow{}
	return &this
}

// GetStartTime returns the StartTime field value
func (o *TimeWindow) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *TimeWindow) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *TimeWindow) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetStopTime returns the StopTime field value
func (o *TimeWindow) GetStopTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StopTime
}

// GetStopTimeOk returns a tuple with the StopTime field value
// and a boolean to check if the value has been set.
func (o *TimeWindow) GetStopTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StopTime, true
}

// SetStopTime sets field value
func (o *TimeWindow) SetStopTime(v time.Time) {
	o.StopTime = v
}

func (o TimeWindow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["startTime"] = o.StartTime
	}
	if true {
		toSerialize["stopTime"] = o.StopTime
	}
	return json.Marshal(toSerialize)
}

type NullableTimeWindow struct {
	value *TimeWindow
	isSet bool
}

func (v NullableTimeWindow) Get() *TimeWindow {
	return v.value
}

func (v *NullableTimeWindow) Set(val *TimeWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeWindow(val *TimeWindow) *NullableTimeWindow {
	return &NullableTimeWindow{value: val, isSet: true}
}

func (v NullableTimeWindow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


