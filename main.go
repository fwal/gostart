// +build !appengine
package main

import (
	"fmt"
	"github.com/theDukeW/docopt-go"
	"net/http"
)

func main() {
	usage := `
	My web server.

	Usage:
	  gostart [--ip=<ip>] [--port=<port>]
	  gostart -h | --help
	  gostart --version

	Options:
	  -h --help     Show this screen.
	  --version     Show version.`

	args, _ := docopt.Parse(usage, nil, true, "gostart 1.0", false)

	ip := "localhost"
	if args["--ip"] != nil {
		ip = args["--ip"].(string)
	}

	port := "3000"
	if args["--port"] != nil {
		port = args["--port"].(string)
	}

	fmt.Printf("Server running at %s:%s\n", ip, port)

	h := httpHandler()

	if err := http.ListenAndServe(ip+":"+port, h); err != nil {
		fmt.Printf("Server error:%v \n", err)
		panic(err)
	}
}
