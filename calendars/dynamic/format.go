package dynamic

import (
	"fmt"
)

type Format struct {
	Calendar    *Calendar `json:"-"`
	Name        string    `json:"name"`
	String      string    `json:"string"`
	Description string    `json:"description"`
}

func (self *Calendar) formatFromString(in string) (*Format) {
	out := Format{
		Calendar: self,
		Name: in,
		String: in,
		Description: "format from string",
	}

	return &out
}

func (self *Format) parse(in string) (out map[string]int, err error) {
	out = make(map[string]int, 0)
	fragments := self.getFragments()

	start := 0
	isEscaped := false
	for _, character := range self.String {
		if start >= len(in) {
			break
		}

		if character == '\\' {
			isEscaped = true
			continue
		} else if fragment, ok := fragments[character]; !isEscaped && ok {
			var unit string
			var value, end int
			unit, value, end, err = fragment.parse(string(in[start:]))
			if err != nil {
				return
			}
			out[unit] = value
			start += end
			continue
		} else if []rune(in[start:])[0] == character {
			start += len(string(character))
			isEscaped = false
			continue
		}

		err = fmt.Errorf("Input %s doesn't match format %s", in, self.Name)
		break
	}

	return
}

func (self *Format) format(in map[string]int) (out string) {
	fragments := self.getFragments()

	isEscaped := false
	for _, character := range self.String {
		if character == '\\' {
			isEscaped = true
			continue
		}

		if isEscaped {
			isEscaped = false
		} else if fragment, ok := fragments[character]; ok {
			out += fragment.format(in)
			continue
		}

		out += string(character)
	}

	return
}

func (self *Format) getFragments() (fragments map[rune]*Fragment) {
	fragments = make(map[rune]*Fragment, 0)
	for _, character := range self.String {
		for _, fragment := range self.Calendar.Fragments {
			if fragment.Code == character {
				fragments[character] = fragment
				break
			}
		}
	}

	return
}
