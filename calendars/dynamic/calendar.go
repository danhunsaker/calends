package dynamic

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
)

type Calendar struct {
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	Units         []*Unit     `json:"units"`
	BaseUnit      *Unit       `json:"base_unit,name"`
	Formats       []*Format   `json:"formats"`
	DefaultFormat *Format     `json:"default_format,name"`
	Fragments     []*Fragment `json:"fragments"`
}

func (self *Calendar) ToTimestamp(date interface{}, format string) (stamp *big.Float, err error) {
	var dateString string

	switch date.(type) {
	case big.Float:
		tmp := date.(big.Float)
		dateString = tmp.String()
	case *big.Float:
		dateString = date.(*big.Float).String()
	case float64:
		dateString = fmt.Sprintf("%f", date.(float64))
	case int:
		dateString = fmt.Sprintf("%d", date.(int))
	case string:
		dateString = date.(string)
	case []byte:
		dateString = string(date.([]byte))
	default:
		err = errors.New("Unsupported Value")
		return
	}

	var formatObject *Format
	if format == "" {
		formatObject = self.DefaultFormat
	} else {
		for _, f := range self.Formats {
			if f.Name == format {
				formatObject = f
				break
			}
		}
	}
	if formatObject == nil {
		formatObject = self.formatFromString(format)
	}

	stamp = self.unitsToTime(self.dateToUnits(dateString, formatObject))

	return
}

func (self *Calendar) FromTimestamp(stamp *big.Float, format string) (string, error) {
	var formatObject *Format
	if format == "" {
		formatObject = self.DefaultFormat
	} else {
		for _, f := range self.Formats {
			if f.Name == format {
				formatObject = f
				break
			}
		}
	}
	if formatObject == nil {
		formatObject = self.formatFromString(format)
	}

	return self.unitsToDate(self.timeToUnits(stamp), formatObject), nil
}

func (self *Calendar) Offset(in *big.Float, offset interface{}) (out *big.Float, err error) {
	var offsetString string
	units := make(map[string]int, 0)

	switch offset.(type) {
	case big.Float:
		tmp := offset.(big.Float)
		offsetString = tmp.String()
	case *big.Float:
		offsetString = offset.(*big.Float).String()
	case float64:
		offsetString = fmt.Sprintf("%f", offset.(float64))
	case int:
		offsetString = fmt.Sprintf("%d", offset.(int))
	case string:
		offsetString = offset.(string)
	case []byte:
		offsetString = string(offset.([]byte))
	default:
		err = errors.New("Unsupported Value")
		return
	}

	units, err = self.unitsWithOffset(self.timeToUnits(in), offsetString)
	if err == nil {
		out = self.unitsToTime(units)
	}

	return
}

func (self *Calendar) dateToUnits(in string, format *Format) (out map[string]int) {
	formats := append([]*Format{format, self.DefaultFormat}, self.Formats...)

	for _, f := range formats {
		out, err := f.parse(in)
		if err == nil {
			return out
		}
	}

	return self.epochUnits(true)
}

func (self *Calendar) unitsToDate(in map[string]int, format *Format) (out string) {
	return format.format(in)
}

func (self *Calendar) unitsWithOffset(in map[string]int, offset string) (out map[string]int, err error) {
	out = make(map[string]int, 0)
	names := make(map[string]int, 0)
	var matchUnit *Unit
	var matchValue int64

	re := regexp.MustCompile("(?P<value>[-+]?[0-9]+)\\s*(?P<unit>\\S+)?")
	for idx, name := range re.SubexpNames() {
		if name == "" {
			continue
		}
		names[name] = idx
	}

	for _, match := range re.FindAllStringSubmatch(offset, -1) {
		matchUnit = nil
	FindUnit:
		for _, unit := range self.Units {
			if unit.InternalName == match[names["unit"]] {
				matchUnit = unit
				break FindUnit
			}

			for _, name := range unit.Names {
				if name.UnitName == match[names["unit"]] {
					matchUnit = unit
					break FindUnit
				}
			}
		}
		if matchUnit == nil {
			matchUnit = self.BaseUnit
		}

		matchValue, err = strconv.ParseInt(match[names["value"]], 0, 0)
		if err != nil {
			return
		}

		unitName, unitValue := matchUnit.reduceAuxiliary(int(matchValue))
		if currentValue, ok := out[unitName]; ok {
			unitValue += currentValue
		}

		out[unitName] = unitValue
	}

	out = self.BaseUnit.carryOver(out)

	return
}

