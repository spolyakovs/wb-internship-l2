package main

import (
	"flag"
	"log"
	"strconv"
	"time"
)

var (
	timeout time.Duration
)

func init() {
	flag.CommandLine.DurationVar(&timeout, "timeout", time.Duration(10)*time.Second, "timeout duration; default: 10s")
}

func main() {
	flag.Parse()
	var host string
	var port int

	switch len(flag.Args()) {
	case 0:
		log.Fatalln("Enter host to connect telnet to")
	case 1:
		host = flag.Arg(0)
		port = 80
	case 2:
		host = flag.Arg(0)
		var err error
		port, err = strconv.Atoi(flag.Arg(1))
		if err != nil {
			log.Fatalln("Wrong port")
		}
	default:
		log.Fatalln("Too many arguments")
	}

	client := NewTelnet(host, port, timeout)
	if err := client.Start(); err != nil {
		log.Fatalln(err)
	}
}
