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

// LocationFilterAnyOf the model 'LocationFilterAnyOf'
type LocationFilterAnyOf string

// List of LocationFilter_anyOf
const (
	LOCATIONFILTERANYOF_TAI LocationFilterAnyOf = "TAI"
	LOCATIONFILTERANYOF_CELL_ID LocationFilterAnyOf = "CELL_ID"
	LOCATIONFILTERANYOF_RAN_NODE LocationFilterAnyOf = "RAN_NODE"
	LOCATIONFILTERANYOF_N3_IWF LocationFilterAnyOf = "N3IWF"
	LOCATIONFILTERANYOF_UE_IP LocationFilterAnyOf = "UE_IP"
	LOCATIONFILTERANYOF_UDP_PORT LocationFilterAnyOf = "UDP_PORT"
	LOCATIONFILTERANYOF_TNAP_ID LocationFilterAnyOf = "TNAP_ID"
	LOCATIONFILTERANYOF_GLI LocationFilterAnyOf = "GLI"
	LOCATIONFILTERANYOF_TWAP_ID LocationFilterAnyOf = "TWAP_ID"
)

// All allowed values of LocationFilterAnyOf enum
var AllowedLocationFilterAnyOfEnumValues = []LocationFilterAnyOf{
	"TAI",
	"CELL_ID",
	"RAN_NODE",
	"N3IWF",
	"UE_IP",
	"UDP_PORT",
	"TNAP_ID",
	"GLI",
	"TWAP_ID",
}

func (v *LocationFilterAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LocationFilterAnyOf(value)
	for _, existing := range AllowedLocationFilterAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LocationFilterAnyOf", value)
}

// NewLocationFilterAnyOfFromValue returns a pointer to a valid LocationFilterAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocationFilterAnyOfFromValue(v string) (*LocationFilterAnyOf, error) {
	ev := LocationFilterAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LocationFilterAnyOf: valid values are %v", v, AllowedLocationFilterAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LocationFilterAnyOf) IsValid() bool {
	for _, existing := range AllowedLocationFilterAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LocationFilter_anyOf value
func (v LocationFilterAnyOf) Ptr() *LocationFilterAnyOf {
	return &v
}

type NullableLocationFilterAnyOf struct {
	value *LocationFilterAnyOf
	isSet bool
}

func (v NullableLocationFilterAnyOf) Get() *LocationFilterAnyOf {
	return v.value
}

func (v *NullableLocationFilterAnyOf) Set(val *LocationFilterAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationFilterAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationFilterAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationFilterAnyOf(val *LocationFilterAnyOf) *NullableLocationFilterAnyOf {
	return &NullableLocationFilterAnyOf{value: val, isSet: true}
}

func (v NullableLocationFilterAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationFilterAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

