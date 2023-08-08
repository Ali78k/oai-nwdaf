/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// SliceLoadLevelInformation - Contains load level information applicable for one or several slices.
type SliceLoadLevelInformation struct {

	// Load level information of the network slice and the optionally associated network slice instance.
	LoadLevelInformation int32 `json:"loadLevelInformation"`

	// Identification(s) of network slice to which the subscription applies.
	Snssais []Snssai `json:"snssais"`

	NumOfUes NumberAverage `json:"numOfUes,omitempty"`

	NumOfPduSess NumberAverage `json:"numOfPduSess,omitempty"`

	ExceedLoadLevelThrInd bool `json:"exceedLoadLevelThrInd,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty"`
}

// AssertSliceLoadLevelInformationRequired checks if the required fields are not zero-ed
func AssertSliceLoadLevelInformationRequired(obj SliceLoadLevelInformation) error {
	elements := map[string]interface{}{
		"loadLevelInformation": obj.LoadLevelInformation,
		"snssais":              obj.Snssais,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Snssais {
		if err := AssertSnssaiRequired(el); err != nil {
			return err
		}
	}
	if err := AssertNumberAverageRequired(obj.NumOfUes); err != nil {
		return err
	}
	if err := AssertNumberAverageRequired(obj.NumOfPduSess); err != nil {
		return err
	}
	return nil
}

// AssertRecurseSliceLoadLevelInformationRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of SliceLoadLevelInformation (e.g. [][]SliceLoadLevelInformation), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseSliceLoadLevelInformationRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aSliceLoadLevelInformation, ok := obj.(SliceLoadLevelInformation)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertSliceLoadLevelInformationRequired(aSliceLoadLevelInformation)
	})
}
