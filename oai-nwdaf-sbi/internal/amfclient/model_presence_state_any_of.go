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

// PresenceStateAnyOf the model 'PresenceStateAnyOf'
type PresenceStateAnyOf string

// List of PresenceState_anyOf
const (
	PRESENCESTATEANYOF_IN_AREA PresenceStateAnyOf = "IN_AREA"
	PRESENCESTATEANYOF_OUT_OF_AREA PresenceStateAnyOf = "OUT_OF_AREA"
	PRESENCESTATEANYOF_UNKNOWN PresenceStateAnyOf = "UNKNOWN"
	PRESENCESTATEANYOF_INACTIVE PresenceStateAnyOf = "INACTIVE"
)

// All allowed values of PresenceStateAnyOf enum
var AllowedPresenceStateAnyOfEnumValues = []PresenceStateAnyOf{
	"IN_AREA",
	"OUT_OF_AREA",
	"UNKNOWN",
	"INACTIVE",
}

func (v *PresenceStateAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PresenceStateAnyOf(value)
	for _, existing := range AllowedPresenceStateAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PresenceStateAnyOf", value)
}

// NewPresenceStateAnyOfFromValue returns a pointer to a valid PresenceStateAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPresenceStateAnyOfFromValue(v string) (*PresenceStateAnyOf, error) {
	ev := PresenceStateAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PresenceStateAnyOf: valid values are %v", v, AllowedPresenceStateAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PresenceStateAnyOf) IsValid() bool {
	for _, existing := range AllowedPresenceStateAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PresenceState_anyOf value
func (v PresenceStateAnyOf) Ptr() *PresenceStateAnyOf {
	return &v
}

type NullablePresenceStateAnyOf struct {
	value *PresenceStateAnyOf
	isSet bool
}

func (v NullablePresenceStateAnyOf) Get() *PresenceStateAnyOf {
	return v.value
}

func (v *NullablePresenceStateAnyOf) Set(val *PresenceStateAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePresenceStateAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePresenceStateAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePresenceStateAnyOf(val *PresenceStateAnyOf) *NullablePresenceStateAnyOf {
	return &NullablePresenceStateAnyOf{value: val, isSet: true}
}

func (v NullablePresenceStateAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePresenceStateAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

