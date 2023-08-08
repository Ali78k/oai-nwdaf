/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

// ContextData - Contains context information related to analytics subscriptions corresponding with one or more context identifiers.
type ContextData struct {

	// List of items that contain context information corresponding with a context identifier.
	ContextElems []ContextElement `json:"contextElems"`
}

// AssertContextDataRequired checks if the required fields are not zero-ed
func AssertContextDataRequired(obj ContextData) error {
	elements := map[string]interface{}{
		"contextElems": obj.ContextElems,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.ContextElems {
		if err := AssertContextElementRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseContextDataRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ContextData (e.g. [][]ContextData), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseContextDataRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aContextData, ok := obj.(ContextData)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertContextDataRequired(aContextData)
	})
}
