package dynamic

import ()

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

func (f *Format) parse(in string) (out map[string]int, err error) {
	return
}

func (f *Format) getFragments() (fragments []*Fragment) {
	return
}
