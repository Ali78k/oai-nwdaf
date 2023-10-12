/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// InvalidParam - It contains an invalid parameter and a related description.
type InvalidParam struct {

	// If the invalid parameter is an attribute in a JSON body, this IE shall contain the attribute's name and shall be encoded as a JSON Pointer. If the invalid parameter is an HTTP header, this IE shall be formatted as the concatenation of the string \"header \" plus the name of such header. If the invalid parameter is a query parameter, this IE shall be formatted as the concatenation of the string \"query \" plus the name of such query parameter. If the invalid parameter is a variable part in the path of a resource URI, this IE shall contain the name of the variable, including the symbols \"{\" and \"}\" used in OpenAPI specification as the notation to represent variable path segments.
	Param string `json:"param"`

	// A human-readable reason, e.g. \"must be a positive integer\". In cases involving failed operations in a PATCH request, the reason string should identify the operation that failed using the operation's array index to assist in correlation of the invalid parameter with the failed operation, e.g.\" Replacement value invalid for attribute (failed operation index= 4)\"
	Reason string `json:"reason,omitempty"`
}

// AssertInvalidParamRequired checks if the required fields are not zero-ed
func AssertInvalidParamRequired(obj InvalidParam) error {
	elements := map[string]interface{}{
		"param": obj.Param,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseInvalidParamRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of InvalidParam (e.g. [][]InvalidParam), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseInvalidParamRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aInvalidParam, ok := obj.(InvalidParam)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertInvalidParamRequired(aInvalidParam)
	})
}