/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// AnalyticsMetadataIndication1 - Contains analytics metadata information requested to be used during analytics generation.
type AnalyticsMetadataIndication1 struct {
	DataWindow TimeWindow `json:"dataWindow,omitempty"`

	DataStatProps []DatasetStatisticalProperty `json:"dataStatProps,omitempty"`

	Strategy OutputStrategy `json:"strategy,omitempty"`

	AggrNwdafIds []string `json:"aggrNwdafIds,omitempty"`
}

// AssertAnalyticsMetadataIndication1Required checks if the required fields are not zero-ed
func AssertAnalyticsMetadataIndication1Required(obj AnalyticsMetadataIndication1) error {
	if err := AssertTimeWindowRequired(obj.DataWindow); err != nil {
		return err
	}
	for _, el := range obj.DataStatProps {
		if err := AssertDatasetStatisticalPropertyRequired(el); err != nil {
			return err
		}
	}
	if err := AssertOutputStrategyRequired(obj.Strategy); err != nil {
		return err
	}
	return nil
}

// AssertRecurseAnalyticsMetadataIndication1Required recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AnalyticsMetadataIndication1 (e.g. [][]AnalyticsMetadataIndication1), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAnalyticsMetadataIndication1Required(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAnalyticsMetadataIndication1, ok := obj.(AnalyticsMetadataIndication1)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAnalyticsMetadataIndication1Required(aAnalyticsMetadataIndication1)
	})
}
