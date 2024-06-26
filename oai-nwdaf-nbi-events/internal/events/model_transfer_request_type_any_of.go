/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

type TransferRequestTypeAnyOf string

// List of TransferRequestTypeAnyOf
const (
	TRANSFERREQUESTTYPEANYOF_PREPARE  TransferRequestTypeAnyOf = "PREPARE"
	TRANSFERREQUESTTYPEANYOF_TRANSFER TransferRequestTypeAnyOf = "TRANSFER"
)

// AssertTransferRequestTypeAnyOfRequired checks if the required fields are not zero-ed
func AssertTransferRequestTypeAnyOfRequired(obj TransferRequestTypeAnyOf) error {
	return nil
}

// AssertRecurseTransferRequestTypeAnyOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TransferRequestTypeAnyOf (e.g. [][]TransferRequestTypeAnyOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTransferRequestTypeAnyOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTransferRequestTypeAnyOf, ok := obj.(TransferRequestTypeAnyOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTransferRequestTypeAnyOfRequired(aTransferRequestTypeAnyOf)
	})
}
