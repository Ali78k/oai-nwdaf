/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.0.0-alpha.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nbiml

// TargetUeInformation - Identifies the target UE information.
type TargetUeInformation struct {
	AnyUe bool `json:"anyUe,omitempty"`

	Supis []string `json:"supis,omitempty"`

	Gpsis []string `json:"gpsis,omitempty"`

	IntGroupIds []string `json:"intGroupIds,omitempty"`
}

// AssertTargetUeInformationRequired checks if the required fields are not zero-ed
func AssertTargetUeInformationRequired(obj TargetUeInformation) error {
	return nil
}

// AssertRecurseTargetUeInformationRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TargetUeInformation (e.g. [][]TargetUeInformation), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTargetUeInformationRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTargetUeInformation, ok := obj.(TargetUeInformation)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTargetUeInformationRequired(aTargetUeInformation)
	})
}
