/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// Exception1 - Represents the Exception information.
type Exception1 struct {
	ExcepId ExceptionId `json:"excepId"`

	ExcepLevel int32 `json:"excepLevel,omitempty"`

	ExcepTrend ExceptionTrend `json:"excepTrend,omitempty"`
}

// AssertException1Required checks if the required fields are not zero-ed
func AssertException1Required(obj Exception1) error {
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

// AssertRecurseException1Required recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Exception1 (e.g. [][]Exception1), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseException1Required(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aException1, ok := obj.(Exception1)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertException1Required(aException1)
	})
}
