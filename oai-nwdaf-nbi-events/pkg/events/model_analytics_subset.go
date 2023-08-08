/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

// AnalyticsSubset - Possible values are   - NUM_OF_UE_REG: The number of UE registered. This value is only applicable to NSI_LOAD_LEVEL event, SLICE_LOAD_LEVEL event and LOAD_LEVEL_INFORMATION event.   - NUM_OF_PDU_SESS_ESTBL: The number of PDU sessions established. This value is only applicable to NSI_LOAD_LEVEL event, SLICE_LOAD_LEVEL event and LOAD_LEVEL_INFORMATION event.   - RES_USAGE: The current usage of the virtual resources assigned to the NF instances belonging to a particular network slice instance. This value is only applicable to NSI_LOAD_LEVEL event.   - NUM_OF_EXCEED_RES_USAGE_LOAD_LEVEL_THR: The number of times the resource usage threshold of the network slice instance is reached or exceeded if a threshold value is provided by the consumer. This value is only applicable to NSI_LOAD_LEVEL event.   - PERIOD_OF_EXCEED_RES_USAGE_LOAD_LEVEL_THR: The time interval between each time the threshold being met or exceeded on the network slice (instance). This value is only applicable to NSI_LOAD_LEVEL event.   - EXCEED_LOAD_LEVEL_THR_IND: Whether the Load Level Threshold is met or exceeded by the statistics value. This value is only applicable to NSI_LOAD_LEVEL event, SLICE_LOAD_LEVEL event and LOAD_LEVEL_INFORMATION event.   - LIST_OF_TOP_APP_UL: The list of applications that contribute the most to the traffic in the UL direction. This value is only applicable to USER_DATA_CONGESTION event.   - LIST_OF_TOP_APP_DL: The list of applications that contribute the most to the traffic in the DL direction. This value is only applicable to USER_DATA_CONGESTION event.   - NF_STATUS: The availability status of the NF on the Analytics target period, expressed as a percentage of time per status value (registered, suspended, undiscoverable). This value is only applicable to NF_LOAD event.   - NF_RESOURCE_USAGE: The average usage of assigned resources (CPU, memory, storage). This value is only applicable to NF_LOAD event.   - NF_LOAD: The average load of the NF instance over the Analytics target period. This value is only applicable to NF_LOAD event.   - NF_PEAK_LOAD: The maximum load of the NF instance over the Analytics target period. This value is only applicable to NF_LOAD event.   - DISPER_AMOUNT: Indicates the dispersion amount of the reported data volume or transaction dispersion type. This value is only applicable to DISPERSION event.   - DISPER_CLASS: Indicates the dispersion mobility class: fixed, camper, traveller upon set its usage threshold, and/or the top-heavy class upon set its percentile rating threshold. This value is only applicable to DISPERSION event.   - RANKING: Data/transaction usage ranking high (i.e.value 1), medium (2) or low (3). This value is only applicable to DISPERSION event.   - PERCENTILE_RANKING: Percentile ranking of the target UE in the Cumulative Distribution Function of data usage for the population of all UEs. This value is only applicable to DISPERSION event.   - RSSI: Indicated the RSSI in the unit of dBm. This value is only applicable to WLAN_PERFORMANCE event.   - RTT: Indicates the RTT in the unit of millisecond. This value is only applicable to WLAN_PERFORMANCE event.   - TRAFFIC_INFO: Traffic information including UL/DL data rate and/or Traffic volume. This value is only applicable to WLAN_PERFORMANCE event.   - NUMBER_OF_UES: Number of UEs observed for the SSID. This value is only applicable to WLAN_PERFORMANCE event.   - APP_LIST_FOR_UE_COMM: The analytics of the application list used by UE. This value is only applicable to UE_COMM event.   - N4_SESS_INACT_TIMER_FOR_UE_COMM: The N4 Session inactivity timer. This value is only applicable to UE_COMM event.   - AVG_TRAFFIC_RATE: Indicates average traffic rate. This value is only applicable to DN_PERFORMANCE event.   - MAX_TRAFFIC_RATE: Indicates maximum traffic rate. This value is only applicable to DN_PERFORMANCE event.   - AVG_PACKET_DELAY: Indicates average Packet Delay. This value is only applicable to DN_PERFORMANCE event.   - MAX_PACKET_DELAY: Indicates maximum Packet Delay. This value is only applicable to DN_PERFORMANCE event.   - AVG_PACKET_LOSS_RATE: Indicates average Loss Rate. This value is only applicable to DN_PERFORMANCE event.
type AnalyticsSubset struct {
}

// AssertAnalyticsSubsetRequired checks if the required fields are not zero-ed
func AssertAnalyticsSubsetRequired(obj AnalyticsSubset) error {
	return nil
}

// AssertRecurseAnalyticsSubsetRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AnalyticsSubset (e.g. [][]AnalyticsSubset), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAnalyticsSubsetRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAnalyticsSubset, ok := obj.(AnalyticsSubset)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAnalyticsSubsetRequired(aAnalyticsSubset)
	})
}