func (self *Calendar) unitsToTime(in map[string]int) *big.Float {
	return self.BaseUnit.toSeconds(self.sumUnits(self.epochUnits(false), in))
}

func (self *Calendar) timeToUnits(in *big.Float) map[string]int {
	return self.sumUnits(self.epochUnits(true), self.BaseUnit.fromSeconds(in))
}

func (self *Calendar) sumUnits(a, b map[string]int) (out map[string]int) {
	out = make(map[string]int, 0)
	for aKey, aVal := range a {
		out[aKey] = aVal

		if bVal, ok := b[aKey]; ok {
			out[aKey] = aVal + bVal
		}
	}

	for bKey, bVal := range b {
		if _, ok := out[bKey]; !ok {
			out[bKey] = bVal
		}
	}

	out = self.BaseUnit.carryOver(out)

	return
}

func (self *Calendar) epochUnits(positive bool) (out map[string]int) {
	out = make(map[string]int, 0)
	for _, unit := range self.Units {
		if unit.IsAuxiliary {
			continue
		}

		if positive {
			out[unit.InternalName] = unit.UnixEpoch
		} else {
			zero := 1
			if unit.UsesZero {
				zero = 0
			}

			out[unit.InternalName] = -1 * (unit.UnixEpoch - zero)
		}
	}

	return
}

// We implement `enocding/json.[Un]Marshaler` because the `name` option in the
// `json:` tags, above, is entirely made up, so we need to do what it indicates
// "manually".

func (self *Calendar) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"name":           self.Name,
		"description":    self.Description,
		"units":          self.Units,
		"base_unit":      self.BaseUnit.InternalName,
		"formats":        self.Formats,
		"default_format": self.DefaultFormat.Name,
		"fragments":      self.Fragments,
	})
}

