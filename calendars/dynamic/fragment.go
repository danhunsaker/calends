package dynamic

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type Fragment struct {
	// The Calendar object this fragment belongs to
	Calendar *Calendar `json:"-"`
	// A single character used to select this fragmet in a Format.String
	Code rune `json:"code"`
	// Follows the same rules as for fmt.Sprintf(), with the addition of
	// {value}, {length}, or {era} immediately following the % character to
	// indicate which of these values should be passed at each point. In the
	// case of {value}, parsing and formatting will work with either the textual
	// version (for "%{value}s"), or the numerical version (for any numerical
	// formats).
	String string `json:"string"`
	// An explanation of what the fragment is meant to encode and how
	Description string `json:"description"`
	// Whether the fragment encodes an era value, vice a unit value directly
	IsEra bool `json:"is_era"`
	// An Era object if Fragment.IsEra is true; otherwise is nil
	Era *Era `json:"era,name"`
	// A Unit object if Fragment.IsEra is false; otherwise is nil
	Unit *Unit `json:"unit,name"`
	// Any textual representations for valid values
	Texts []*FragmentText `json:"texts"`
}

func (self *Fragment) parse(in string) (unit string, value, end int, err error) {
	types := map[rune]string{
		'b': "binary number",
		'c': "character",
		'd': "decimal integer",
		'e': "exponentiated float",
		'f': "floating point number",
		'g': "e or f",
		'o': "octal integer",
		's': "string",
		'x': "hexadecimal integer (lowercase)",
		'X': "hexadecimal integer (uppercase)",
	}

	tags := [4]string{
		"value",
		"length",
		"era",
		"value",
	}

	inSpec := false
	inTag := false
	lastTag := 0
	spec := ""
	tag := ""
	parsed := map[string]interface{}{}
	for _, character := range self.String {
		if inSpec {
			if inTag {
				if character == '}' {
					inTag = false
				} else {
					tag += string(character)
				}
			} else if character == '{' {
				inTag = true
			} else if character == '%' {
				inSpec = false
				if []rune(in[end:])[0] == character {
					end += len(string(character))
				} else {
					err = fmt.Errorf("Input '%c' in '%s' doesn't match fragment '%s'", []rune(in[end:])[0], in, self.String)
					return
				}
			} else if _, ok := types[character]; ok {
				inSpec = false
				// handle strings without tags according to Sprintf rules
				if tag == "" {
					tag = tags[lastTag + 1]
				}
				for t, v := range tags {
					if v == tag {
						lastTag = t
						break
					}
				}
				// now the actual parsing
				working := ""
				// [width][.[precision]]
				split := strings.Split(spec, ".")
				// use default width ("whatever it takes") if width is empty
				width := -1
				if (len(split[0]) > 0) {
					width, _ = strconv.Atoi(split[0])
				}
				// use default precision if precision is unset, and 0 if empty
				prec := -1
				if (len(split) > 1) {
					prec, _ = strconv.Atoi(split[1])
				}

				switch character {
				case 'b': // binary integer
					working, end, err = self.parseSegment(func(chr string) bool {
						return chr == "0" || chr == "1"
					}, in, end, width, prec)
					if err != nil {
						return
					}
					parsed[tag], _ = strconv.ParseInt(working, 2, 0)
				case 'o': // octal integer
					working, end, err = self.parseSegment(func(chr string) bool {
						return chr >= "0" && chr <= "8"
					}, in, end, width, prec)
					if err != nil {
						return
					}
					parsed[tag], _ = strconv.ParseInt(working, 8, 0)
				case 'd': // decimal integer
					working, end, err = self.parseSegment(func(chr string) bool {
						return unicode.IsNumber([]rune(chr)[0])
					}, in, end, width, prec)
					if err != nil {
						return
					}
					parsed[tag], _ = strconv.Atoi(working)
				case 'x', 'X': // hexadecimal integer
					working, end, err = self.parseSegment(func(chr string) bool {
						return (chr >= "0" && chr <= "9") || (chr >= "A" && chr <= "F") || (chr >= "a" && chr <= "f")
					}, in, end, width, prec)
					if err != nil {
						return
					}
					parsed[tag], _ = strconv.ParseInt(working, 16, 0)
				case 'e', 'f', 'g': // floating point decimal number
					working, end, err = self.parseSegment(func(chr string) bool {
						return unicode.IsNumber([]rune(chr)[0]) || chr == "." || chr == "e" || chr == "E"
					}, in, end, width, prec)
					if err != nil {
						return
					}
					parsed[tag], _ = strconv.ParseFloat(working, 64)
				case 'c': // single character
					parsed[tag] = []rune(in[end:])[0]
					end += len(string(parsed[tag].(rune)))
				case 's': // string
					var valid func(string) bool
					if width > 0 {
						valid = func(chr string) bool {
							return true
						}
					} else {
						valid = func(chr string) bool {
							return chr != " "
						}
					}
					parsed[tag], end, err = self.parseSegment(valid, in, end, width, prec)
					if err != nil {
						return
					}
				}
				tag = ""
				spec = ""
			} else if !(character == '.' || unicode.IsNumber(character)) {
				err = fmt.Errorf("Fragment '%s' is malformatted", self.String)
				return
			} else {
				spec += string(character)
			}
		} else if character == '%' {
			inSpec = true
			inTag = false
		} else if []rune(in[end:])[0] == character {
			end += len(string(character))
		} else {
			err = fmt.Errorf("Input '%c' in '%s' doesn't match fragment '%s'", []rune(in[end:])[0], in, self.String)
			return
		}
	}

	// TODO: The input has been parsed; now generate the appropriate output

	return
}

