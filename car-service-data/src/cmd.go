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
	"time"
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
	trip{
		name:  "Embarcadero to city hall",
		start: *geo.NewPoint(37.793660, -122.395452),
		end:   *geo.NewPoint(37.779704, -122.418914),
	},
	trip{
		name:  "downtown LA to Hollywood",
		start: *geo.NewPoint(34.048555, -118.245216),
		end:   *geo.NewPoint(34.090904, -118.325662),
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
	driverAggregation()
	fareAggregation()
	t := time.NewTicker(10 * time.Minute)
	for _ = range t.C {
		driverAggregation()
		fareAggregation()
	}
}

func driverAggregation() {
	//Driver data
	iters := len(cities) * len(services)
	doneDriver := make(chan bool, iters)
	driverJobs := createDriverJobs()
	getDriverData(driverJobs, doneDriver)
	awaitCompletion(doneDriver, iters)

}

func fareAggregation() {
	//Fare data
	iters := len(trips) * len(services)
	doneFare := make(chan bool, iters)
	fareJobs := createFareJobs()
	getFareData(fareJobs, doneFare)
	awaitCompletion(doneFare, iters)

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
		for job := range jobs {
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
				log.Printf("StoreDriver(): %s\n", err)
			}
			done <- true
		}
	}()
}

func getFareData(jobs chan fareJob, done chan bool) {
	go func() {
		for job := range jobs {
			log.Printf("Traveling with %s\n", job.service)
			option := collect.CollectOption{
				Start:   job.start,
				End:     job.end,
				Service: job.service,
			}
			data, err := collect.CollectFare(option)
			if err != nil {
				log.Printf("CollectFare(): %s\n", err)
			}
			err = store.StoreFare(data)
			if err != nil {
				log.Printf("StoreFare(): %s\n", err)
			}
			done <- true
		}
	}()
}

func awaitCompletion(done chan bool, workers int) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(done)
}
