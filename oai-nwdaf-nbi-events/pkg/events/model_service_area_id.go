/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// ServiceAreaId - Contains a Service Area Identifier as defined in 3GPP TS 23.003, clause 12.5.
type ServiceAreaId struct {
	PlmnId PlmnId `json:"plmnId"`

	// Location Area Code.
	Lac string `json:"lac"`

	// Service Area Code.
	Sac string `json:"sac"`
}

// AssertServiceAreaIdRequired checks if the required fields are not zero-ed
func AssertServiceAreaIdRequired(obj ServiceAreaId) error {
	elements := map[string]interface{}{
		"plmnId": obj.PlmnId,
		"lac":    obj.Lac,
		"sac":    obj.Sac,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertPlmnIdRequired(obj.PlmnId); err != nil {
		return err
	}
	return nil
}

// AssertRecurseServiceAreaIdRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ServiceAreaId (e.g. [][]ServiceAreaId), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseServiceAreaIdRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aServiceAreaId, ok := obj.(ServiceAreaId)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertServiceAreaIdRequired(aServiceAreaId)
	})
}
