/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// CongestionInfo1 - Represents the congestion information.
type CongestionInfo1 struct {
	CongType CongestionType `json:"congType"`

	TimeIntev TimeWindow `json:"timeIntev"`

	Nsi ThresholdLevel1 `json:"nsi"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty"`

	TopAppListUl []TopApplication1 `json:"topAppListUl,omitempty"`

	TopAppListDl []TopApplication1 `json:"topAppListDl,omitempty"`
}

// AssertCongestionInfo1Required checks if the required fields are not zero-ed
func AssertCongestionInfo1Required(obj CongestionInfo1) error {
	elements := map[string]interface{}{
		"congType":  obj.CongType,
		"timeIntev": obj.TimeIntev,
		"nsi":       obj.Nsi,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertCongestionTypeRequired(obj.CongType); err != nil {
		return err
	}
	if err := AssertTimeWindowRequired(obj.TimeIntev); err != nil {
		return err
	}
	if err := AssertThresholdLevel1Required(obj.Nsi); err != nil {
		return err
	}
	for _, el := range obj.TopAppListUl {
		if err := AssertTopApplication1Required(el); err != nil {
			return err
		}
	}
	for _, el := range obj.TopAppListDl {
		if err := AssertTopApplication1Required(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseCongestionInfo1Required recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CongestionInfo1 (e.g. [][]CongestionInfo1), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCongestionInfo1Required(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCongestionInfo1, ok := obj.(CongestionInfo1)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCongestionInfo1Required(aCongestionInfo1)
	})
}
