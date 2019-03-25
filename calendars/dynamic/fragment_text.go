package dynamic

import ()

type FragmentText struct {
	Fragment *Fragment `json:"-"`
	Value    int       `json:"value"`
	Text     string    `json:"text"`
}
