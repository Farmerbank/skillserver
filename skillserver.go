package main

import (
	alexa "github.com/mikeflynn/go-alexa/skillserver"
)

var Applications = map[string]interface{}{
	"/echo/helloworld": alexa.EchoApplication{ // Route
		AppID:    "amzn1.ask.skill.8c0cabc9-c18b-4d53-ac7f-61e2d6f367a2", // Echo App ID from Amazon Dashboard
		OnLaunch: launchIntentHandler,
		OnIntent: echoIntentHandler,
	},
}

func main() {
	alexa.Run(Applications, "3000")
}

func launchIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoResp.OutputSpeech("You just launched the Farmerbank app!")
}

func echoIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoResp.OutputSpeech("Hello world from my new Echo test app!").Card("Hello World", "This is a test card.")
}
