// store stores car service data to cluster point

package store

import (
	"collect"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

func StoreDriver(data []collect.DriverData) error {
	return writeDriver(data)
}

func StoreFare(data []collect.FareData) error {
	return writeFare(data)
}

func writeDriver(data []collect.DriverData) error {
	for _, driver := range data {
		request := gorequest.New()
		driverLiteral := `{"id":%d,` +
			`"service":"%s",` +
			`"driverid":"%s",` +
			`"lat":%f,` +
			`"lon":%f,` +
			`"time":"%s"}`
		body := fmt.Sprintf(driverLiteral, driver.Id, driver.Service,
			driver.DriverId, driver.Loc.Lat(),
			driver.Loc.Lng(), driver.Time)
		resp, _, errs := request.Post("https://api-us.clusterpoint.com/100882/car-service-test.json").
			SetBasicAuth("field.zackery@gmail.com", "angelhack").
			Send(body).
			End()
		if len(errs) > 0 {
			fmt.Println(resp)
			fmt.Println(errs)
		}
	}
	return nil
}

func writeFare(data []collect.FareData) error {
	for _, option := range data {
		fmt.Println(option)
		request := gorequest.New()
		fareLiteral := `{"id":%d,` +
			`"service":"%s",` +
			`"type":"%s",` +
			`"cost":"%d",` +
			`"eta":"%d",` +
			`"start_lat":%f,` +
			`"start_lon":%f,` +
			`"end_lat":%f,` +
			`"end_lon":%f,` +
			`"time":"%s"}`
		body := fmt.Sprintf(fareLiteral, option.Id, option.Service, option.Type,
			option.Cost, option.Eta,
			option.Start.Lat(), option.Start.Lng(),
			option.End.Lat(), option.End.Lng(), option.Time)
		resp, _, errs := request.Post("https://api-us.clusterpoint.com/100882/car-service-test.json").
			SetBasicAuth("field.zackery@gmail.com", "angelhack").
			Send(body).
			End()
		if len(errs) > 0 {
			fmt.Println(resp)
			fmt.Println(errs)
		}
	}
	return nil
}
