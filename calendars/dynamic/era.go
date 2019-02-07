package dynamic

import ()

type Era struct {
	Calendar     *Calendar
	InternalName string
	Unit         *Unit
	Ranges       []*EraRange
	DefaultRange *EraRange
	Formats      []*Fragment
}
