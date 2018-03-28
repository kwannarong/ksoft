package main

import (
	"encoding/json"
	"fmt"
)

type BuyOfferAcceptRequest struct {
	Context               Context               `json:"CONTEXT_DATA"`
	CustomerIdentifier    CustomerIdentifier    `json:"CUSTOMER_IDENTIFIER"`
	FullFillMentDetail    FullFillMentDetail    `json:"FULLFILLMENT_DETAILS"`
	Intent                string                `json:"INTENT"`
	RecordType            string                `json:"RECORD_TYPE"`
	RequesterApplication  string                `json:"REQUESTER_APPLICATION"`
	RequesterChannel      string                `json:"REQUESTER_CHANNEL"`
	RequesterLocationData RequesterLocationData `json:"REQUESTER_LOCATION"`
	RequesterZone         string                `json:"REQUESTER_ZONE"`
	RequesterId           string                `json:"REQUEST_ID"`
}

type CustomerIdentifier struct {
	Msisdn string `json:"MSISDN"`
}

type Context struct {
	AccountType       string `json:"ACCOUNT_TYPE,omitempty"`
	NewCustomerNumber int    `json:"NEW_CUSTNUM_POS,omitempty"`
	UserCode          string `json:"USER_CODE,omitempty"`
	CommanyCode       string `json:"COMPANY_CODE,omitempty"`
	CustomerID        string `json:"CUSTOMER_ID,omitempty"`
}

type FullFillMentDetail struct {
	ErrorCode          string `json:"ERROR_CODE"`
	ErrorDescription   string `json:"ERROR_DESCRIPTION"`
	FullFillMentSystem string `json:"FULLFILLMENT_SYSTEM"`
	PackageStartDate   string `json:"PACKAGE_START_DATE"`
	Status             string `json:"STATUS"`
}

type RequesterLocationData struct {
	Location []Location `json:"LOCATIONS"`
}

type Location struct {
	Id     string   `json:"id"`
	FlowId []string `json:"FLOW_ID"`
}

type BuyOfferAcceptResponse struct {
	Status             string             `json:"STATUS"`
	CustomerIdentifier CustomerIdentifier `json:"CUSTOMER_IDENTIFIER,omitempty"`
	ErrorCode          string             `json:"ERROR_CODE,omitempty"`
	ErrorMessage       string             `json:"ERROR_MESSAGE,omitempty"`
	RequestId          string             `json:"REQUEST_ID,omitempty"`
}

func main() {
	input := BuyOfferAcceptRequest{}
	userCode := "LLCHAILUCK"

	input.Context = Context{AccountType: "PREPAID", NewCustomerNumber: 13234, UserCode: userCode}
	input.CustomerIdentifier = CustomerIdentifier{Msisdn: "66834567890"}
	input.FullFillMentDetail = FullFillMentDetail{ErrorCode: "", ErrorDescription: "", FullFillMentSystem: "CCB", PackageStartDate: "14950004170000", Status: "SUCCESS"}
	input.Intent = "BUY_NOTIFICATION"
	input.RecordType = "HANDLE_CUSTOMER_INTENT"
	input.RequesterApplication = "OMR_P2P"
	input.RequesterChannel = "OMR_PRE"
	input.RequesterLocationData = RequesterLocationData{Location: []Location{{Id: "Recommend_P2P", FlowId: []string{"6694241111829927"}}}}
	input.RequesterZone = "OMR_P2P"
	input.RequesterId = "66834567890-1511942606095"
	var payload []byte
	payload, _ = json.Marshal(input)

	fmt.Println(string(payload))
}
