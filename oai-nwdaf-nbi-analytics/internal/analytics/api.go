/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package analytics

import (
	"context"
	"net/http"
)

// NWDAFAnalyticsDocumentApiRouter defines the required methods for binding the api requests to a responses for the NWDAFAnalyticsDocumentApi
// The NWDAFAnalyticsDocumentApiRouter implementation should parse necessary information from the http request,
// pass the data to a NWDAFAnalyticsDocumentApiServicer to perform the required actions, then write the service results to the http response.
type NWDAFAnalyticsDocumentApiRouter interface {
	GetNWDAFAnalytics(http.ResponseWriter, *http.Request)
}

// NWDAFContextDocumentApiRouter defines the required methods for binding the api requests to a responses for the NWDAFContextDocumentApi
// The NWDAFContextDocumentApiRouter implementation should parse necessary information from the http request,
// pass the data to a NWDAFContextDocumentApiServicer to perform the required actions, then write the service results to the http response.
type NWDAFContextDocumentApiRouter interface {
	GetNwdafContext(http.ResponseWriter, *http.Request)
}

// NWDAFAnalyticsDocumentApiServicer defines the api actions for the NWDAFAnalyticsDocumentApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type NWDAFAnalyticsDocumentApiServicer interface {
	GetNWDAFAnalytics(context.Context, EventIdAnyOf, EventReportingRequirement, EventFilter, string, TargetUeInformation) (ImplResponse, error)
}

// NWDAFContextDocumentApiServicer defines the api actions for the NWDAFContextDocumentApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type NWDAFContextDocumentApiServicer interface {
	GetNwdafContext(context.Context, ContextIdList, RequestedContext) (ImplResponse, error)
}
