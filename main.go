package main

import (
	"emcs-relay-go/api"
	"emcs-relay-go/configs"
	"emcs-relay-go/desktop"
	"emcs-relay-go/udp"
)

func main() {
	udp.Run(configs.UDPListenAddr)
	if configs.EnableDesktop {
		go func() {
			api.Run()
		}()
		desktop.Run()
	} else {
		api.Run()
	}
}
