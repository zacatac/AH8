package collect

import (
	"log"
	"testing"
)

func TestDriverDataDecode(t *testing.T) {
	log.SetFlags(0)
	log.Println("TEST bigdigits")
	data := []byte(`{"Data":[{"App":"Flywheel","Brand":"#e91219","Drivers":[{"id":"140886","latitude":34.1052515,"longitude":-118.2880089},{"id":"140921","latitude":34.0808819,"longitude":-118.2721651}]}],"error":false,"status":200}`)
	decoded, err := driverDataDecode(data)
	if err != nil {
		t.Fatal(err)
	}
	if decoded.Data[0].App != "Flywheel" {
		t.Fatal("driverDataDecode(): invalid parsing of data")
	}
	if decoded.Data[0].Drivers[0].Id != "140886" {
		t.Fatal("driverDataDecode(): invalid parsing of data")
	}
}

func TestParseDriver(t *testing.T) {

}
