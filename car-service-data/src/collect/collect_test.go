package collect

import (
	"fmt"
	"log"
	"testing"
)

func TestDriverDataDecode(t *testing.T) {
	log.SetFlags(0)
	log.Println("TEST collect")
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

func TestFareDataDecode(t *testing.T) {
	log.SetFlags(0)
	data := []byte(`{"Data":[{"App":"Uber","Brand":"#1a1a1a","Options":[{"price":23,"service":"X","eta":5},{"price":34,"service":"XL","eta":6},{"price":39,"service":"BLACK","eta":9},{"price":53,"service":"SUV","eta":9}]}],"error":false,"status":200}`)
	decoded, err := fareDataDecode(data)
	fmt.Println(decoded)
	if err != nil {
		t.Fatal(err)
	}
	if decoded.Data[0].App != "Uber" {
		t.Fatal("fareDataDecode(): invalid parsing of data")
	}
	if decoded.Data[0].Options[0].Price != 23 {
		t.Fatal("fareDataDecode(): invalid parsing of data")
	}
}

func TestParseFare(t *testing.T) {

}
