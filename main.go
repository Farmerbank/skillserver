package main

import (
	"fmt"
	"strconv"

	"flag"
	"log"

	alexa "github.com/mikeflynn/go-alexa/skillserver"
	"github.com/naipath/bwfclient"
)

var (
	bwfClient    = bwfclient.New()
	applications = map[string]interface{}{
		"/echo/farmerbank": alexa.EchoApplication{
			AppID:    "amzn1.ask.skill.a1c73b55-4e76-45f2-8478-bc79b77cc537",
			OnLaunch: launchIntentHandler,
			OnIntent: echoIntentHandler,
		},
	}
)

func main() {
	var (
		httpPort = flag.String("port", "3000", "HTTP server port")
	)
	flag.Parse()

	log.Println("Starting Farmerbank Skillserver")
	log.Printf("Service listening on %s", *httpPort)

	alexa.Run(applications, *httpPort)
}

func launchIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoResp.OutputSpeech("You just launched the Farmerbank app!    Moooh").EndSession(false)
}

func echoIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	s := make([]Intent, 4)
	s[0] = ElevatorPitch{}
	s[1] = GetBalance{}
	s[2] = CancelIntent{}
	s[3] = StopIntent{}

	handled := false
	for _, element := range s {
		if element.name() == echoReq.GetIntentName() {
			element.handle(echoReq, echoResp)
			handled = true
		}
	}

	if handled == false {
		echoResp.OutputSpeech("Unrecognized command").EndSession(false)
	}
}

func retrieveKoopsomBedr() int {
	resp, err := bwfClient.Request(bwfclient.BwfRequest{
		AanvragerBrutoJaarinkomenBedr: 40000,
		PartnerBrutoJaarinkomenBedr:   0,
	})
	if err != nil {
		fmt.Print(err)
	}
	return resp.MaxTeLenenObvInkomen.Tienjaarsrente.KoopsomBedr
}

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
	echoResp.OutputSpeech("You can loan " + strconv.Itoa(retrieveKoopsomBedr())).EndSession(false)
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
