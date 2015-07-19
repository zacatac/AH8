// collect.go collects data made available by Chalmers API
// Modeled as a concurrent pipeline.

package main

import (
	"collect"
	"constants"
	"flag"
	"fmt"
	"github.com/kellydunn/golang-geo"
	"log"
	"os"
	"path/filepath"
	"store"
)

var (
	help    bool
	service string
	Usage   func()
)

func init() {
	Usage = func() { // Change the default flag message
		fmt.Fprintf(os.Stderr, "Usage of %s -sv [service]", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	flag.BoolVar(&help, "help", false, "print usage information")
	flag.StringVar(&service, "sv", "all", "specify a service [uber|lyft|sidecar]")
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	option := collect.CollectOption{
		Loc:     *geo.NewPoint(34.1463926, -118.25027349999999),
		Service: constants.Lyft,
	}
	log.Println("Collecting")
	data, _ := collect.CollectDriver(option)
	for _, d := range data {
		fmt.Println(d)
	}
	log.Println("Storing")
	err := store.StoreDriver(data)
	if err != nil {
		fmt.Errorf("store.Store(): %s", err)
	}
}
