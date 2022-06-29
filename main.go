package main

import (
	"flag"
	"os"

	"github.com/amandinedrebes/orness/internal/api"
)

func main() {
	//parse commandline
	ip := ""
	port := ""
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flag.StringVar(&ip, "ip", "127.0.0.1", "Specify the ip (by default 127.0.0.1)")
	flag.StringVar(&port, "port", "4000", "Specify the port configuration (by default 4000)")
	flag.Parse()

	//start webserver
	api.InitAPI(ip, port)
}
