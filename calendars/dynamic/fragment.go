package dynamic

import ()

type Fragment struct {
	Calendar    *Calendar
	Code        rune
	String      string
	Description string
	IsEra       bool
	Era         *Era
	Unit        *Unit
	Texts       []*FragmentText
}

func (self *Fragment) parse(in string) (unit string, value, end int, err error) {
	return
}

func (self *Fragment) format(units map[string]int) (out string) {
	return
}

func (self *Fragment) value(units map[string]int) (value, length int, era string) {
	var unit *Unit

	if self.IsEra {
		unit = self.Era.Unit
	} else {
		unit = self.Unit
	}

	value = units[unit.InternalName]
	if len(unit.Lengths) > 0 {
		length = 0
		for _, l := range unit.Lengths {
			if l.UnitValue == value {
				length = l.ScaleAmount
				break
			}
		}
	} else {
		length = unit.ScaleAmount
	}
	era = ""

	if self.IsEra {
		var myRange *EraRange = self.Era.DefaultRange
		var adjusted int

		for _, r := range self.Era.Ranges {
			if r.Ascending {
				adjusted = value - r.StartDisplay
				if (r.OpenEnded || r.EndValue >= adjusted) && r.StartValue <= adjusted {
					myRange = r
					break
				}
			} else {
				adjusted = r.StartDisplay - value
				if (r.OpenEnded || r.EndValue <= adjusted) && r.StartValue >= adjusted {
					myRange = r
					break
				}
			}
		}

		value = myRange.StartValue + adjusted
		era = myRange.RangeCode
	}

	return
}
