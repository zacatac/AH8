package constants

import (
	"github.com/kellydunn/golang-geo"
)

type ServiceFlag int
type MetropolitanArea struct {
	Name string
	Loc  geo.Point
}

const (
	Uber ServiceFlag = iota
	Lyft
	Sidecar
	Flywheel
)

var (
	LA  MetropolitanArea = MetropolitanArea{"Los Angeles", *geo.NewPoint(34.039834, -118.246349)}
	NYC                  = MetropolitanArea{"New York City", *geo.NewPoint(40.747687, -73.987328)}
	SF                   = MetropolitanArea{"San Francisco", *geo.NewPoint(37.793317, -122.400607)}
)

func (flag ServiceFlag) String() string {
	switch flag {
	case Uber:
		return "Uber"
	case Lyft:
		return "Lyft"
	case Sidecar:
		return "Sidecar"
	case Flywheel:
		return "Flywheel"
	}
	return "?"
}

func (city MetropolitanArea) String() string {
	return city.Name
}
