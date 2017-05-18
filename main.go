package main

import (
	"fmt"
	"strconv"

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
	alexa.Run(applications, "3000")
}

func launchIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoResp.OutputSpeech("You just launched the Farmerbank app!").EndSession(false)
}

func echoIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoResp.OutputSpeech("Hello world from my new Echo test app! " + strconv.Itoa(retrieveKoopsomBedr())).EndSession(true)
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
