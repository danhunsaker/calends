package dynamic

import (
	"encoding/json"
	"strconv"
)

type Era struct {
	Calendar     *Calendar   `json:"-"`
	InternalName string      `json:"internal_name"`
	Unit         *Unit       `json:"-"`
	Ranges       []*EraRange `json:"ranges"`
	DefaultRange *EraRange   `json:"default_range,name"`
	Formats      []*Fragment `json:"-"`
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

func (self *Era) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"internal_name": self.InternalName,
		"ranges":        self.Ranges,
		"default_range": self.DefaultRange.RangeCode,
	})
}
