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
		"/echo/helloworld": alexa.EchoApplication{
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
