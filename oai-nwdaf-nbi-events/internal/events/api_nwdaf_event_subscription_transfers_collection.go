/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.7
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package events

import (
	"encoding/json"
	"net/http"
	"strings"
)

// NWDAFEventSubscriptionTransfersCollectionApiController binds http requests to an api service and writes the service results to the http response
type NWDAFEventSubscriptionTransfersCollectionApiController struct {
	service      NWDAFEventSubscriptionTransfersCollectionApiServicer
	errorHandler ErrorHandler
}

// NWDAFEventSubscriptionTransfersCollectionApiOption for how the controller is set up.
type NWDAFEventSubscriptionTransfersCollectionApiOption func(*NWDAFEventSubscriptionTransfersCollectionApiController)

// WithNWDAFEventSubscriptionTransfersCollectionApiErrorHandler inject ErrorHandler into controller
func WithNWDAFEventSubscriptionTransfersCollectionApiErrorHandler(h ErrorHandler) NWDAFEventSubscriptionTransfersCollectionApiOption {
	return func(c *NWDAFEventSubscriptionTransfersCollectionApiController) {
		c.errorHandler = h
	}
}

// NewNWDAFEventSubscriptionTransfersCollectionApiController creates a default api controller
func NewNWDAFEventSubscriptionTransfersCollectionApiController(s NWDAFEventSubscriptionTransfersCollectionApiServicer, opts ...NWDAFEventSubscriptionTransfersCollectionApiOption) Router {
	controller := &NWDAFEventSubscriptionTransfersCollectionApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the NWDAFEventSubscriptionTransfersCollectionApiController
func (c *NWDAFEventSubscriptionTransfersCollectionApiController) Routes() Routes {
	return Routes{
		{
			"CreateNWDAFEventSubscriptionTransfer",
			strings.ToUpper("Post"),
			"/nnwdaf-eventssubscription/v1/transfers",
			c.CreateNWDAFEventSubscriptionTransfer,
		},
	}
}

// CreateNWDAFEventSubscriptionTransfer - Provide information about requested analytics subscriptions transfer and potentially create a new Individual NWDAF Event Subscription Transfer resource.
func (c *NWDAFEventSubscriptionTransfersCollectionApiController) CreateNWDAFEventSubscriptionTransfer(w http.ResponseWriter, r *http.Request) {
	analyticsSubscriptionsTransferParam := AnalyticsSubscriptionsTransfer{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&analyticsSubscriptionsTransferParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertAnalyticsSubscriptionsTransferRequired(analyticsSubscriptionsTransferParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateNWDAFEventSubscriptionTransfer(r.Context(), analyticsSubscriptionsTransferParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}