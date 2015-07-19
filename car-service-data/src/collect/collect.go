// collect.go collects data made available by Chalmers API
// Modeled as a concurrent pipeline.

package collect

import (
	"constants"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kellydunn/golang-geo"
	"github.com/parnurzeal/gorequest"
	"runtime"
	"time"
)

type CollectOption struct {
	Service    constants.ServiceFlag
	Loc        geo.Point
	Address    string
	To, From   string
	Start, End geo.Point
}

type DriverData struct {
	Id       int64
	Service  constants.ServiceFlag
	DriverId string
	Loc      geo.Point
	Time     time.Time
}

type driverSearch struct {
	Data   []driverSearchData `json:Data`
	Error  bool               `json:error`
	Status int                `json:status`
}

type driverSearchData struct {
	App     string               `json:App`
	Brand   string               `json:Brand`
	Drivers []driverSearchDriver `json:Drivers`
}

type driverSearchDriver struct {
	Id        string  `json:id`
	Latitude  float64 `json:float64`
	Longitude float64 `json:float64`
}

type FareData struct {
	Id         int64
	Service    constants.ServiceFlag
	Start, End geo.Point
	Time       time.Time
	Cost       float64
}

func CollectDriver(collect CollectOption) (data []DriverData, err error) {
	if collect.Address == "" && collect.Loc == *geo.NewPoint(0, 0) {
		return nil, errors.New("CollectDriver: Missing location data")
	}
	responses := requestDriver(collect)
	data, err = parseDriver(responses, collect.Service)
	return data, nil
}

func CollectFare(collect CollectOption) (data []FareData, err error) {
	if (collect.To == "" || collect.From == "") &&
		(collect.Start == *geo.NewPoint(0, 0) ||
			collect.End == *geo.NewPoint(0, 0)) {
		return nil, errors.New("CollectDriver: Missing location data")
	}
	responses := requestFare(collect)
	data = parseFare(responses)
	return data, nil
}

func requestDriver(collect CollectOption) <-chan []byte {
	maxRoutines := runtime.GOMAXPROCS(runtime.NumCPU())
	responses := make(chan []byte, maxRoutines)
	go func() {

		request := gorequest.New()
		location := collect.Address
		if collect.Loc != *geo.NewPoint(0, 0) {
			location = fmt.Sprintf("%f,%f", collect.Loc.Lat(), collect.Loc.Lng())
		}
		url := fmt.Sprintf("http://api.rydrapp.co/v1/Ride/Drivers?pickup_location=%s&app=%s", location, collect.Service)
		fmt.Println(url)
		resp, body, errs := request.Get(url).End()
		if resp.StatusCode != 200 || len(errs) > 0 {
			responses <- []byte(fmt.Sprint(resp.StatusCode, errs))
		}
		responses <- []byte(body)
		close(responses)
	}()
	return responses
}

func requestFare(collect CollectOption) <-chan []byte {
	maxRoutines := runtime.GOMAXPROCS(runtime.NumCPU())
	responses := make(chan []byte, maxRoutines)
	return responses
}

func parseDriver(responses <-chan []byte, service constants.ServiceFlag) (data []DriverData, err error) {
	data = make([]DriverData, cap(responses))
	for b := range responses {
		searchData, err := driverDataDecode(b)
		if err != nil {
			return nil, err
		}
		for i, d := range searchData.Data[0].Drivers {
			now := time.Now()
			driver := DriverData{
				Id:       now.Unix() + int64(i),
				Service:  service,
				DriverId: d.Id,
				Loc:      *geo.NewPoint(d.Latitude, d.Longitude),
				Time:     now,
			}
			data = append(data, driver)
		}
	}
	return data, nil
}

func parseFare(responses <-chan []byte) []FareData {
	data := make([]FareData, cap(responses))
	for r := range responses {
		if r != nil {
			data = append(data, FareData{})
		}
	}
	return data
}

func driverDataDecode(b []byte) (driverSearch, error) {
	var driverData driverSearch
	err := json.Unmarshal(b, &driverData)
	if err != nil {
		return driverSearch{}, err
	}
	return driverData, nil
}
