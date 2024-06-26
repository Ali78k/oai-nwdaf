/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

// ScheduledCommunicationTime - Represents an offered scheduled communication time.
type ScheduledCommunicationTime struct {

	// Identifies the day(s) of the week. If absent, it indicates every day of the week.
	DaysOfWeek []int32 `json:"daysOfWeek,omitempty"`

	// String with format partial-time or full-time as defined in subclause 5.6 of IETF RFC 3339. Examples, 20:15:00, 20:15:00-08:00 (for 8 hours behind UTC).
	TimeOfDayStart string `json:"timeOfDayStart,omitempty"`

	// String with format partial-time or full-time as defined in subclause 5.6 of IETF RFC 3339. Examples, 20:15:00, 20:15:00-08:00 (for 8 hours behind UTC).
	TimeOfDayEnd string `json:"timeOfDayEnd,omitempty"`
}

// AssertScheduledCommunicationTimeRequired checks if the required fields are not zero-ed
func AssertScheduledCommunicationTimeRequired(obj ScheduledCommunicationTime) error {
	return nil
}

// AssertRecurseScheduledCommunicationTimeRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ScheduledCommunicationTime (e.g. [][]ScheduledCommunicationTime), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseScheduledCommunicationTimeRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aScheduledCommunicationTime, ok := obj.(ScheduledCommunicationTime)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertScheduledCommunicationTimeRequired(aScheduledCommunicationTime)
	})
}
