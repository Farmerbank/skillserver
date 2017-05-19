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
	echoResp.OutputSpeech("Hello world from my new Echo test app! " + strconv.Itoa(retrieveKoopsomBedr()))
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
