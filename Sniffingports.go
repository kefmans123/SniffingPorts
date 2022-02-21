package main

import (

	// "net"
	"flag"
	"log"
	"strconv"
	"strings"
)

var address string
var port int64
var portrange string
var portcheck bool

var startport int
var endport int

var method string

func init() {
	flag.StringVar(&address, "d", "", "Fill in any domain that you want to use.")
	flag.StringVar(&address, "ip", "", "Fill in any IP address that you want to use.")
	flag.Int64Var(&port, "p", 0, "Fill in any port that you want to use.")
	flag.StringVar(&portrange, "pr", "", "Fill in any port range that you want to use.")
	flag.BoolVar(&portcheck, "pch", false, "Do you want to check all the common ports? Use true/false")
	flag.Parse()

	if address == "" {
		log.Fatal("I need an address to sniff.")
	}

	if !portcheck {
		if port == 0 && portrange == "" {
			log.Fatal("I need a port or portrange to continue. If you want to check common ports please use the right flag.")
		} else if port != 0 {
			method = "port"
		} else if portrange != "" {
			ports := strings.Split(portrange, "-")

			if len(ports) != 2 {
				log.Fatal("Detected an error in the port range. Please try again! Example = 80-443")
			}

			f_startport, _ := strconv.Atoi(ports[0])
			f_endport, _ := strconv.Atoi(ports[1])

			startport = f_startport
			endport = f_endport

			method = "portrange"

		}
	}

}

func main() {

	switch method {
	case "portrange":

	case "port":
		log.Println("Checking port!")
		log.Println(sniff(address, port))

	case "common":
	}

}

func sniff(address string, port int) {

}
