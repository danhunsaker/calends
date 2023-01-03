//go:build js && wasm

package main

import (
	"syscall/js"

	"github.com/danhunsaker/calends"
)

func CalendsRelease(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())

	instances.Delete(p)
	return js.Undefined()
}

func CalendsCreateString(this js.Value, args []js.Value) interface{} {
	stamp := args[0].String()
	calendar := args[1].String()
	format := args[2].String()

	return js.ValueOf(calends_create(stamp, calendar, format))
}

func CalendsCreateStringRange(this js.Value, args []js.Value) interface{} {
	start := args[0].String()
	end := args[1].String()
	calendar := args[2].String()
	format := args[3].String()

	return js.ValueOf(calends_create(map[string]interface{}{
		"start": start,
		"end":   end,
	}, calendar, format))
}

func CalendsCreateStringStartPeriod(this js.Value, args []js.Value) interface{} {
	start := args[0].String()
	duration := args[1].String()
	calendar := args[2].String()
	format := args[3].String()

	return js.ValueOf(calends_create(map[string]interface{}{
		"start":    start,
		"duration": duration,
	}, calendar, format))
}

func CalendsCreateStringEndPeriod(this js.Value, args []js.Value) interface{} {
	duration := args[0].String()
	end := args[1].String()
	calendar := args[2].String()
	format := args[3].String()

	return js.ValueOf(calends_create(map[string]interface{}{
		"duration": duration,
		"end":      end,
	}, calendar, format))
}

func CalendsCreateInt64(this js.Value, args []js.Value) interface{} {
	stamp := args[0].Int()
	calendar := args[1].String()
	format := args[2].String()

	return js.ValueOf(calends_create(int(stamp), calendar, format))
}

func CalendsCreateInt64Range(this js.Value, args []js.Value) interface{} {
	start := args[0].Int()
	end := args[1].Int()
	calendar := args[2].String()
	format := args[3].String()

	return js.ValueOf(calends_create(map[string]interface{}{
		"start": int(start),
		"end":   int(end),
	}, calendar, format))
}

func CalendsCreateInt64StartPeriod(this js.Value, args []js.Value) interface{} {
	start := args[0].Int()
	duration := args[1].Int()
	calendar := args[2].String()
	format := args[3].String()

	return js.ValueOf(calends_create(map[string]interface{}{
		"start":    int(start),
		"duration": int(duration),
	}, calendar, format))
}

func CalendsCreateInt64EndPeriod(this js.Value, args []js.Value) interface{} {
	duration := args[0].Int()
	end := args[1].Int()
	calendar := args[2].String()
	format := args[3].String()

	return js.ValueOf(calends_create(map[string]interface{}{
		"duration": int(duration),
		"end":      int(end),
	}, calendar, format))
}

func CalendsCreateDouble(this js.Value, args []js.Value) interface{} {
	stamp := args[0].Float()
	calendar := args[1].String()
	format := args[2].String()

	return js.ValueOf(calends_create(float64(stamp), calendar, format))
}

func CalendsCreateDoubleRange(this js.Value, args []js.Value) interface{} {
	start := args[0].Float()
	end := args[1].Float()
	calendar := args[2].String()
	format := args[3].String()

	return js.ValueOf(calends_create(map[string]interface{}{
		"start": float64(start),
		"end":   float64(end),
	}, calendar, format))
}

func CalendsCreateDoubleStartPeriod(this js.Value, args []js.Value) interface{} {
	start := args[0].Float()
	duration := args[1].Float()
	calendar := args[2].String()
	format := args[3].String()

	return js.ValueOf(calends_create(map[string]interface{}{
		"start":    float64(start),
		"duration": float64(duration),
	}, calendar, format))
}

func CalendsCreateDoubleEndPeriod(this js.Value, args []js.Value) interface{} {
	duration := args[0].Float()
	end := args[1].Float()
	calendar := args[2].String()
	format := args[3].String()

	return js.ValueOf(calends_create(map[string]interface{}{
		"duration": float64(duration),
		"end":      float64(end),
	}, calendar, format))
}

func CalendsDate(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	calendar := args[1].String()
	format := args[2].String()

	c := instGet(p)
	out, err := c.Date(calendar, format)
	if err != nil {
		panic(err)
	}
	return js.ValueOf(out)
}

func CalendsDuration(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())

	c := instGet(p)
	return js.ValueOf(c.Duration().String())
}

func CalendsEndDate(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	calendar := args[1].String()
	format := args[2].String()

	c := instGet(p)
	out, err := c.EndDate(calendar, format)
	if err != nil {
		panic(err)
	}
	return js.ValueOf(out)
}

func CalendsString(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())

	c := instGet(p)
	return js.ValueOf(c.String())
}

func CalendsEncodeText(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())

	c := instGet(p)
	out, err := c.MarshalText()
	if err != nil {
		panic(err)
	}
	return js.ValueOf(string(out))
}

func CalendsDecodeText(this js.Value, args []js.Value) interface{} {
	in := args[0].String()

	c := calends.Calends{}
	err := c.UnmarshalText([]byte(in))
	if err != nil {
		panic(err)
	}
	return js.ValueOf(instNum(c))
}

func CalendsEncodeJson(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())

	c := instGet(p)
	out, err := c.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return js.ValueOf(string(out))
}

func CalendsDecodeJson(this js.Value, args []js.Value) interface{} {
	in := args[0].String()

	c := calends.Calends{}
	err := c.UnmarshalJSON([]byte(in))
	if err != nil {
		panic(err)
	}
	return js.ValueOf(instNum(c))
}

func CalendsRegisterPanicHandler(this js.Value, args []js.Value) interface{} {
	handler := args[0]

	id := nextPanHandle.Id()
	panicHandlers.Store(id, handler)

	return js.Undefined()
}
