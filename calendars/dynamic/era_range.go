package dynamic

import ()

type EraRange struct {
	Era          *Era
	RangeCode    string
	StartValue   int
	StartDisplay int
	EndValue     int
	OpenEnded    bool
	Ascending    bool
}
