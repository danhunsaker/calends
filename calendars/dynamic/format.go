package dynamic

import (
	"errors"
	"fmt"
)

type Format struct {
	Calendar    *Calendar
	Name        string
	String      string
	Description string
}

func formatFromString(c *Calendar, in string) (out *Format) {
	out.Calendar = c
	out.Name = in
	out.String = in
	out.Description = "format from string"

	return
}

func (self *Format) parse(in string) (out map[string]int, err error) {
	fragments := self.getFragments()

	start := 0
	isEscaped := false
	for _, character := range self.String {
		if character == '\\' {
			isEscaped = true
			continue
		}

		if isEscaped && []rune(in[start:])[0] == character {
			start += len(string(character))
			isEscaped = false
			continue
		}

		if fragment, ok := fragments[character]; !isEscaped && ok {
			var unit string
			var value, end int
			unit, value, end, err = fragment.parse(string(in[start:]))
			if err != nil {
				return
			}
			out[unit] = value
			start += end
			continue
		}

		err = errors.New(fmt.Sprintf("Input %s doesn't match format %s", in, self.Name))
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
