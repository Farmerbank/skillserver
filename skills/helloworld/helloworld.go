package helloworld

import (
	"fmt"
	"strconv"

	alexa "github.com/mikeflynn/go-alexa/skillserver"

	"github.com/naipath/bwfclient"
)

//HWGetApp returns application for helloworld
func HWGetApp() alexa.EchoApplication {
	return alexa.EchoApplication{
		AppID:    "amzn1.ask.skill.8c0cabc9-c18b-4d53-ac7f-61e2d6f367a2",
		OnLaunch: launchIntentHandler,
		OnIntent: echoIntentHandler,
	}
}

var (
	bwfClient = bwfclient.New()
)

func launchIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoResp.OutputSpeech("You just launched the Farmerbank app!")
}

func echoIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	switch echoReq.GetIntentName() {
	case "GetBalance":
		echoResp.OutputSpeech("You can loan " + strconv.Itoa(retrieveKoopsomBedr())).EndSession(false)

	case "ElevatorPitch":
		echoResp.OutputSpeech("Starting the elevatorpitch. With this application it is possible to interact with your bank account in various ways using a natural dialogue.").EndSession(false)

	case "AMAZON.CancelIntent":
		echoResp.OutputSpeech("Closing the farmerbank application.").EndSession(true)

	case "AMAZON.StopIntent":
		echoResp.OutputSpeech("Closing the farmerbank application.").EndSession(true)

	default:
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
