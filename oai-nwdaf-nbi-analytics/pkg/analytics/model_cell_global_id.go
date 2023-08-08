/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

// CellGlobalId - Contains a Cell Global Identification as defined in 3GPP TS 23.003, clause 4.3.1.
type CellGlobalId struct {
	PlmnId PlmnId `json:"plmnId"`

	Lac string `json:"lac"`

	CellId string `json:"cellId"`
}

// AssertCellGlobalIdRequired checks if the required fields are not zero-ed
func AssertCellGlobalIdRequired(obj CellGlobalId) error {
	elements := map[string]interface{}{
		"plmnId": obj.PlmnId,
		"lac":    obj.Lac,
		"cellId": obj.CellId,
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

// AssertRecurseCellGlobalIdRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CellGlobalId (e.g. [][]CellGlobalId), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCellGlobalIdRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCellGlobalId, ok := obj.(CellGlobalId)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCellGlobalIdRequired(aCellGlobalId)
	})
}
