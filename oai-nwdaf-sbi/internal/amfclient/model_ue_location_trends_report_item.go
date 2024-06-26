/*
Namf_EventExposure

AMF Event Exposure Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.2.0-alpha.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package amfclient

import (
	"encoding/json"
	"time"
)

// UeLocationTrendsReportItem Report Item for UE Location Trends event.
type UeLocationTrendsReportItem struct {
	Tai *Tai `json:"tai,omitempty"`
	Ncgi *Ncgi `json:"ncgi,omitempty"`
	Ecgi *Ecgi `json:"ecgi,omitempty"`
	N3gaLocation *N3gaLocation `json:"n3gaLocation,omitempty"`
	// indicating a time in seconds.
	Spacing int32 `json:"spacing"`
	// indicating a time in seconds.
	Duration int32 `json:"duration"`
	// string with format 'date-time' as defined in OpenAPI.
	Timestamp time.Time `json:"timestamp"`
}

// NewUeLocationTrendsReportItem instantiates a new UeLocationTrendsReportItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeLocationTrendsReportItem(spacing int32, duration int32, timestamp time.Time) *UeLocationTrendsReportItem {
	this := UeLocationTrendsReportItem{}
	this.Spacing = spacing
	this.Duration = duration
	this.Timestamp = timestamp
	return &this
}

// NewUeLocationTrendsReportItemWithDefaults instantiates a new UeLocationTrendsReportItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeLocationTrendsReportItemWithDefaults() *UeLocationTrendsReportItem {
	this := UeLocationTrendsReportItem{}
	return &this
}

// GetTai returns the Tai field value if set, zero value otherwise.
func (o *UeLocationTrendsReportItem) GetTai() Tai {
	if o == nil || o.Tai == nil {
		var ret Tai
		return ret
	}
	return *o.Tai
}

// GetTaiOk returns a tuple with the Tai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeLocationTrendsReportItem) GetTaiOk() (*Tai, bool) {
	if o == nil || o.Tai == nil {
		return nil, false
	}
	return o.Tai, true
}

// HasTai returns a boolean if a field has been set.
func (o *UeLocationTrendsReportItem) HasTai() bool {
	if o != nil && o.Tai != nil {
		return true
	}

	return false
}

// SetTai gets a reference to the given Tai and assigns it to the Tai field.
func (o *UeLocationTrendsReportItem) SetTai(v Tai) {
	o.Tai = &v
}

// GetNcgi returns the Ncgi field value if set, zero value otherwise.
func (o *UeLocationTrendsReportItem) GetNcgi() Ncgi {
	if o == nil || o.Ncgi == nil {
		var ret Ncgi
		return ret
	}
	return *o.Ncgi
}

// GetNcgiOk returns a tuple with the Ncgi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeLocationTrendsReportItem) GetNcgiOk() (*Ncgi, bool) {
	if o == nil || o.Ncgi == nil {
		return nil, false
	}
	return o.Ncgi, true
}

// HasNcgi returns a boolean if a field has been set.
func (o *UeLocationTrendsReportItem) HasNcgi() bool {
	if o != nil && o.Ncgi != nil {
		return true
	}

	return false
}

// SetNcgi gets a reference to the given Ncgi and assigns it to the Ncgi field.
func (o *UeLocationTrendsReportItem) SetNcgi(v Ncgi) {
	o.Ncgi = &v
}

// GetEcgi returns the Ecgi field value if set, zero value otherwise.
func (o *UeLocationTrendsReportItem) GetEcgi() Ecgi {
	if o == nil || o.Ecgi == nil {
		var ret Ecgi
		return ret
	}
	return *o.Ecgi
}

// GetEcgiOk returns a tuple with the Ecgi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeLocationTrendsReportItem) GetEcgiOk() (*Ecgi, bool) {
	if o == nil || o.Ecgi == nil {
		return nil, false
	}
	return o.Ecgi, true
}

// HasEcgi returns a boolean if a field has been set.
func (o *UeLocationTrendsReportItem) HasEcgi() bool {
	if o != nil && o.Ecgi != nil {
		return true
	}

	return false
}

// SetEcgi gets a reference to the given Ecgi and assigns it to the Ecgi field.
func (o *UeLocationTrendsReportItem) SetEcgi(v Ecgi) {
	o.Ecgi = &v
}

// GetN3gaLocation returns the N3gaLocation field value if set, zero value otherwise.
func (o *UeLocationTrendsReportItem) GetN3gaLocation() N3gaLocation {
	if o == nil || o.N3gaLocation == nil {
		var ret N3gaLocation
		return ret
	}
	return *o.N3gaLocation
}

// GetN3gaLocationOk returns a tuple with the N3gaLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeLocationTrendsReportItem) GetN3gaLocationOk() (*N3gaLocation, bool) {
	if o == nil || o.N3gaLocation == nil {
		return nil, false
	}
	return o.N3gaLocation, true
}

// HasN3gaLocation returns a boolean if a field has been set.
func (o *UeLocationTrendsReportItem) HasN3gaLocation() bool {
	if o != nil && o.N3gaLocation != nil {
		return true
	}

	return false
}

// SetN3gaLocation gets a reference to the given N3gaLocation and assigns it to the N3gaLocation field.
func (o *UeLocationTrendsReportItem) SetN3gaLocation(v N3gaLocation) {
	o.N3gaLocation = &v
}

// GetSpacing returns the Spacing field value
func (o *UeLocationTrendsReportItem) GetSpacing() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Spacing
}

// GetSpacingOk returns a tuple with the Spacing field value
// and a boolean to check if the value has been set.
func (o *UeLocationTrendsReportItem) GetSpacingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spacing, true
}

// SetSpacing sets field value
func (o *UeLocationTrendsReportItem) SetSpacing(v int32) {
	o.Spacing = v
}

// GetDuration returns the Duration field value
func (o *UeLocationTrendsReportItem) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *UeLocationTrendsReportItem) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *UeLocationTrendsReportItem) SetDuration(v int32) {
	o.Duration = v
}

// GetTimestamp returns the Timestamp field value
func (o *UeLocationTrendsReportItem) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *UeLocationTrendsReportItem) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *UeLocationTrendsReportItem) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o UeLocationTrendsReportItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tai != nil {
		toSerialize["tai"] = o.Tai
	}
	if o.Ncgi != nil {
		toSerialize["ncgi"] = o.Ncgi
	}
	if o.Ecgi != nil {
		toSerialize["ecgi"] = o.Ecgi
	}
	if o.N3gaLocation != nil {
		toSerialize["n3gaLocation"] = o.N3gaLocation
	}
	if true {
		toSerialize["spacing"] = o.Spacing
	}
	if true {
		toSerialize["duration"] = o.Duration
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableUeLocationTrendsReportItem struct {
	value *UeLocationTrendsReportItem
	isSet bool
}

func (v NullableUeLocationTrendsReportItem) Get() *UeLocationTrendsReportItem {
	return v.value
}

func (v *NullableUeLocationTrendsReportItem) Set(val *UeLocationTrendsReportItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUeLocationTrendsReportItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUeLocationTrendsReportItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeLocationTrendsReportItem(val *UeLocationTrendsReportItem) *NullableUeLocationTrendsReportItem {
	return &NullableUeLocationTrendsReportItem{value: val, isSet: true}
}

func (v NullableUeLocationTrendsReportItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeLocationTrendsReportItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


