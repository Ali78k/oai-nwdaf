/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

import (
	"time"
)

// AdditionalMeasurement - Represents additional measurement information.
type AdditionalMeasurement struct {
	UnexpLoc NetworkAreaInfo `json:"unexpLoc,omitempty"`

	UnexpFlowTeps []IpEthFlowDescription `json:"unexpFlowTeps,omitempty"`

	UnexpWakes []time.Time `json:"unexpWakes,omitempty"`

	DdosAttack AddressList `json:"ddosAttack,omitempty"`

	WrgDest AddressList `json:"wrgDest,omitempty"`

	Circums []CircumstanceDescription `json:"circums,omitempty"`
}

// AssertAdditionalMeasurementRequired checks if the required fields are not zero-ed
func AssertAdditionalMeasurementRequired(obj AdditionalMeasurement) error {
	if err := AssertNetworkAreaInfoRequired(obj.UnexpLoc); err != nil {
		return err
	}
	for _, el := range obj.UnexpFlowTeps {
		if err := AssertIpEthFlowDescriptionRequired(el); err != nil {
			return err
		}
	}
	if err := AssertAddressListRequired(obj.DdosAttack); err != nil {
		return err
	}
	if err := AssertAddressListRequired(obj.WrgDest); err != nil {
		return err
	}
	for _, el := range obj.Circums {
		if err := AssertCircumstanceDescriptionRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseAdditionalMeasurementRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AdditionalMeasurement (e.g. [][]AdditionalMeasurement), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAdditionalMeasurementRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAdditionalMeasurement, ok := obj.(AdditionalMeasurement)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAdditionalMeasurementRequired(aAdditionalMeasurement)
	})
}