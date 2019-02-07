package dynamic

import (
	"math/big"
)

type Unit struct {
	Calendar     *Calendar
	InternalName string
	Names        []*UnitName
	Lengths      []*UnitLength
	Eras         []*Era
	Formats      []*Fragment
	ScaleTo      *Unit
	ScaleAmount  int
	ScaleInverse bool // true: ScaleAmount Units per ScaleTo (false: ScaleTos per Unit)
	UsesZero     bool
	UnixEpoch    int
	IsAuxiliary  bool
}

func (u *Unit) toSeconds(units map[string]int) (seconds *big.Float) {
	return
}

func (u *Unit) fromSeconds(seconds *big.Float) (units map[string]int) {
	return
}

func (u *Unit) carryOver(units map[string]int) (out map[string]int) {
	return
}
