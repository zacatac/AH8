package constants

type ServiceFlag int

const (
	Uber ServiceFlag = iota
	Lyft
	Sidecar
	Flywheel
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
