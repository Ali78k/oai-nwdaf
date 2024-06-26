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

// ReachabilityFilterAnyOf the model 'ReachabilityFilterAnyOf'
type ReachabilityFilterAnyOf string

// List of ReachabilityFilter_anyOf
const (
	REACHABILITYFILTERANYOF_REACHABILITY_STATUS_CHANGE ReachabilityFilterAnyOf = "UE_REACHABILITY_STATUS_CHANGE"
	REACHABILITYFILTERANYOF_REACHABLE_DL_TRAFFIC ReachabilityFilterAnyOf = "UE_REACHABLE_DL_TRAFFIC"
)

// All allowed values of ReachabilityFilterAnyOf enum
var AllowedReachabilityFilterAnyOfEnumValues = []ReachabilityFilterAnyOf{
	"UE_REACHABILITY_STATUS_CHANGE",
	"UE_REACHABLE_DL_TRAFFIC",
}

func (v *ReachabilityFilterAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReachabilityFilterAnyOf(value)
	for _, existing := range AllowedReachabilityFilterAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReachabilityFilterAnyOf", value)
}

// NewReachabilityFilterAnyOfFromValue returns a pointer to a valid ReachabilityFilterAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReachabilityFilterAnyOfFromValue(v string) (*ReachabilityFilterAnyOf, error) {
	ev := ReachabilityFilterAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReachabilityFilterAnyOf: valid values are %v", v, AllowedReachabilityFilterAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReachabilityFilterAnyOf) IsValid() bool {
	for _, existing := range AllowedReachabilityFilterAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReachabilityFilter_anyOf value
func (v ReachabilityFilterAnyOf) Ptr() *ReachabilityFilterAnyOf {
	return &v
}

type NullableReachabilityFilterAnyOf struct {
	value *ReachabilityFilterAnyOf
	isSet bool
}

func (v NullableReachabilityFilterAnyOf) Get() *ReachabilityFilterAnyOf {
	return v.value
}

func (v *NullableReachabilityFilterAnyOf) Set(val *ReachabilityFilterAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReachabilityFilterAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReachabilityFilterAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReachabilityFilterAnyOf(val *ReachabilityFilterAnyOf) *NullableReachabilityFilterAnyOf {
	return &NullableReachabilityFilterAnyOf{value: val, isSet: true}
}

func (v NullableReachabilityFilterAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReachabilityFilterAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

