/*
Namf_EventExposure

AMF Event Exposure Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.2.0-alpha.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package amfclient

import (
	"encoding/json"
	"fmt"
)

// CmStateAnyOf the model 'CmStateAnyOf'
type CmStateAnyOf string

// List of CmState_anyOf
const (
	CMSTATEANYOF_IDLE CmStateAnyOf = "IDLE"
	CMSTATEANYOF_CONNECTED CmStateAnyOf = "CONNECTED"
)

// All allowed values of CmStateAnyOf enum
var AllowedCmStateAnyOfEnumValues = []CmStateAnyOf{
	"IDLE",
	"CONNECTED",
}

func (v *CmStateAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CmStateAnyOf(value)
	for _, existing := range AllowedCmStateAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CmStateAnyOf", value)
}

// NewCmStateAnyOfFromValue returns a pointer to a valid CmStateAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCmStateAnyOfFromValue(v string) (*CmStateAnyOf, error) {
	ev := CmStateAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CmStateAnyOf: valid values are %v", v, AllowedCmStateAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CmStateAnyOf) IsValid() bool {
	for _, existing := range AllowedCmStateAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CmState_anyOf value
func (v CmStateAnyOf) Ptr() *CmStateAnyOf {
	return &v
}

type NullableCmStateAnyOf struct {
	value *CmStateAnyOf
	isSet bool
}

func (v NullableCmStateAnyOf) Get() *CmStateAnyOf {
	return v.value
}

func (v *NullableCmStateAnyOf) Set(val *CmStateAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCmStateAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCmStateAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmStateAnyOf(val *CmStateAnyOf) *NullableCmStateAnyOf {
	return &NullableCmStateAnyOf{value: val, isSet: true}
}

func (v NullableCmStateAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmStateAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

