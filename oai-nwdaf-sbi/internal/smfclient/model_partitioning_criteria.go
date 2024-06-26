/*
Nsmf_EventExposure

Session Management Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smfclient

import (
	"encoding/json"
	"fmt"
)

// PartitioningCriteria Possible values are: - \"TAC\": Type Allocation Code - \"SUBPLMN\": Subscriber PLMN ID - \"GEOAREA\": Geographical area, i.e. list(s) of TAI(s) - \"SNSSAI\": S-NSSAI - \"DNN\": DNN 
type PartitioningCriteria struct {
	PartitioningCriteriaAnyOf *PartitioningCriteriaAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PartitioningCriteria) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PartitioningCriteriaAnyOf
	err = json.Unmarshal(data, &dst.PartitioningCriteriaAnyOf);
	if err == nil {
		jsonPartitioningCriteriaAnyOf, _ := json.Marshal(dst.PartitioningCriteriaAnyOf)
		if string(jsonPartitioningCriteriaAnyOf) == "{}" { // empty struct
			dst.PartitioningCriteriaAnyOf = nil
		} else {
			return nil // data stored in dst.PartitioningCriteriaAnyOf, return on the first match
		}
	} else {
		dst.PartitioningCriteriaAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(PartitioningCriteria)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PartitioningCriteria) MarshalJSON() ([]byte, error) {
	if src.PartitioningCriteriaAnyOf != nil {
		return json.Marshal(&src.PartitioningCriteriaAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePartitioningCriteria struct {
	value *PartitioningCriteria
	isSet bool
}

func (v NullablePartitioningCriteria) Get() *PartitioningCriteria {
	return v.value
}

func (v *NullablePartitioningCriteria) Set(val *PartitioningCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullablePartitioningCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullablePartitioningCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartitioningCriteria(val *PartitioningCriteria) *NullablePartitioningCriteria {
	return &NullablePartitioningCriteria{value: val, isSet: true}
}

func (v NullablePartitioningCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartitioningCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


