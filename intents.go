package main

import (
	"strconv"

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
