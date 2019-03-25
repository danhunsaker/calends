package dynamic

import (
	"encoding/json"
	"math/big"
)

type Unit struct {
	Calendar     *Calendar     `json:"-"`
	InternalName string        `json:"internal_name"`
	Names        []*UnitName   `json:"names"`
	Lengths      []*UnitLength `json:"lengths"`
	Eras         []*Era        `json:"eras"`
	Formats      []*Fragment   `json:"-"`
	ScalesToMe   []*Unit       `json:"-"`
	ScaleTo      *Unit         `json:"scale_to,name"`
	ScaleAmount  int           `json:"scale_amount"`
	ScaleInverse bool          `json:"scale_inverse"` // true if 1 ScaleTo is ScaleAmount Units
	UsesZero     bool          `json:"uses_zero"`
	UnixEpoch    int           `json:"unix_epoch"`
	IsAuxiliary  bool          `json:"is_auxiliary"`
}

func (self *Unit) toSeconds(units map[string]int) (seconds *big.Float) {
	seconds = big.NewFloat(0.0)
	value, _ := self.consolidate(units)
	seconds.SetString(value)

	return
}

func (self *Unit) consolidate(in map[string]int) (name string, value int) {
	name = "0"
	value = 0

	if len(in) > 0 {
		for _, unit := range self.ScalesToMe {
			if !unit.IsAuxiliary {
				unitName, unitValue := unit.consolidate(in)
				delete(in, unit.InternalName)
				in[unitName] = unitValue
			}
		}

		if myVal, ok := in[self.InternalName]; ok {
			if !self.UsesZero {
				myVal--
			}
			scaled := self.scale(big.NewFloat(float64(myVal)))
			name = scaled.Text('g', 45)
			value = 0

			if self.ScaleTo != nil {
				if unitValue, ok := in[self.ScaleTo.InternalName]; ok {
					tmp, _ := scaled.Add(scaled, big.NewFloat(float64(unitValue))).Int64()
					name = self.ScaleTo.InternalName
					value = int(tmp)
				}
			}
		}
	}

	return
}

func (self *Unit) fromSeconds(seconds *big.Float) (units map[string]int) {
	units = make(map[string]int, 0)
	var unitVal int64
	var round big.Accuracy
	if self.ScaleInverse {
		unitVal, round = seconds.Quo(seconds, big.NewFloat(float64(self.ScaleAmount))).Int64()
	} else {
		unitVal, round = seconds.Mul(seconds, big.NewFloat(float64(self.ScaleAmount))).Int64()
	}

	units[self.InternalName] = int(unitVal)

	if round.String() != "Exact" {
		for _, unit := range self.ScalesToMe {
			if unit.IsAuxiliary || !unit.ScaleInverse {
				continue
			}

			tmpUnits := unit.fromSeconds(big.NewFloat(0.0).Mul(
				big.NewFloat(0.0).Sub(seconds, big.NewFloat(float64(unitVal))),
				big.NewFloat(float64(unit.ScaleAmount)),
			))

			for name, value := range tmpUnits {
				if _, exists := units[name]; exists {
					value += units[name]
				}
				units[name] = value
			}
		}
	}

	return
}

