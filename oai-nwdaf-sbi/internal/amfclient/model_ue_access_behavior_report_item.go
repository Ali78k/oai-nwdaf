/*
Namf_EventExposure

AMF Event Exposure Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.2.0-alpha.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package amfclient

import (
	"encoding/json"
)

// UeAccessBehaviorReportItem Report Item for UE Access Behavior Trends event.
type UeAccessBehaviorReportItem struct {
	StateTransitionType AccessStateTransitionType `json:"stateTransitionType"`
	// indicating a time in seconds.
	Spacing int32 `json:"spacing"`
	// indicating a time in seconds.
	Duration int32 `json:"duration"`
}

// NewUeAccessBehaviorReportItem instantiates a new UeAccessBehaviorReportItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeAccessBehaviorReportItem(stateTransitionType AccessStateTransitionType, spacing int32, duration int32) *UeAccessBehaviorReportItem {
	this := UeAccessBehaviorReportItem{}
	this.StateTransitionType = stateTransitionType
	this.Spacing = spacing
	this.Duration = duration
	return &this
}

// NewUeAccessBehaviorReportItemWithDefaults instantiates a new UeAccessBehaviorReportItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeAccessBehaviorReportItemWithDefaults() *UeAccessBehaviorReportItem {
	this := UeAccessBehaviorReportItem{}
	return &this
}

// GetStateTransitionType returns the StateTransitionType field value
func (o *UeAccessBehaviorReportItem) GetStateTransitionType() AccessStateTransitionType {
	if o == nil {
		var ret AccessStateTransitionType
		return ret
	}

	return o.StateTransitionType
}

// GetStateTransitionTypeOk returns a tuple with the StateTransitionType field value
// and a boolean to check if the value has been set.
func (o *UeAccessBehaviorReportItem) GetStateTransitionTypeOk() (*AccessStateTransitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateTransitionType, true
}

// SetStateTransitionType sets field value
func (o *UeAccessBehaviorReportItem) SetStateTransitionType(v AccessStateTransitionType) {
	o.StateTransitionType = v
}

// GetSpacing returns the Spacing field value
func (o *UeAccessBehaviorReportItem) GetSpacing() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Spacing
}

// GetSpacingOk returns a tuple with the Spacing field value
// and a boolean to check if the value has been set.
func (o *UeAccessBehaviorReportItem) GetSpacingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spacing, true
}

// SetSpacing sets field value
func (o *UeAccessBehaviorReportItem) SetSpacing(v int32) {
	o.Spacing = v
}

// GetDuration returns the Duration field value
func (o *UeAccessBehaviorReportItem) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *UeAccessBehaviorReportItem) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *UeAccessBehaviorReportItem) SetDuration(v int32) {
	o.Duration = v
}

func (o UeAccessBehaviorReportItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["stateTransitionType"] = o.StateTransitionType
	}
	if true {
		toSerialize["spacing"] = o.Spacing
	}
	if true {
		toSerialize["duration"] = o.Duration
	}
	return json.Marshal(toSerialize)
}

type NullableUeAccessBehaviorReportItem struct {
	value *UeAccessBehaviorReportItem
	isSet bool
}

func (v NullableUeAccessBehaviorReportItem) Get() *UeAccessBehaviorReportItem {
	return v.value
}

func (v *NullableUeAccessBehaviorReportItem) Set(val *UeAccessBehaviorReportItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUeAccessBehaviorReportItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUeAccessBehaviorReportItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeAccessBehaviorReportItem(val *UeAccessBehaviorReportItem) *NullableUeAccessBehaviorReportItem {
	return &NullableUeAccessBehaviorReportItem{value: val, isSet: true}
}

func (v NullableUeAccessBehaviorReportItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeAccessBehaviorReportItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


