/*
Nsmf_EventExposure

Session Management Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smfclient

import (
	"encoding/json"
	"time"
)

// NsmfEventExposure Represents an Individual SMF Notification Subscription resource. The serviveName property corresponds to the serviceName in the main body of the specification.
type NsmfEventExposure struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause 2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2 of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2 of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of 3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall only contain characters allowed according to the \"lower-with-hyphen\" naming convention defined in 3GPP TS 29.501. 
	Supi *string `json:"supi,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-<extid>, where <extid>  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi *string `json:"gpsi,omitempty"`
	// Any UE indication. This IE shall be present if the event subscription is applicable to any UE. Default value \"false\" is used, if not present.
	AnyUeInd *bool `json:"anyUeInd,omitempty"`
	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.
	GroupId *string `json:"groupId,omitempty"`
	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.  
	PduSeId *int32 `json:"pduSeId,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003; it shall contain either a DNN Network Identifier, or a full DNN with both the Network Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots (e.g. \"Label1.Label2.Label3\").
	Dnn *string `json:"dnn,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
	// Identifies an Individual SMF Notification Subscription. To enable that the value is used as part of a URI, the string shall only contain characters allowed according to the \"lower-with-hyphen\" naming convention defined in 3GPP TS 29.501. In an OpenAPI schema, the format shall be designated as \"SubId\".
	SubId *string `json:"subId,omitempty"`
	// Notification Correlation ID assigned by the NF service consumer.
	NotifId string `json:"notifId"`
	// String providing an URI formatted according to RFC 3986
	NotifUri string `json:"notifUri"`
	// Alternate or backup IPv4 address(es) where to send Notifications.
	AltNotifIpv4Addrs []string `json:"altNotifIpv4Addrs,omitempty"`
	// Alternate or backup IPv6 address(es) where to send Notifications.
	AltNotifIpv6Addrs []Ipv6Addr `json:"altNotifIpv6Addrs,omitempty"`
	// Alternate or backup FQDN(s) where to send Notifications.
	AltNotifFqdns []string `json:"altNotifFqdns,omitempty"`
	// Subscribed events
	EventSubs []EventSubscription `json:"eventSubs"`
	EventNotifs []EventNotification `json:"eventNotifs,omitempty"`
	ImmeRep *bool `json:"ImmeRep,omitempty"`
	NotifMethod *NotificationMethod `json:"notifMethod,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxReportNbr *int32 `json:"maxReportNbr,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty"`
	// indicating a time in seconds.
	RepPeriod *int32 `json:"repPeriod,omitempty"`
	Guami *Guami `json:"guami,omitempty"`
	ServiveName *ServiceName `json:"serviveName,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	SampRatio *int32 `json:"sampRatio,omitempty"`
	// Criteria for partitioning the UEs before applying the sampling ratio.
	PartitionCriteria []PartitioningCriteria `json:"partitionCriteria,omitempty"`
	// indicating a time in seconds.
	GrpRepTime *int32 `json:"grpRepTime,omitempty"`
	NotifFlag *NotificationFlag `json:"notifFlag,omitempty"`
}

// NewNsmfEventExposure instantiates a new NsmfEventExposure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNsmfEventExposure(notifId string, notifUri string, eventSubs []EventSubscription) *NsmfEventExposure {
	this := NsmfEventExposure{}
	this.NotifId = notifId
	this.NotifUri = notifUri
	this.EventSubs = eventSubs
	return &this
}

// NewNsmfEventExposureWithDefaults instantiates a new NsmfEventExposure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNsmfEventExposureWithDefaults() *NsmfEventExposure {
	this := NsmfEventExposure{}
	return &this
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetSupi() string {
	if o == nil || o.Supi == nil {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetSupiOk() (*string, bool) {
	if o == nil || o.Supi == nil {
		return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasSupi() bool {
	if o != nil && o.Supi != nil {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *NsmfEventExposure) SetSupi(v string) {
	o.Supi = &v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetGpsi() string {
	if o == nil || o.Gpsi == nil {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetGpsiOk() (*string, bool) {
	if o == nil || o.Gpsi == nil {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasGpsi() bool {
	if o != nil && o.Gpsi != nil {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *NsmfEventExposure) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetAnyUeInd returns the AnyUeInd field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetAnyUeInd() bool {
	if o == nil || o.AnyUeInd == nil {
		var ret bool
		return ret
	}
	return *o.AnyUeInd
}

// GetAnyUeIndOk returns a tuple with the AnyUeInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetAnyUeIndOk() (*bool, bool) {
	if o == nil || o.AnyUeInd == nil {
		return nil, false
	}
	return o.AnyUeInd, true
}

// HasAnyUeInd returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasAnyUeInd() bool {
	if o != nil && o.AnyUeInd != nil {
		return true
	}

	return false
}

// SetAnyUeInd gets a reference to the given bool and assigns it to the AnyUeInd field.
func (o *NsmfEventExposure) SetAnyUeInd(v bool) {
	o.AnyUeInd = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *NsmfEventExposure) SetGroupId(v string) {
	o.GroupId = &v
}

// GetPduSeId returns the PduSeId field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetPduSeId() int32 {
	if o == nil || o.PduSeId == nil {
		var ret int32
		return ret
	}
	return *o.PduSeId
}

// GetPduSeIdOk returns a tuple with the PduSeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetPduSeIdOk() (*int32, bool) {
	if o == nil || o.PduSeId == nil {
		return nil, false
	}
	return o.PduSeId, true
}

// HasPduSeId returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasPduSeId() bool {
	if o != nil && o.PduSeId != nil {
		return true
	}

	return false
}

// SetPduSeId gets a reference to the given int32 and assigns it to the PduSeId field.
func (o *NsmfEventExposure) SetPduSeId(v int32) {
	o.PduSeId = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetDnn() string {
	if o == nil || o.Dnn == nil {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetDnnOk() (*string, bool) {
	if o == nil || o.Dnn == nil {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasDnn() bool {
	if o != nil && o.Dnn != nil {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *NsmfEventExposure) SetDnn(v string) {
	o.Dnn = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetSnssai() Snssai {
	if o == nil || o.Snssai == nil {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || o.Snssai == nil {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasSnssai() bool {
	if o != nil && o.Snssai != nil {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *NsmfEventExposure) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetSubId returns the SubId field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetSubId() string {
	if o == nil || o.SubId == nil {
		var ret string
		return ret
	}
	return *o.SubId
}

// GetSubIdOk returns a tuple with the SubId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetSubIdOk() (*string, bool) {
	if o == nil || o.SubId == nil {
		return nil, false
	}
	return o.SubId, true
}

// HasSubId returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasSubId() bool {
	if o != nil && o.SubId != nil {
		return true
	}

	return false
}

// SetSubId gets a reference to the given string and assigns it to the SubId field.
func (o *NsmfEventExposure) SetSubId(v string) {
	o.SubId = &v
}

// GetNotifId returns the NotifId field value
func (o *NsmfEventExposure) GetNotifId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifId
}

// GetNotifIdOk returns a tuple with the NotifId field value
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetNotifIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifId, true
}

// SetNotifId sets field value
func (o *NsmfEventExposure) SetNotifId(v string) {
	o.NotifId = v
}

// GetNotifUri returns the NotifUri field value
func (o *NsmfEventExposure) GetNotifUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifUri
}

// GetNotifUriOk returns a tuple with the NotifUri field value
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetNotifUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifUri, true
}

// SetNotifUri sets field value
func (o *NsmfEventExposure) SetNotifUri(v string) {
	o.NotifUri = v
}

// GetAltNotifIpv4Addrs returns the AltNotifIpv4Addrs field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetAltNotifIpv4Addrs() []string {
	if o == nil || o.AltNotifIpv4Addrs == nil {
		var ret []string
		return ret
	}
	return o.AltNotifIpv4Addrs
}

// GetAltNotifIpv4AddrsOk returns a tuple with the AltNotifIpv4Addrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetAltNotifIpv4AddrsOk() ([]string, bool) {
	if o == nil || o.AltNotifIpv4Addrs == nil {
		return nil, false
	}
	return o.AltNotifIpv4Addrs, true
}

// HasAltNotifIpv4Addrs returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasAltNotifIpv4Addrs() bool {
	if o != nil && o.AltNotifIpv4Addrs != nil {
		return true
	}

	return false
}

// SetAltNotifIpv4Addrs gets a reference to the given []string and assigns it to the AltNotifIpv4Addrs field.
func (o *NsmfEventExposure) SetAltNotifIpv4Addrs(v []string) {
	o.AltNotifIpv4Addrs = v
}

// GetAltNotifIpv6Addrs returns the AltNotifIpv6Addrs field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetAltNotifIpv6Addrs() []Ipv6Addr {
	if o == nil || o.AltNotifIpv6Addrs == nil {
		var ret []Ipv6Addr
		return ret
	}
	return o.AltNotifIpv6Addrs
}

// GetAltNotifIpv6AddrsOk returns a tuple with the AltNotifIpv6Addrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetAltNotifIpv6AddrsOk() ([]Ipv6Addr, bool) {
	if o == nil || o.AltNotifIpv6Addrs == nil {
		return nil, false
	}
	return o.AltNotifIpv6Addrs, true
}

// HasAltNotifIpv6Addrs returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasAltNotifIpv6Addrs() bool {
	if o != nil && o.AltNotifIpv6Addrs != nil {
		return true
	}

	return false
}

// SetAltNotifIpv6Addrs gets a reference to the given []Ipv6Addr and assigns it to the AltNotifIpv6Addrs field.
func (o *NsmfEventExposure) SetAltNotifIpv6Addrs(v []Ipv6Addr) {
	o.AltNotifIpv6Addrs = v
}

// GetAltNotifFqdns returns the AltNotifFqdns field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetAltNotifFqdns() []string {
	if o == nil || o.AltNotifFqdns == nil {
		var ret []string
		return ret
	}
	return o.AltNotifFqdns
}

// GetAltNotifFqdnsOk returns a tuple with the AltNotifFqdns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetAltNotifFqdnsOk() ([]string, bool) {
	if o == nil || o.AltNotifFqdns == nil {
		return nil, false
	}
	return o.AltNotifFqdns, true
}

// HasAltNotifFqdns returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasAltNotifFqdns() bool {
	if o != nil && o.AltNotifFqdns != nil {
		return true
	}

	return false
}

// SetAltNotifFqdns gets a reference to the given []string and assigns it to the AltNotifFqdns field.
func (o *NsmfEventExposure) SetAltNotifFqdns(v []string) {
	o.AltNotifFqdns = v
}

// GetEventSubs returns the EventSubs field value
func (o *NsmfEventExposure) GetEventSubs() []EventSubscription {
	if o == nil {
		var ret []EventSubscription
		return ret
	}

	return o.EventSubs
}

// GetEventSubsOk returns a tuple with the EventSubs field value
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetEventSubsOk() ([]EventSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventSubs, true
}

// SetEventSubs sets field value
func (o *NsmfEventExposure) SetEventSubs(v []EventSubscription) {
	o.EventSubs = v
}

// GetEventNotifs returns the EventNotifs field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetEventNotifs() []EventNotification {
	if o == nil || o.EventNotifs == nil {
		var ret []EventNotification
		return ret
	}
	return o.EventNotifs
}

// GetEventNotifsOk returns a tuple with the EventNotifs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetEventNotifsOk() ([]EventNotification, bool) {
	if o == nil || o.EventNotifs == nil {
		return nil, false
	}
	return o.EventNotifs, true
}

// HasEventNotifs returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasEventNotifs() bool {
	if o != nil && o.EventNotifs != nil {
		return true
	}

	return false
}

// SetEventNotifs gets a reference to the given []EventNotification and assigns it to the EventNotifs field.
func (o *NsmfEventExposure) SetEventNotifs(v []EventNotification) {
	o.EventNotifs = v
}

// GetImmeRep returns the ImmeRep field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetImmeRep() bool {
	if o == nil || o.ImmeRep == nil {
		var ret bool
		return ret
	}
	return *o.ImmeRep
}

// GetImmeRepOk returns a tuple with the ImmeRep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetImmeRepOk() (*bool, bool) {
	if o == nil || o.ImmeRep == nil {
		return nil, false
	}
	return o.ImmeRep, true
}

// HasImmeRep returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasImmeRep() bool {
	if o != nil && o.ImmeRep != nil {
		return true
	}

	return false
}

// SetImmeRep gets a reference to the given bool and assigns it to the ImmeRep field.
func (o *NsmfEventExposure) SetImmeRep(v bool) {
	o.ImmeRep = &v
}

// GetNotifMethod returns the NotifMethod field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetNotifMethod() NotificationMethod {
	if o == nil || o.NotifMethod == nil {
		var ret NotificationMethod
		return ret
	}
	return *o.NotifMethod
}

// GetNotifMethodOk returns a tuple with the NotifMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetNotifMethodOk() (*NotificationMethod, bool) {
	if o == nil || o.NotifMethod == nil {
		return nil, false
	}
	return o.NotifMethod, true
}

// HasNotifMethod returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasNotifMethod() bool {
	if o != nil && o.NotifMethod != nil {
		return true
	}

	return false
}

// SetNotifMethod gets a reference to the given NotificationMethod and assigns it to the NotifMethod field.
func (o *NsmfEventExposure) SetNotifMethod(v NotificationMethod) {
	o.NotifMethod = &v
}

// GetMaxReportNbr returns the MaxReportNbr field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetMaxReportNbr() int32 {
	if o == nil || o.MaxReportNbr == nil {
		var ret int32
		return ret
	}
	return *o.MaxReportNbr
}

// GetMaxReportNbrOk returns a tuple with the MaxReportNbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetMaxReportNbrOk() (*int32, bool) {
	if o == nil || o.MaxReportNbr == nil {
		return nil, false
	}
	return o.MaxReportNbr, true
}

// HasMaxReportNbr returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasMaxReportNbr() bool {
	if o != nil && o.MaxReportNbr != nil {
		return true
	}

	return false
}

// SetMaxReportNbr gets a reference to the given int32 and assigns it to the MaxReportNbr field.
func (o *NsmfEventExposure) SetMaxReportNbr(v int32) {
	o.MaxReportNbr = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetExpiry() time.Time {
	if o == nil || o.Expiry == nil {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetExpiryOk() (*time.Time, bool) {
	if o == nil || o.Expiry == nil {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasExpiry() bool {
	if o != nil && o.Expiry != nil {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *NsmfEventExposure) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetRepPeriod returns the RepPeriod field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetRepPeriod() int32 {
	if o == nil || o.RepPeriod == nil {
		var ret int32
		return ret
	}
	return *o.RepPeriod
}

// GetRepPeriodOk returns a tuple with the RepPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetRepPeriodOk() (*int32, bool) {
	if o == nil || o.RepPeriod == nil {
		return nil, false
	}
	return o.RepPeriod, true
}

// HasRepPeriod returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasRepPeriod() bool {
	if o != nil && o.RepPeriod != nil {
		return true
	}

	return false
}

// SetRepPeriod gets a reference to the given int32 and assigns it to the RepPeriod field.
func (o *NsmfEventExposure) SetRepPeriod(v int32) {
	o.RepPeriod = &v
}

// GetGuami returns the Guami field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetGuami() Guami {
	if o == nil || o.Guami == nil {
		var ret Guami
		return ret
	}
	return *o.Guami
}

// GetGuamiOk returns a tuple with the Guami field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetGuamiOk() (*Guami, bool) {
	if o == nil || o.Guami == nil {
		return nil, false
	}
	return o.Guami, true
}

// HasGuami returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasGuami() bool {
	if o != nil && o.Guami != nil {
		return true
	}

	return false
}

// SetGuami gets a reference to the given Guami and assigns it to the Guami field.
func (o *NsmfEventExposure) SetGuami(v Guami) {
	o.Guami = &v
}

// GetServiveName returns the ServiveName field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetServiveName() ServiceName {
	if o == nil || o.ServiveName == nil {
		var ret ServiceName
		return ret
	}
	return *o.ServiveName
}

// GetServiveNameOk returns a tuple with the ServiveName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetServiveNameOk() (*ServiceName, bool) {
	if o == nil || o.ServiveName == nil {
		return nil, false
	}
	return o.ServiveName, true
}

// HasServiveName returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasServiveName() bool {
	if o != nil && o.ServiveName != nil {
		return true
	}

	return false
}

// SetServiveName gets a reference to the given ServiceName and assigns it to the ServiveName field.
func (o *NsmfEventExposure) SetServiveName(v ServiceName) {
	o.ServiveName = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetSupportedFeatures() string {
	if o == nil || o.SupportedFeatures == nil {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || o.SupportedFeatures == nil {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasSupportedFeatures() bool {
	if o != nil && o.SupportedFeatures != nil {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *NsmfEventExposure) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetSampRatio returns the SampRatio field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetSampRatio() int32 {
	if o == nil || o.SampRatio == nil {
		var ret int32
		return ret
	}
	return *o.SampRatio
}

// GetSampRatioOk returns a tuple with the SampRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetSampRatioOk() (*int32, bool) {
	if o == nil || o.SampRatio == nil {
		return nil, false
	}
	return o.SampRatio, true
}

// HasSampRatio returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasSampRatio() bool {
	if o != nil && o.SampRatio != nil {
		return true
	}

	return false
}

// SetSampRatio gets a reference to the given int32 and assigns it to the SampRatio field.
func (o *NsmfEventExposure) SetSampRatio(v int32) {
	o.SampRatio = &v
}

// GetPartitionCriteria returns the PartitionCriteria field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetPartitionCriteria() []PartitioningCriteria {
	if o == nil || o.PartitionCriteria == nil {
		var ret []PartitioningCriteria
		return ret
	}
	return o.PartitionCriteria
}

// GetPartitionCriteriaOk returns a tuple with the PartitionCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetPartitionCriteriaOk() ([]PartitioningCriteria, bool) {
	if o == nil || o.PartitionCriteria == nil {
		return nil, false
	}
	return o.PartitionCriteria, true
}

// HasPartitionCriteria returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasPartitionCriteria() bool {
	if o != nil && o.PartitionCriteria != nil {
		return true
	}

	return false
}

// SetPartitionCriteria gets a reference to the given []PartitioningCriteria and assigns it to the PartitionCriteria field.
func (o *NsmfEventExposure) SetPartitionCriteria(v []PartitioningCriteria) {
	o.PartitionCriteria = v
}

// GetGrpRepTime returns the GrpRepTime field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetGrpRepTime() int32 {
	if o == nil || o.GrpRepTime == nil {
		var ret int32
		return ret
	}
	return *o.GrpRepTime
}

// GetGrpRepTimeOk returns a tuple with the GrpRepTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetGrpRepTimeOk() (*int32, bool) {
	if o == nil || o.GrpRepTime == nil {
		return nil, false
	}
	return o.GrpRepTime, true
}

// HasGrpRepTime returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasGrpRepTime() bool {
	if o != nil && o.GrpRepTime != nil {
		return true
	}

	return false
}

// SetGrpRepTime gets a reference to the given int32 and assigns it to the GrpRepTime field.
func (o *NsmfEventExposure) SetGrpRepTime(v int32) {
	o.GrpRepTime = &v
}

// GetNotifFlag returns the NotifFlag field value if set, zero value otherwise.
func (o *NsmfEventExposure) GetNotifFlag() NotificationFlag {
	if o == nil || o.NotifFlag == nil {
		var ret NotificationFlag
		return ret
	}
	return *o.NotifFlag
}

// GetNotifFlagOk returns a tuple with the NotifFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsmfEventExposure) GetNotifFlagOk() (*NotificationFlag, bool) {
	if o == nil || o.NotifFlag == nil {
		return nil, false
	}
	return o.NotifFlag, true
}

// HasNotifFlag returns a boolean if a field has been set.
func (o *NsmfEventExposure) HasNotifFlag() bool {
	if o != nil && o.NotifFlag != nil {
		return true
	}

	return false
}

// SetNotifFlag gets a reference to the given NotificationFlag and assigns it to the NotifFlag field.
func (o *NsmfEventExposure) SetNotifFlag(v NotificationFlag) {
	o.NotifFlag = &v
}

func (o NsmfEventExposure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Supi != nil {
		toSerialize["supi"] = o.Supi
	}
	if o.Gpsi != nil {
		toSerialize["gpsi"] = o.Gpsi
	}
	if o.AnyUeInd != nil {
		toSerialize["anyUeInd"] = o.AnyUeInd
	}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if o.PduSeId != nil {
		toSerialize["pduSeId"] = o.PduSeId
	}
	if o.Dnn != nil {
		toSerialize["dnn"] = o.Dnn
	}
	if o.Snssai != nil {
		toSerialize["snssai"] = o.Snssai
	}
	if o.SubId != nil {
		toSerialize["subId"] = o.SubId
	}
	if true {
		toSerialize["notifId"] = o.NotifId
	}
	if true {
		toSerialize["notifUri"] = o.NotifUri
	}
	if o.AltNotifIpv4Addrs != nil {
		toSerialize["altNotifIpv4Addrs"] = o.AltNotifIpv4Addrs
	}
	if o.AltNotifIpv6Addrs != nil {
		toSerialize["altNotifIpv6Addrs"] = o.AltNotifIpv6Addrs
	}
	if o.AltNotifFqdns != nil {
		toSerialize["altNotifFqdns"] = o.AltNotifFqdns
	}
	if true {
		toSerialize["eventSubs"] = o.EventSubs
	}
	if o.EventNotifs != nil {
		toSerialize["eventNotifs"] = o.EventNotifs
	}
	if o.ImmeRep != nil {
		toSerialize["ImmeRep"] = o.ImmeRep
	}
	if o.NotifMethod != nil {
		toSerialize["notifMethod"] = o.NotifMethod
	}
	if o.MaxReportNbr != nil {
		toSerialize["maxReportNbr"] = o.MaxReportNbr
	}
	if o.Expiry != nil {
		toSerialize["expiry"] = o.Expiry
	}
	if o.RepPeriod != nil {
		toSerialize["repPeriod"] = o.RepPeriod
	}
	if o.Guami != nil {
		toSerialize["guami"] = o.Guami
	}
	if o.ServiveName != nil {
		toSerialize["serviveName"] = o.ServiveName
	}
	if o.SupportedFeatures != nil {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if o.SampRatio != nil {
		toSerialize["sampRatio"] = o.SampRatio
	}
	if o.PartitionCriteria != nil {
		toSerialize["partitionCriteria"] = o.PartitionCriteria
	}
	if o.GrpRepTime != nil {
		toSerialize["grpRepTime"] = o.GrpRepTime
	}
	if o.NotifFlag != nil {
		toSerialize["notifFlag"] = o.NotifFlag
	}
	return json.Marshal(toSerialize)
}

type NullableNsmfEventExposure struct {
	value *NsmfEventExposure
	isSet bool
}

func (v NullableNsmfEventExposure) Get() *NsmfEventExposure {
	return v.value
}

func (v *NullableNsmfEventExposure) Set(val *NsmfEventExposure) {
	v.value = val
	v.isSet = true
}

func (v NullableNsmfEventExposure) IsSet() bool {
	return v.isSet
}

func (v *NullableNsmfEventExposure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNsmfEventExposure(val *NsmfEventExposure) *NullableNsmfEventExposure {
	return &NullableNsmfEventExposure{value: val, isSet: true}
}

func (v NullableNsmfEventExposure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNsmfEventExposure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

