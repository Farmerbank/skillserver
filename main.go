package main

import (
	"github.com/farmerbank/skillserver/skills/helloworld"
	alexa "github.com/mikeflynn/go-alexa/skillserver"
	"flag"
	"log"
)

func getApps() map[string]interface{} {
	a := make(map[string]interface{})

	a["/echo/helloworld"] = helloworld.HWGetApp()

	return a
}

func main() {

	var (
		httpPort = flag.String("port", "3000", "HTTP server port")
	)
	flag.Parse()

	log.Println("Starting Farmerbank Skillserver")
	log.Printf("Service listening on %s", *httpPort)

	alexa.Run(getApps(), *httpPort)

}
