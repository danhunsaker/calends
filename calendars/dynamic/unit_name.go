package dynamic

import ()

type UnitName struct {
	Unit        *Unit  `json:"-"`
	UnitName    string `json:"unit_name"`
	NameContext string `json:"name_context"`
}
