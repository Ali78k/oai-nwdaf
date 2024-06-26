/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

// RequestedContext - Contains types of analytics context information.
type RequestedContext struct {

	// List of analytics context types.
	Contexts []ContextType `json:"contexts"`
}

// AssertRequestedContextRequired checks if the required fields are not zero-ed
func AssertRequestedContextRequired(obj RequestedContext) error {
	elements := map[string]interface{}{
		"contexts": obj.Contexts,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Contexts {
		if err := AssertContextTypeRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseRequestedContextRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of RequestedContext (e.g. [][]RequestedContext), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseRequestedContextRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aRequestedContext, ok := obj.(RequestedContext)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertRequestedContextRequired(aRequestedContext)
	})
}
