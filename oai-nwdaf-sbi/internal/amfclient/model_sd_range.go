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

// SdRange A range of SDs (Slice Differentiators)
type SdRange struct {
	// First value identifying the start of an SD range. This string shall be formatted as specified for the sd attribute of the Snssai data type in clause 5.4.4.2.
	Start *string `json:"start,omitempty"`
	// Last value identifying the end of an SD range. This string shall be formatted as specified for the sd attribute of the Snssai data type in clause 5.4.4.2.
	End *string `json:"end,omitempty"`
}

// NewSdRange instantiates a new SdRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdRange() *SdRange {
	this := SdRange{}
	return &this
}

// NewSdRangeWithDefaults instantiates a new SdRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdRangeWithDefaults() *SdRange {
	this := SdRange{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *SdRange) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdRange) GetStartOk() (*string, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *SdRange) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *SdRange) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *SdRange) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdRange) GetEndOk() (*string, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *SdRange) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *SdRange) SetEnd(v string) {
	o.End = &v
}

func (o SdRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	return json.Marshal(toSerialize)
}

type NullableSdRange struct {
	value *SdRange
	isSet bool
}

func (v NullableSdRange) Get() *SdRange {
	return v.value
}

func (v *NullableSdRange) Set(val *SdRange) {
	v.value = val
	v.isSet = true
}

func (v NullableSdRange) IsSet() bool {
	return v.isSet
}

func (v *NullableSdRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdRange(val *SdRange) *NullableSdRange {
	return &NullableSdRange{value: val, isSet: true}
}

func (v NullableSdRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


