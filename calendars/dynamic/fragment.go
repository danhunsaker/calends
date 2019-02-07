package dynamic

import ()

type Fragment struct {
	Calendar    *Calendar
	Code        string
	String      string
	Description string
	Type        string // era or unit
	Era         *Era
	Unit        *Unit
	Texts       []*FragmentText
}