func (out *Calendar) UnmarshalJSON(in []byte) (err error) {
	var baseUnit, defaultFormat string
	raw := make(map[string]interface{}, 7)
	scaleTo := make(map[*Unit]string, 0)
	defaultRange := make(map[*Era]string, 0)
	fragmentEras := make(map[*Fragment]string, 0)
	fragmentUnits := make(map[*Fragment]string, 0)

	err = json.Unmarshal(in, &raw)
	if err != nil {
		return
	}

	for k, v := range raw {
		switch k {
		case "name":
			out.Name = v.(string)
		case "description":
			out.Description = v.(string)
		case "units":
			if v != nil {
				for _, ui := range v.([]interface{}) {
					unit := Unit{Calendar: out}
					for uk, uv := range ui.(map[string]interface{}) {
						switch uk {
						case "internal_name":
							unit.InternalName = uv.(string)
						case "names":
							if uv != nil {
								for _, uni := range uv.([]interface{}) {
									unitName := UnitName{Unit: &unit}
									for unk, unv := range uni.(map[string]interface{}) {
										switch unk {
										case "unit_name":
											unitName.UnitName = unv.(string)
										case "name_context":
											unitName.NameContext = unv.(string)
										}
									}
									unit.Names = append(unit.Names, &unitName)
								}
							}
						case "lengths":
							if uv != nil {
								for _, uli := range uv.([]interface{}) {
									unitLength := UnitLength{Unit: &unit}
									for ulk, ulv := range uli.(map[string]interface{}) {
										switch ulk {
										case "unit_value":
											unitLength.UnitValue = int(ulv.(float64))
										case "scale_amount":
											unitLength.ScaleAmount = int(ulv.(float64))
										}
									}
									unit.Lengths = append(unit.Lengths, &unitLength)
								}
							}
						case "eras":
							if uv != nil {
								for _, uei := range uv.([]interface{}) {
									era := Era{Calendar: out, Unit: &unit}
									for uek, uev := range uei.(map[string]interface{}) {
										switch uek {
										case "internal_name":
											era.InternalName = uev.(string)
										case "ranges":
											if uev != nil {
												for _, ueri := range uev.([]interface{}) {
													eraRange := EraRange{Era: &era}
													for uerk, uerv := range ueri.(map[string]interface{}) {
														switch uerk {
														case "range_code":
															eraRange.RangeCode = uerv.(string)
														case "start_value":
															eraRange.StartValue = int(uerv.(float64))
														case "start_display":
															eraRange.StartDisplay = int(uerv.(float64))
														case "end_value":
															eraRange.EndValue = int(uerv.(float64))
														case "open_ended":
															eraRange.OpenEnded = uerv.(bool)
														case "ascending":
															eraRange.Ascending = uerv.(bool)
														}
													}
													era.Ranges = append(era.Ranges, &eraRange)
												}
											}
										case "default_range":
											defaultRange[&era] = uev.(string)
										}
									}
									unit.Eras = append(unit.Eras, &era)
								}
							}
						case "scale_to":
							if uv != nil {
								scaleTo[&unit] = uv.(string)
							}
						case "scale_amount":
							unit.ScaleAmount = int(uv.(float64))
						case "scale_inverse":
							unit.ScaleInverse = uv.(bool)
						case "uses_zero":
							unit.UsesZero = uv.(bool)
						case "unix_epoch":
							unit.UnixEpoch = int(uv.(float64))
						case "is_auxiliary":
							unit.IsAuxiliary = uv.(bool)
						}
					}
					out.Units = append(out.Units, &unit)
				}
			}
		case "base_unit":
			baseUnit = v.(string)
		case "formats":
			if v != nil {
				for _, fi := range v.([]interface{}) {
					format := Format{Calendar: out}
					for fk, fv := range fi.(map[string]interface{}) {
						switch fk {
						case "name":
							format.Name = fv.(string)
						case "string":
							format.String = fv.(string)
						case "description":
							format.Description = fv.(string)
						}
					}
					out.Formats = append(out.Formats, &format)
				}
			}
		case "default_format":
			defaultFormat = v.(string)
		case "fragments":
			if v != nil {
				for _, fi := range v.([]interface{}) {
					fragment := Fragment{Calendar: out}
					for fk, fv := range fi.(map[string]interface{}) {
						switch fk {
						case "code":
							fragment.Code = []rune(fv.(string))[0]
						case "string":
							fragment.String = fv.(string)
						case "description":
							fragment.Description = fv.(string)
						case "is_era":
							fragment.IsEra = fv.(bool)
						case "era":
							fragmentEras[&fragment] = fv.(string)
						case "unit":
							fragmentUnits[&fragment] = fv.(string)
						case "texts":
							if fv != nil {
								for _, fti := range fv.([]interface{}) {
									fragmentText := FragmentText{Fragment: &fragment}
									for ftk, ftv := range fti.(map[string]interface{}) {
										switch ftk {
										case "value":
											fragmentText.Value = int(ftv.(float64))
										case "text":
											fragmentText.Text = ftv.(string)
										}
									}
									fragment.Texts = append(fragment.Texts, &fragmentText)
								}
							}
						}
					}
					out.Fragments = append(out.Fragments, &fragment)
				}
			}
		}
	}

	// Assign out.BaseUnit
	for _, u := range out.Units {
		if u.InternalName == baseUnit {
			out.BaseUnit = u
			break
		}
	}
	// Assign out.DefaultFormat
	for _, f := range out.Formats {
		if f.Name == defaultFormat {
			out.DefaultFormat = f
			break
		}
	}
	// Assign unit ScaleTo values
	for u1, name := range scaleTo {
		for _, u2 := range out.Units {
			if u2.InternalName == name {
				u1.ScaleTo = u2
				break
			}
		}
	}
	// Assign era DefaultRange values
	for e, name := range defaultRange {
		for _, r := range e.Ranges {
			if r.RangeCode == name {
				e.DefaultRange = r
				break
			}
		}
	}
	// Assign fragment Era values
	for f, name := range fragmentEras {
		for _, u := range out.Units {
			for _, e := range u.Eras {
				if e.InternalName == name {
					f.Era = e
					break
				}
			}

			if f.Era != nil {
				break
			}
		}
	}
	// Assign fragment Unit values
	for f, name := range fragmentUnits {
		for _, u := range out.Units {
			if u.InternalName == name {
				f.Unit = u
				break
			}
		}
	}

	// Fill unit and era Formats lists
	for _, f := range out.Fragments {
		if f.IsEra {
			f.Era.Formats = append(f.Era.Formats, f)
		} else {
			f.Unit.Formats = append(f.Unit.Formats, f)
		}
	}
	// Fill unit ScalesToMe lists
	for _, u := range out.Units {
		if u.ScaleTo != nil {
			u.ScaleTo.ScalesToMe = append(u.ScaleTo.ScalesToMe, u)
		}
	}

	return
}
