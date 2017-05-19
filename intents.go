package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"fmt"

	alexa "github.com/mikeflynn/go-alexa/skillserver"
)

type Intent interface {
	name() string
	handle(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse)
}

type ElevatorPitch struct {
}

func (r ElevatorPitch) name() string {
	return "ElevatorPitch"
}
func (r ElevatorPitch) handle(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoResp.OutputSpeech("Starting the elevatorpitch. With this application it is possible to interact with your bank account in various ways using a natural dialogue.").EndSession(false)
}

type GetBalance struct {
}

func (r GetBalance) name() string {
	return "GetBalance"
}
func (r GetBalance) handle(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoResp.OutputSpeech("You can loan " + strconv.Itoa(retrieveKoopsomBedr()) + "euros").EndSession(false)
}

type CancelIntent struct {
}

func (r CancelIntent) name() string {
	return "AMAZON.CancelIntent"
}
func (r CancelIntent) handle(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoResp.OutputSpeech("Closing the farmerbank application.").EndSession(true)
}

type StopIntent struct {
}

func (r StopIntent) name() string {
	return "AMAZON.StopIntent"
}
func (r StopIntent) handle(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoResp.OutputSpeech("Closing the farmerbank application.").EndSession(true)
}

type HouseEstimation struct {
}

func (r HouseEstimation) name() string {
	return "HouseEstimation"
}

func (r HouseEstimation) handle(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	fmt.Print(echoReq.AllSlots())

	number, _ := echoReq.GetSlotValue("number")
	numberConvert, _ := strconv.Atoi(number)

	secondNumber, _ := echoReq.GetSlotValue("secondNumber")
	secondNumberConvert, _ := strconv.Atoi(secondNumber)

	echoResp.OutputSpeech("The value is " + strconv.Itoa(numberConvert+secondNumberConvert)).EndSession(false)
}

type YesOrNoResponse struct {
	Answer string `json:"answer"`
	Forced bool   `json:"forced"`
	Image  string `json:"image"`
}

type YesOrNo struct {
}

func (r YesOrNo) name() string {
	return "YesOrNo"
}

func (r YesOrNo) handle(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	resp, _ := http.Get("https://yesno.wtf/api")
	defer resp.Body.Close()

	var data YesOrNoResponse
	json.NewDecoder(resp.Body).Decode(&data)

	echoResp.OutputSpeech("The dice has been rolled .... the answer is: " + data.Answer).EndSession(false)
}
