package dynamic

import ()

type UnitLength struct {
	Unit        *Unit `json:"-"`
	UnitValue   int   `json:"unit_value"`
	ScaleAmount int   `json:"scale_amount"`
}
