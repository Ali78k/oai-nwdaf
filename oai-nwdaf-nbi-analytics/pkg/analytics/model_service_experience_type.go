/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

// ServiceExperienceType - Possible values are   - VOICE: Indicates that the service experience analytics is for voice service.   - VIDEO: Indicates that the service experience analytics is for video service.
type ServiceExperienceType struct {
}

// AssertServiceExperienceTypeRequired checks if the required fields are not zero-ed
func AssertServiceExperienceTypeRequired(obj ServiceExperienceType) error {
	return nil
}

// AssertRecurseServiceExperienceTypeRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ServiceExperienceType (e.g. [][]ServiceExperienceType), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseServiceExperienceTypeRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aServiceExperienceType, ok := obj.(ServiceExperienceType)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertServiceExperienceTypeRequired(aServiceExperienceType)
	})
}
