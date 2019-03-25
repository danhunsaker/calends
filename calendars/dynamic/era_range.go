package dynamic

import ()

type EraRange struct {
	Era          *Era   `json:"-"`
	RangeCode    string `json:"range_code"`
	StartValue   int    `json:"start_value"`
	StartDisplay int    `json:"start_display"`
	EndValue     int    `json:"end_value"`
	OpenEnded    bool   `json:"open_ended"`
	Ascending    bool   `json:"ascending"`
}
