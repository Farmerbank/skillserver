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
	echoResp.OutputSpeechSSML(
		`
		<speak>
			<p>Starting the elevator pitch.</p> 
			<p>With our newly build application <emphasis level="moderate">farmer bank</emphasis>, 
			it is possible to interact with your bank account in various ways using a natural dialogue.</p>
			<p><emphasis level="moderate">Farmer bank</emphasis> will be your new personal financial assistant who can take care of all your financial needs and questions.</p>
			<p>Our app can give you house estimations, mortgage capabilities, account balance information and much more.
			Customer interaction has never been this easy.</p>
			<p>All you have to do is ask questions.</p>
		</speak>
	`).EndSession(false)
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
	echoResp.OutputSpeech("Thanks for using farmer bank.").EndSession(true)
}

type StopIntent struct {
}

func (r StopIntent) name() string {
	return "AMAZON.StopIntent"
}
func (r StopIntent) handle(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoResp.OutputSpeech("CThanks for using farmer bank.").EndSession(true)
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