func (self *Unit) carryOver(units map[string]int) (out map[string]int) {
	out = units

	var myZero int = 0
	if !self.UsesZero {
		myZero = 1
	}

	var myVal int = 0
	if !self.IsAuxiliary {
		myVal, ok := units[self.InternalName]
		if !ok {
			myVal = myZero
		}
		myVal -= myZero
	}

	var inverse = func(myVal, scale int) int {
		return (myVal * scale) % scale
	}
	var normal = func(myVal, scale int) int {
		return (myVal - (myVal % scale)) / scale
	}

	for _, unit := range self.ScalesToMe {
		var unitZero int = 0
		if !unit.UsesZero {
			unitZero = 1
		}

		unitVal, ok := units[unit.InternalName]
		if !ok {
			unitVal = unitZero
		}
		unitVal -= unitZero

		unitAdjustment := 0
		myAdjustment := 0

		var compare int
		var unitExpression func(int, int) int
		var myExpression func(int, int) int
		if unit.ScaleInverse {
			compare = 1
			unitExpression = inverse
			myExpression = func(adjust, scale int) int {
				return adjust / scale
			}
		} else {
			compare = myVal
			unitExpression = normal
			myExpression = func(adjust, scale int) int {
				return adjust * scale
			}
		}

		if unit.ScaleAmount > 0 {
			unitAdjustment = unitExpression(myVal, unit.ScaleAmount)
			myAdjustment = myExpression(unitAdjustment, unit.ScaleAmount)
		} else {
			lengthCount := len(unit.Lengths)

			var lengthSum = 0
			for _, length := range unit.Lengths {
				lengthSum += length.ScaleAmount
			}

			if lengthCount > 0 {
				unitAdjustment += unitExpression(myVal, lengthSum) * lengthCount
				myAdjustment += myExpression(unitAdjustment/lengthCount, lengthSum)

				var lengthNum int
				for lengthNum = 0; compare > myAdjustment && myVal >= unit.Lengths[lengthNum].ScaleAmount; lengthNum++ {
					unitAdjustment += 1
					myAdjustment += unit.Lengths[lengthNum%lengthCount].ScaleAmount
				}

				if compare < myAdjustment {
					unitAdjustment -= 1
					myAdjustment -= unit.Lengths[(lengthNum-1)%lengthCount].ScaleAmount
				}
			}
		}

		out[unit.InternalName] = unitVal + unitAdjustment + unitZero

		if !unit.IsAuxiliary {
			myVal -= myAdjustment - myZero
			out[self.InternalName] = myVal
		}

		out = unit.carryOver(out)
	}

	return
}

func (self *Unit) reduceAuxiliary(in int) (name string, out int) {
	name = self.InternalName
	out = in

	if self.IsAuxiliary {
		scaled, _ := self.scale(big.NewFloat(float64(in))).Int64()
		return self.ScaleTo.reduceAuxiliary(int(scaled))
	}

	return
}

func (self *Unit) scale(in *big.Float) (out *big.Float) {
	out = big.NewFloat(0.0)
	if self.ScaleAmount != 0 {
		if self.ScaleInverse {
			out.Quo(in, big.NewFloat(float64(self.ScaleAmount)))
		} else {
			out.Mul(in, big.NewFloat(float64(self.ScaleAmount)))
		}
	} else if len(self.Lengths) < 1 {
		out.SetInt64(0)
	} else {
		var lengthCount = len(self.Lengths)

		var lengthSum = 0
		for _, length := range self.Lengths {
			lengthSum += length.ScaleAmount
		}

		adjustmentLoops, _ := big.NewFloat(0.0).Quo(in, big.NewFloat(float64(lengthCount))).Int64()
		adjustmentUnits := big.NewFloat(0.0).Add(
			big.NewFloat(0.0).Sub(in, big.NewFloat(float64(int(adjustmentLoops)*lengthCount))),
			big.NewFloat(float64(self.UnixEpoch)),
		)
		adjUnitsInt, _ := adjustmentUnits.Int64()
		adjustmentRemains := big.NewFloat(0.0).Sub(adjustmentUnits, big.NewFloat(float64(adjUnitsInt)))

		adjustment := big.NewFloat(float64(int(adjustmentLoops) * lengthSum))

		var lengthNum int
		for lengthNum = self.UnixEpoch; lengthNum <= int(adjUnitsInt); lengthNum++ {
			adjustment.Add(
				adjustment,
				big.NewFloat(float64(self.Lengths[lengthNum%lengthCount].ScaleAmount)),
			)
		}

		out.Add(adjustment, big.NewFloat(0.0).Mul(
			adjustmentRemains,
			big.NewFloat(float64(self.Lengths[(lengthNum-1)%lengthCount].ScaleAmount)),
		))
	}

	return
}

func (self *Unit) MarshalJSON() ([]byte, error) {
	object := map[string]interface{}{
		"internal_name": self.InternalName,
		"names":         self.Names,
		"lengths":       self.Lengths,
		"eras":          self.Eras,
		"scale_amount":  self.ScaleAmount,
		"scale_inverse": self.ScaleInverse,
		"uses_zero":     self.UsesZero,
		"unix_epoch":    self.UnixEpoch,
		"is_auxiliary":  self.IsAuxiliary,
	}

	if self.ScaleTo != nil {
		object["scale_to"] = self.ScaleTo.InternalName
	} else {
		object["scale_to"] = nil
	}

	return json.Marshal(object)
}
