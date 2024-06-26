/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.0.0-alpha.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nbiml

type DnPerfOrderingCriterionAnyOf string

// List of DnPerfOrderingCriterionAnyOf
const (
	DNPERFORDERINGCRITERIONANYOF_AVERAGE_TRAFFIC_RATE     DnPerfOrderingCriterionAnyOf = "AVERAGE_TRAFFIC_RATE"
	DNPERFORDERINGCRITERIONANYOF_MAXIMUM_TRAFFIC_RATE     DnPerfOrderingCriterionAnyOf = "MAXIMUM_TRAFFIC_RATE"
	DNPERFORDERINGCRITERIONANYOF_AVERAGE_PACKET_DELAY     DnPerfOrderingCriterionAnyOf = "AVERAGE_PACKET_DELAY"
	DNPERFORDERINGCRITERIONANYOF_MAXIMUM_PACKET_DELAY     DnPerfOrderingCriterionAnyOf = "MAXIMUM_PACKET_DELAY"
	DNPERFORDERINGCRITERIONANYOF_AVERAGE_PACKET_LOSS_RATE DnPerfOrderingCriterionAnyOf = "AVERAGE_PACKET_LOSS_RATE"
)

// AssertDnPerfOrderingCriterionAnyOfRequired checks if the required fields are not zero-ed
func AssertDnPerfOrderingCriterionAnyOfRequired(obj DnPerfOrderingCriterionAnyOf) error {
	return nil
}

// AssertRecurseDnPerfOrderingCriterionAnyOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of DnPerfOrderingCriterionAnyOf (e.g. [][]DnPerfOrderingCriterionAnyOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseDnPerfOrderingCriterionAnyOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aDnPerfOrderingCriterionAnyOf, ok := obj.(DnPerfOrderingCriterionAnyOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertDnPerfOrderingCriterionAnyOfRequired(aDnPerfOrderingCriterionAnyOf)
	})
}
