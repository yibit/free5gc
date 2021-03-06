/*
 * Npcf_PolicyAuthorization Service API
 *
 * This is the Policy Authorization Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Identifies the modifications to an Individual Application Session Context and may include the modifications to the sub-resource Events Subscription.
type AppSessionContextUpdateData struct {
	// Contains an AF application identifier.
	AfAppId   string                  `json:"afAppId,omitempty" bson:"afAppId"`
	AfRoutReq *AfRoutingRequirementRm `json:"afRoutReq,omitempty" bson:"afRoutReq"`
	// Contains an identity of an application service provider.
	AspId string `json:"aspId,omitempty" bson:"aspId"`
	// string identifying a BDT Reference ID as defined in subclause 5.3.3 of 3GPP TS 29.154.
	BdtRefId      string                      `json:"bdtRefId,omitempty" bson:"bdtRefId"`
	EvSubsc       *EventsSubscReqDataRm       `json:"evSubsc,omitempty" bson:"evSubsc"`
	MedComponents map[string]MediaComponentRm `json:"medComponents,omitempty" bson:"medComponents"`
	// indication of MPS service request
	MpsId string `json:"mpsId,omitempty" bson:"mpsId"`
	// Contains an identity of a sponsor.
	SponId     string           `json:"sponId,omitempty" bson:"sponId"`
	SponStatus SponsoringStatus `json:"sponStatus,omitempty" bson:"sponStatus"`
}
