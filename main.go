package main

import (
	"flag"
	"time"
	"webstart/server"
)

var port int
var templateDirName string
var wait time.Duration

func init() {
	flag.IntVar(&port, "p", 8081, "Port number on which server should run")
	flag.StringVar(&templateDirName, "t", "templates", "Name of the template directory Names")
	flag.DurationVar(&wait, "timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
}

func main() {
	flag.Parse()
	s := server.Server{
		Port:         port,
		WaitDuration: wait,
		TemplateDir:  templateDirName,
	}
	s.Start()
}
