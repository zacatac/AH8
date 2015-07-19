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

func writeDriver(data []collect.DriverData) error {
	for _, driver := range data {
		request := gorequest.New()
		driverLiteral := `{"id":%d,` +
			`"service":"%s",` +
			`"driverid":"%s",` +
			`"lat":%f,` +
			`"lon":%f,` +
			`"time":"%s"}`
		body := fmt.Sprintf(driverLiteral, driver.Id, driver.Service, driver.DriverId, driver.Loc.Lat(), driver.Loc.Lng(), driver.Time)
		fmt.Println(body)
		resp, _, errs := request.Post("https://api-us.clusterpoint.com/100882/car-service-test.json").
			SetBasicAuth("field.zackery@gmail.com", "angelhack").
			Send(body).
			End()
		fmt.Println(resp)
		fmt.Println(errs)
	}
	return nil
}
