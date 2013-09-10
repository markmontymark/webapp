package main

import (
	"flag"
	"appserver/server"
	"./handlers"
)

func main () {
	var (
      port = flag.Int("p",8000, "Port to listen on")
      host = flag.String("h","localhost","Hostname to listen on")
      config_file = flag.String("c","./app.yaml","Path to app server config file.")
		handlers = handlers.NewMyHandlers()
   )
	server.Start( *host, *port, *config_file, handlers)
}

