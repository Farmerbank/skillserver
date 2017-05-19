package main

import (
	"github.com/farmerbank/skillserver/skills/helloworld"
	alexa "github.com/mikeflynn/go-alexa/skillserver"
)

func getApps() map[string]interface{} {
	a := make(map[string]interface{})

	a["/echo/helloworld"] = helloworld.HWGetApp()

	return a
}

func main() {

	alexa.Run(getApps(), "3000")

}
