/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// NwdafEvent - Possible values are - SLICE_LOAD_LEVEL: Indicates that the event subscribed is load level information of Network Slice - NETWORK_PERFORMANCE: Indicates that the event subscribed is network performance information. - NF_LOAD: Indicates that the event subscribed is load level and status of one or several Network Functions. - SERVICE_EXPERIENCE: Indicates that the event subscribed is service experience. - UE_MOBILITY: Indicates that the event subscribed is UE mobility information. - UE_COMMUNICATION: Indicates that the event subscribed is UE communication information. - QOS_SUSTAINABILITY: Indicates that the event subscribed is QoS sustainability. - ABNORMAL_BEHAVIOUR: Indicates that the event subscribed is abnormal behaviour. - USER_DATA_CONGESTION: Indicates that the event subscribed is user data congestion information. - NSI_LOAD_LEVEL: Indicates that the event subscribed is load level information of Network Slice and the optionally associated Network Slice Instance - DN_PERFORMANCE: Indicates that the event subscribed is DN performance information. - DISPERSION: Indicates that the event subscribed is dispersion information. - RED_TRANS_EXP: Indicates that the event subscribed is redundant transmission experience. - WLAN_PERFORMANCE: Indicates that the event subscribed is WLAN performance.
type NwdafEvent string

// List of NwdafEvent
const (
	NWDAFEVENT_SLICE_LOAD_LEVEL     NwdafEvent = "SLICE_LOAD_LEVEL"
	NWDAFEVENT_NETWORK_PERFORMANCE  NwdafEvent = "NETWORK_PERFORMANCE"
	NWDAFEVENT_NF_LOAD              NwdafEvent = "NF_LOAD"
	NWDAFEVENT_SERVICE_EXPERIENCE   NwdafEvent = "SERVICE_EXPERIENCE"
	NWDAFEVENT_UE_MOBILITY          NwdafEvent = "UE_MOBILITY"
	NWDAFEVENT_UE_COMMUNICATION     NwdafEvent = "UE_COMMUNICATION"
	NWDAFEVENT_QOS_SUSTAINABILITY   NwdafEvent = "QOS_SUSTAINABILITY"
	NWDAFEVENT_ABNORMAL_BEHAVIOUR   NwdafEvent = "ABNORMAL_BEHAVIOUR"
	NWDAFEVENT_USER_DATA_CONGESTION NwdafEvent = "USER_DATA_CONGESTION"
	NWDAFEVENT_NSI_LOAD_LEVEL       NwdafEvent = "NSI_LOAD_LEVEL"
	NWDAFEVENT_DN_PERFORMANCE       NwdafEvent = "DN_PERFORMANCE"
	NWDAFEVENT_DISPERSION           NwdafEvent = "DISPERSION"
	NWDAFEVENT_RED_TRANS_EXP        NwdafEvent = "RED_TRANS_EXP"
	NWDAFEVENT_WLAN_PERFORMANCE     NwdafEvent = "WLAN_PERFORMANCE"
)

// AssertNwdafEventRequired checks if the required fields are not zero-ed
func AssertNwdafEventRequired(obj NwdafEvent) error {
	return nil
}

// AssertRecurseNwdafEventRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of NwdafEvent (e.g. [][]NwdafEvent), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNwdafEventRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNwdafEvent, ok := obj.(NwdafEvent)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNwdafEventRequired(aNwdafEvent)
	})
}
