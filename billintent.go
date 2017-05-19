package main

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	alexa "github.com/mikeflynn/go-alexa/skillserver"
)

type Bill struct {
	Amount      string    `json:"Amount"`
	Beneficiary string    `json:"Beneficiary"`
	Due         time.Time `json:"Date"`
}

type Bills []Bill

func getAllOutstandingBills() Bills {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, _ := client.Get("https://farmerbank.nl/transactions/Bills")

	defer resp.Body.Close()

	var data Bills
	json.NewDecoder(resp.Body).Decode(&data)
	return data
}

type ListBills struct {
}

func (r ListBills) name() string {
	return "ListBills"
}

func (r ListBills) handle(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	beneficiary, _ := echoReq.GetSlotValue("beneficiary")

	bills := getAllOutstandingBills()

	for _, v := range bills {
		if strings.HasPrefix(v.Beneficiary, beneficiary) {
			echoResp.OutputSpeech("You have one outstanding bill from " + v.Beneficiary + "in the amount of" + v.Amount).EndSession(false)
		}
	}

}
