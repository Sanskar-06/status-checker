package main

import (
	"github.com/Sanskar-06/status-checker/constants"
	"github.com/Sanskar-06/status-checker/server"
	"github.com/Sanskar-06/status-checker/status"
)

func main() {
	go status.PollUpdateAllSites(constants.POLLING_RATE)
	server.ServeRequests()
}
