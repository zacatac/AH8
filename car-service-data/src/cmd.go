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

var services = []constants.ServiceFlag{
	constants.Uber,
	constants.Lyft,
	constants.Sidecar,
	constants.Flywheel,
}

var cities = []constants.MetropolitanArea{
	constants.LA,
	constants.SF,
	constants.NYC,
}

var trips = []trip{
	trip{
		name:  "penn station to wall street",
		start: *geo.NewPoint(40.751247, -73.993819),
		end:   *geo.NewPoint(40.704681, -74.007712),
	},
}

type driverJob struct {
	service constants.ServiceFlag
	city    constants.MetropolitanArea
}

type fareJob struct {
	service    constants.ServiceFlag
	start, end geo.Point
}

type trip struct {
	name       string
	start, end geo.Point
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	iters := len(cities) * len(services)
	done := make(chan bool, iters)
	jobs := createDriverJobs()
	getDriverData(jobs, done)
	// iters := len(trips) * len(services)
	// done := make(chan bool, iters)
	// jobs := createFareJobs()
	// getFareData(jobs, done)
	awaitCompletion(done, iters)
}

func createDriverJobs() (jobs chan driverJob) {
	jobs = make(chan driverJob, len(services)*len(cities))
	go func() {
		for _, service := range services {
			for _, city := range cities {
				jobs <- driverJob{service, city}
			}
		}
		close(jobs)
	}()
	return jobs
}

func createFareJobs() (jobs chan fareJob) {
	jobs = make(chan fareJob, len(trips)*len(services))
	go func() {
		for _, trip := range trips {
			for _, service := range services {
				jobs <- fareJob{service, trip.start, trip.end}
			}
		}
		close(jobs)
	}()
	return jobs
}

func getDriverData(jobs chan driverJob, done chan bool) {
	go func() {
		job := <-jobs
		log.Printf("%s in %s\n", job.service, job.city)
		option := collect.CollectOption{
			Loc:     job.city.Loc,
			Service: job.service,
		}
		data, err := collect.CollectDriver(option)
		if err != nil {
			log.Printf("CollectDriver(): %s\n", err)
		}
		log.Printf("drivers: %d\n", len(data))
		err = store.StoreDriver(data)
		if err != nil {
			log.Printf("StoreDriver(): %s", err)
		}
		done <- true
	}()
}

func getFareData(jobs chan fareJob, done chan bool) {
	go func() {
		job := <-jobs
		log.Printf("Traveling with %s", job.service)
		option := collect.CollectOption{
			Start:   job.start,
			End:     job.end,
			Service: job.service,
		}
		fmt.Println(option)
		data, err := collect.CollectFare(option)
		fmt.Println(data)
		if err != nil {
			log.Printf("CollectFare(): %s\n", err)
		}
		err = store.StoreFare(data)
		if err != nil {
			log.Printf("StoreFare(): %s", err)
		}
		done <- true
	}()

}

func awaitCompletion(done chan bool, workers int) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(done)
}
