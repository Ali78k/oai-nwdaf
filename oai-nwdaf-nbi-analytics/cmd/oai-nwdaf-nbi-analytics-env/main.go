/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.2.0-alpha.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"gitlab.eurecom.fr/mekrache/oai-nwdaf/components/oai-nwdaf-nbi-analytics/pkg/analytics"
)

func main() {

	log.Printf("Server started")

	// load the environment variables from the file .env (no-docker scenario)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	NWDAFAnalyticsDocumentApiService := analytics.NewNWDAFAnalyticsDocumentApiService()
	NWDAFAnalyticsDocumentApiController := analytics.NewNWDAFAnalyticsDocumentApiController(NWDAFAnalyticsDocumentApiService)

	NWDAFContextDocumentApiService := analytics.NewNWDAFContextDocumentApiService()
	NWDAFContextDocumentApiController := analytics.NewNWDAFContextDocumentApiController(NWDAFContextDocumentApiService)

	router := analytics.NewRouter(NWDAFAnalyticsDocumentApiController, NWDAFContextDocumentApiController)

	log.Fatal(http.ListenAndServe(":8880", router))
}