func (self *Fragment) parseSegment(valid func(string) bool, in string, start, width, prec int) (out string, end int, err error) {
	end = start
	if end >= len(in) {
		return
	}
	if width > 0 {
		for idx := 0; idx < width && end < len(in); idx++ {
			chr := string([]rune(in[end:])[0])
			if valid(chr) {
				out += chr
				end += len(chr)
			} else {
				err = fmt.Errorf("Input '%s' in '%s' doesn't match fragment '%s'", chr, in, self.String)
				end -= len(out)
				return
			}
		}
	} else {
		for chr := string([]rune(in[end:])[0]); valid(chr) && end < len(in); chr = string([]rune(in[end:])[0]) {
			out += chr
			end += len(chr)
			if end >= len(in) {
				break
			}
		}
	}
	return
}

func (self *Fragment) format(units map[string]int) (out string) {
	format := self.String
	format = regexp.MustCompile("%\\{value\\}([\\d.]*)([bcdeEfFgGoxX])").ReplaceAllString(format, "%${1}[1]${2}")
	format = regexp.MustCompile("%\\{length\\}([\\d.]*)([bcdeEfFgGoxX])").ReplaceAllString(format, "%${1}[2]${2}")
	format = regexp.MustCompile("%\\{era\\}([\\d.]*)s").ReplaceAllString(format, "%${1}[3]s")
	format = regexp.MustCompile("%\\{value\\}([\\d.]*)s").ReplaceAllString(format, "%${1}[4]s")

	number, length, era, text := self.valueFromUnits(units)
	out = fmt.Sprintf(format, number, length, era, text)

	return
}

func (self *Fragment) valueFromUnits(units map[string]int) (number, length int, era, text string) {
	var unit *Unit

	if self.IsEra {
		unit = self.Era.Unit
	} else {
		unit = self.Unit
	}

	number = units[unit.InternalName]

	if len(unit.Lengths) > 0 {
		length = 0
		for _, l := range unit.Lengths {
			if l.UnitValue == number {
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
				adjusted = number - r.StartDisplay
				if (r.OpenEnded || r.EndValue >= adjusted) && r.StartValue <= adjusted {
					myRange = r
					break
				}
			} else {
				adjusted = r.StartDisplay - number
				if (r.OpenEnded || r.EndValue <= adjusted) && r.StartValue >= adjusted {
					myRange = r
					break
				}
			}
		}

		number = myRange.StartValue + adjusted
		era = myRange.RangeCode
	}

	text = ""
	if len(self.Texts) > 0 {
		for _, t := range self.Texts {
			if t.Value == number {
				text = t.Text
				break
			}
		}
	}

	return
}

func (self *Fragment) MarshalJSON() ([]byte, error) {
	object := map[string]interface{}{
		"code":        string(self.Code),
		"string":      self.String,
		"description": self.Description,
		"is_era":      self.IsEra,
		"texts":       self.Texts,
	}
	if self.IsEra {
		object["era"] = self.Era.InternalName
	} else {
		object["unit"] = self.Unit.InternalName
	}

	return json.Marshal(object)
}
