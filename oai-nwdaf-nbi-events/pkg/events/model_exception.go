/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// Exception - Represents the Exception information.
type Exception struct {
	ExcepId ExceptionId `json:"excepId"`

	ExcepLevel int32 `json:"excepLevel,omitempty"`

	ExcepTrend ExceptionTrend `json:"excepTrend,omitempty"`
}

// AssertExceptionRequired checks if the required fields are not zero-ed
func AssertExceptionRequired(obj Exception) error {
	elements := map[string]interface{}{
		"excepId": obj.ExcepId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertExceptionIdRequired(obj.ExcepId); err != nil {
		return err
	}
	if err := AssertExceptionTrendRequired(obj.ExcepTrend); err != nil {
		return err
	}
	return nil
}

// AssertRecurseExceptionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Exception (e.g. [][]Exception), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseExceptionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aException, ok := obj.(Exception)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertExceptionRequired(aException)
	})
}
