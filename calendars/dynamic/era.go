package dynamic

import (
	"strconv"
)

type Era struct {
	Calendar     *Calendar
	InternalName string
	Unit         *Unit
	Ranges       []*EraRange
	DefaultRange *EraRange
	Formats      []*Fragment
}

func (self *Era) unitValue(parsed map[string]string) (name string, value int) {
	var raw int = 0
	if _, ok := parsed["value"]; ok {
		tmpValue, err := strconv.ParseInt(parsed["value"], 0, 0)
		if err == nil {
			raw = int(tmpValue)
		}
	}

	var myRange *EraRange = self.DefaultRange
	var adjusted int
	if code, ok := parsed["code"]; ok {
		for _, r := range self.Ranges {
			if r.RangeCode != code {
				continue
			}

			if r.Ascending {
				adjusted = raw - r.StartDisplay
				if (r.OpenEnded || r.EndValue >= adjusted) && r.StartValue <= adjusted {
					myRange = r
					break
				}
			} else {
				adjusted = r.StartDisplay - raw
				if (r.OpenEnded || r.EndValue <= adjusted) && r.StartValue >= adjusted {
					myRange = r
					break
				}
			}
		}
	}

	name = self.Unit.InternalName
	value = myRange.StartValue + adjusted

	return
}
