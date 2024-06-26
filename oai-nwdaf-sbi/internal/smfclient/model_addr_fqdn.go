/*
Nsmf_EventExposure

Session Management Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smfclient

import (
	"encoding/json"
)

// AddrFqdn IP address and/or FQDN.
type AddrFqdn struct {
	IpAddr *IpAddr `json:"ipAddr,omitempty"`
	// Indicates an FQDN.
	Fqdn *string `json:"fqdn,omitempty"`
}

// NewAddrFqdn instantiates a new AddrFqdn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddrFqdn() *AddrFqdn {
	this := AddrFqdn{}
	return &this
}

// NewAddrFqdnWithDefaults instantiates a new AddrFqdn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddrFqdnWithDefaults() *AddrFqdn {
	this := AddrFqdn{}
	return &this
}

// GetIpAddr returns the IpAddr field value if set, zero value otherwise.
func (o *AddrFqdn) GetIpAddr() IpAddr {
	if o == nil || o.IpAddr == nil {
		var ret IpAddr
		return ret
	}
	return *o.IpAddr
}

// GetIpAddrOk returns a tuple with the IpAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddrFqdn) GetIpAddrOk() (*IpAddr, bool) {
	if o == nil || o.IpAddr == nil {
		return nil, false
	}
	return o.IpAddr, true
}

// HasIpAddr returns a boolean if a field has been set.
func (o *AddrFqdn) HasIpAddr() bool {
	if o != nil && o.IpAddr != nil {
		return true
	}

	return false
}

// SetIpAddr gets a reference to the given IpAddr and assigns it to the IpAddr field.
func (o *AddrFqdn) SetIpAddr(v IpAddr) {
	o.IpAddr = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *AddrFqdn) GetFqdn() string {
	if o == nil || o.Fqdn == nil {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddrFqdn) GetFqdnOk() (*string, bool) {
	if o == nil || o.Fqdn == nil {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *AddrFqdn) HasFqdn() bool {
	if o != nil && o.Fqdn != nil {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *AddrFqdn) SetFqdn(v string) {
	o.Fqdn = &v
}

func (o AddrFqdn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IpAddr != nil {
		toSerialize["ipAddr"] = o.IpAddr
	}
	if o.Fqdn != nil {
		toSerialize["fqdn"] = o.Fqdn
	}
	return json.Marshal(toSerialize)
}

type NullableAddrFqdn struct {
	value *AddrFqdn
	isSet bool
}

func (v NullableAddrFqdn) Get() *AddrFqdn {
	return v.value
}

func (v *NullableAddrFqdn) Set(val *AddrFqdn) {
	v.value = val
	v.isSet = true
}

func (v NullableAddrFqdn) IsSet() bool {
	return v.isSet
}

func (v *NullableAddrFqdn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddrFqdn(val *AddrFqdn) *NullableAddrFqdn {
	return &NullableAddrFqdn{value: val, isSet: true}
}

func (v NullableAddrFqdn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddrFqdn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


