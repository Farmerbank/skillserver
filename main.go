package main

import (
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
	echoResp.OutputSpeech("Welcome to the Farmerbank app!").EndSession(false)
}

func echoIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	s := make([]Intent, 10)
	s[0] = ElevatorPitch{}
	s[1] = MaximumMortgage{}
	s[2] = HouseEstimation{}
	s[3] = YesOrNo{}
	s[4] = CancelIntent{}
	s[5] = StopIntent{}
	s[6] = MadeBy{}
	s[7] = ListTransactions{}
	s[8] = SalaryDeposit{}
	s[9] = ListBills{}

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
