/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.0.0-alpha.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nbiml

// ScheduledCommunicationType - Possible values are: -DOWNLINK_ONLY: Downlink only -UPLINK_ONLY: Uplink only -BIDIRECTIONA: Bi-directional
type ScheduledCommunicationType struct {
}

// AssertScheduledCommunicationTypeRequired checks if the required fields are not zero-ed
func AssertScheduledCommunicationTypeRequired(obj ScheduledCommunicationType) error {
	return nil
}

// AssertRecurseScheduledCommunicationTypeRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ScheduledCommunicationType (e.g. [][]ScheduledCommunicationType), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseScheduledCommunicationTypeRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aScheduledCommunicationType, ok := obj.(ScheduledCommunicationType)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertScheduledCommunicationTypeRequired(aScheduledCommunicationType)
	})
}
