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

// TacRange Range of TACs (Tracking Area Codes)
type TacRange struct {
	Start *string `json:"start,omitempty"`
	End *string `json:"end,omitempty"`
	Pattern *string `json:"pattern,omitempty"`
}

// NewTacRange instantiates a new TacRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTacRange() *TacRange {
	this := TacRange{}
	return &this
}

// NewTacRangeWithDefaults instantiates a new TacRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTacRangeWithDefaults() *TacRange {
	this := TacRange{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *TacRange) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TacRange) GetStartOk() (*string, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *TacRange) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *TacRange) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *TacRange) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TacRange) GetEndOk() (*string, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *TacRange) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *TacRange) SetEnd(v string) {
	o.End = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *TacRange) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TacRange) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *TacRange) HasPattern() bool {
	if o != nil && o.Pattern != nil {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *TacRange) SetPattern(v string) {
	o.Pattern = &v
}

func (o TacRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	return json.Marshal(toSerialize)
}

type NullableTacRange struct {
	value *TacRange
	isSet bool
}

func (v NullableTacRange) Get() *TacRange {
	return v.value
}

func (v *NullableTacRange) Set(val *TacRange) {
	v.value = val
	v.isSet = true
}

func (v NullableTacRange) IsSet() bool {
	return v.isSet
}

func (v *NullableTacRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTacRange(val *TacRange) *NullableTacRange {
	return &NullableTacRange{value: val, isSet: true}
}

func (v NullableTacRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTacRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


