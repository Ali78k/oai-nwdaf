/*
Namf_EventExposure

AMF Event Exposure Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.2.0-alpha.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package amfclient

import (
	"encoding/json"
	"errors"
	"reflect"
)

// AmfEventState Represents the state of a subscribed event
type AmfEventState struct {
	Active bool // `json:"active"` // use custom unmarshalling
	RemainReports *int32 `json:"remainReports,omitempty"`
	// indicating a time in seconds.
	RemainDuration *int32 `json:"remainDuration,omitempty"`
}

// NewAmfEventState instantiates a new AmfEventState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfEventState(active bool) *AmfEventState {
	this := AmfEventState{}
	this.Active = active
	return &this
}

// NewAmfEventStateWithDefaults instantiates a new AmfEventState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfEventStateWithDefaults() *AmfEventState {
	this := AmfEventState{}
	return &this
}

// GetActive returns the Active field value
func (o *AmfEventState) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *AmfEventState) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *AmfEventState) SetActive(v bool) {
	o.Active = v
}

// GetRemainReports returns the RemainReports field value if set, zero value otherwise.
func (o *AmfEventState) GetRemainReports() int32 {
	if o == nil || o.RemainReports == nil {
		var ret int32
		return ret
	}
	return *o.RemainReports
}

// GetRemainReportsOk returns a tuple with the RemainReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventState) GetRemainReportsOk() (*int32, bool) {
	if o == nil || o.RemainReports == nil {
		return nil, false
	}
	return o.RemainReports, true
}

// HasRemainReports returns a boolean if a field has been set.
func (o *AmfEventState) HasRemainReports() bool {
	if o != nil && o.RemainReports != nil {
		return true
	}

	return false
}

// SetRemainReports gets a reference to the given int32 and assigns it to the RemainReports field.
func (o *AmfEventState) SetRemainReports(v int32) {
	o.RemainReports = &v
}

// GetRemainDuration returns the RemainDuration field value if set, zero value otherwise.
func (o *AmfEventState) GetRemainDuration() int32 {
	if o == nil || o.RemainDuration == nil {
		var ret int32
		return ret
	}
	return *o.RemainDuration
}

// GetRemainDurationOk returns a tuple with the RemainDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventState) GetRemainDurationOk() (*int32, bool) {
	if o == nil || o.RemainDuration == nil {
		return nil, false
	}
	return o.RemainDuration, true
}

// HasRemainDuration returns a boolean if a field has been set.
func (o *AmfEventState) HasRemainDuration() bool {
	if o != nil && o.RemainDuration != nil {
		return true
	}

	return false
}

// SetRemainDuration gets a reference to the given int32 and assigns it to the RemainDuration field.
func (o *AmfEventState) SetRemainDuration(v int32) {
	o.RemainDuration = &v
}

func (o AmfEventState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["active"] = o.Active
	}
	if o.RemainReports != nil {
		toSerialize["remainReports"] = o.RemainReports
	}
	if o.RemainDuration != nil {
		toSerialize["remainDuration"] = o.RemainDuration
	}
	return json.Marshal(toSerialize)
}

type NullableAmfEventState struct {
	value *AmfEventState
	isSet bool
}

func (v NullableAmfEventState) Get() *AmfEventState {
	return v.value
}

func (v *NullableAmfEventState) Set(val *AmfEventState) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfEventState) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfEventState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfEventState(val *AmfEventState) *NullableAmfEventState {
	return &NullableAmfEventState{value: val, isSet: true}
}

func (v NullableAmfEventState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfEventState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (v *AmfEventState) UnmarshalJSON(src []byte) error {
	// custom unmarshaling for Active: string or bool (as used in AMF) => bool
	type AmfEventStateAlias AmfEventState
	alias := &struct {
		Active interface{} `json:"active"`
		*AmfEventStateAlias
	}{
		AmfEventStateAlias: (*AmfEventStateAlias)(v),
	}
    err := json.Unmarshal(src, alias)
    if err != nil { return err; }
	activeVal := reflect.ValueOf(alias.Active)
	switch activeVal.Kind() {
	case reflect.String:
		activeString := activeVal.String()
		if activeString == "TRUE" || activeString == "True" || activeString == "true" {
			v.Active = bool(true)
		} else {
			if activeString == "FALSE" || activeString == "False" || activeString == "false" {
				v.Active = bool(false)
			} else {
				return errors.New("undefined string representation of AmfEventState.Active: NOT true NOR false (case insensitive)")
			}
		}
	case reflect.Bool:
		v.Active = activeVal.Bool()
	default:
		return errors.New("undefined type of AmfEventState.Active: not one of (string, bool)")
	}
	return nil
}
