/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

// ContextType - Possible values are - PENDING_ANALYTICS: Represents context information that relates to pending output analytics. - HISTORICAL_ANALYTICS: Represents context information that relates to historical output analytics. - AGGR_SUBS: Represents context information about the analytics subscriptions that an NWDAF has with other NWDAFs that collectively serve an analytics subscription. - DATA: Represents context information about historical data that is available. - AGGR_INFO: Represents context information that is related to aggregation of analytics from multiple NWDAF subscriptions. - ML_MODELS: Represents context information about used ML models.
type ContextType struct {
}

// AssertContextTypeRequired checks if the required fields are not zero-ed
func AssertContextTypeRequired(obj ContextType) error {
	return nil
}

// AssertRecurseContextTypeRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ContextType (e.g. [][]ContextType), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseContextTypeRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aContextType, ok := obj.(ContextType)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertContextTypeRequired(aContextType)
	})
}