package calendars

import (
	"io/ioutil"
	"testing"

	"github.com/danhunsaker/calends/calendars/dynamic"
)

func TestRegisterDynamic(t *testing.T) {
	calendar := dynamic.Calendar{}
	json, err:= ioutil.ReadFile("../tests/dynamic.json")
	if err != nil {
		t.Fatalf("Failed to load test dynamic calendar: %#v", err)
	}

	if err := calendar.UnmarshalJSON(json); err != nil {
		t.Fatalf("Failed to unmarshal test dynamic calendar: %#v", err)
	}

	RegisterDynamic(calendar)

	if !Registered("test") {
		t.Errorf("RegisterDynamic(%#v) failed", calendar)
	}
}
