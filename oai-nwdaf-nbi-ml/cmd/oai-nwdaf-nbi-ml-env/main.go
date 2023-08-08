/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 1.0.0-alpha.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	nbi_ml "gitlab.eurecom.fr/development/oai-nwdaf/components/oai-nwdaf-nbi-ml/pkg/nbiml"
)

func main() {
	log.Printf("Server started")

	// load the environment variables from the file .env (no-docker scenario)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	IndividualNWDAFMLModelProvisionSubscriptionDocumentApiService := nbi_ml.NewIndividualNWDAFMLModelProvisionSubscriptionDocumentApiService()
	IndividualNWDAFMLModelProvisionSubscriptionDocumentApiController := nbi_ml.NewIndividualNWDAFMLModelProvisionSubscriptionDocumentApiController(IndividualNWDAFMLModelProvisionSubscriptionDocumentApiService)

	SubscriptionsCollectionApiService := nbi_ml.NewSubscriptionsCollectionApiService()
	SubscriptionsCollectionApiController := nbi_ml.NewSubscriptionsCollectionApiController(SubscriptionsCollectionApiService)

	router := nbi_ml.NewRouter(IndividualNWDAFMLModelProvisionSubscriptionDocumentApiController, SubscriptionsCollectionApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
