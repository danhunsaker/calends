//go:build js && wasm

package main

import "syscall/js"

func calendsAdd(p uint64, offset interface{}, calendar string) js.Value {
	c := instGet(p)
	out, err := c.Add(offset, calendar)
	if err != nil {
		panic(err)
	}
	return js.ValueOf(instNum(out))
}

func CalendsAddString(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	offset := args[1].String()
	calendar := args[2].String()

	return calendsAdd(p, offset, calendar)
}

func CalendsAddInt64(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	offset := args[1].Int()
	calendar := args[2].String()

	return calendsAdd(p, offset, calendar)
}

func CalendsAddDouble(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	offset := args[1].Float()
	calendar := args[2].String()

	return calendsAdd(p, offset, calendar)
}

func calendsSubtract(p uint64, offset interface{}, calendar string) js.Value {
	c := instGet(p)
	out, err := c.Subtract(offset, calendar)
	if err != nil {
		panic(err)
	}
	return js.ValueOf(instNum(out))
}

func CalendsSubtractString(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	offset := args[1].String()
	calendar := args[2].String()

	return calendsSubtract(p, offset, calendar)
}

func CalendsSubtractInt64(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	offset := args[1].Int()
	calendar := args[2].String()

	return calendsSubtract(p, offset, calendar)
}

func CalendsSubtractDouble(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	offset := args[1].Float()
	calendar := args[2].String()

	return calendsSubtract(p, offset, calendar)
}

func calendsAddFromEnd(p uint64, offset interface{}, calendar string) js.Value {
	c := instGet(p)
	out, err := c.AddFromEnd(offset, calendar)
	if err != nil {
		panic(err)
	}
	return js.ValueOf(instNum(out))
}

func CalendsAddFromEndString(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	offset := args[1].String()
	calendar := args[2].String()

	return calendsAddFromEnd(p, offset, calendar)
}

func CalendsAddFromEndInt64(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	offset := args[1].Int()
	calendar := args[2].String()

	return calendsAddFromEnd(p, offset, calendar)
}

func CalendsAddFromEndDouble(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	offset := args[1].Float()
	calendar := args[2].String()

	return calendsAddFromEnd(p, offset, calendar)
}

func calendsSubtractFromEnd(p uint64, offset interface{}, calendar string) js.Value {
	c := instGet(p)
	out, err := c.SubtractFromEnd(offset, calendar)
	if err != nil {
		panic(err)
	}
	return js.ValueOf(instNum(out))
}

func CalendsSubtractFromEndString(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	offset := args[1].String()
	calendar := args[2].String()

	return calendsSubtractFromEnd(p, offset, calendar)
}

func CalendsSubtractFromEndInt64(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	offset := args[1].Int()
	calendar := args[2].String()

	return calendsSubtractFromEnd(p, offset, calendar)
}

func CalendsSubtractFromEndDouble(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	offset := args[1].Float()
	calendar := args[2].String()

	return calendsSubtractFromEnd(p, offset, calendar)
}

func calendsNext(p uint64, offset interface{}, calendar string) js.Value {
	c := instGet(p)
	out, err := c.Next(offset, calendar)
	if err != nil {
		panic(err)
	}
	return js.ValueOf(instNum(out))
}

func CalendsNextString(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	offset := args[1].String()
	calendar := args[2].String()

	return calendsNext(p, offset, calendar)
}

func CalendsNextInt64(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	offset := args[1].Int()
	calendar := args[2].String()

	return calendsNext(p, offset, calendar)
}

func CalendsNextDouble(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	offset := args[1].Float()
	calendar := args[2].String()

	return calendsNext(p, offset, calendar)
}

func calendsPrevious(p uint64, offset interface{}, calendar string) js.Value {
	c := instGet(p)
	out, err := c.Previous(offset, calendar)
	if err != nil {
		panic(err)
	}
	return js.ValueOf(instNum(out))
}

func CalendsPreviousString(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	offset := args[1].String()
	calendar := args[2].String()

	return calendsPrevious(p, offset, calendar)
}

func CalendsPreviousInt64(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	offset := args[1].Int()
	calendar := args[2].String()

	return calendsPrevious(p, offset, calendar)
}

func CalendsPreviousDouble(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	offset := args[1].Float()
	calendar := args[2].String()

	return calendsPrevious(p, offset, calendar)
}

func calendsWithDate(p uint64, stamp interface{}, calendar, format string) js.Value {
	c := instGet(p)
	out, err := c.SetDate(stamp, calendar, format)
	if err != nil {
		panic(err)
	}
	return js.ValueOf(instNum(out))
}

func CalendsWithDateString(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	stamp := args[1].String()
	calendar := args[2].String()
	format := args[3].String()

	return calendsWithDate(p, stamp, calendar, format)
}

func CalendsWithDateInt64(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	stamp := args[1].Int()
	calendar := args[2].String()
	format := args[3].String()

	return calendsWithDate(p, stamp, calendar, format)
}

func CalendsWithDateDouble(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	stamp := args[1].Float()
	calendar := args[2].String()
	format := args[3].String()

	return calendsWithDate(p, stamp, calendar, format)
}

func calendsWithEndDate(p uint64, stamp interface{}, calendar, format string) js.Value {
	c := instGet(p)
	out, err := c.SetEndDate(stamp, calendar, format)
	if err != nil {
		panic(err)
	}
	return js.ValueOf(instNum(out))
}

func CalendsWithEndDateString(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	stamp := args[1].String()
	calendar := args[2].String()
	format := args[3].String()

	return calendsWithEndDate(p, stamp, calendar, format)
}

func CalendsWithEndDateInt64(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	stamp := args[1].Int()
	calendar := args[2].String()
	format := args[3].String()

	return calendsWithEndDate(p, stamp, calendar, format)
}

func CalendsWithEndDateDouble(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	stamp := args[1].Float()
	calendar := args[2].String()
	format := args[3].String()

	return calendsWithEndDate(p, stamp, calendar, format)
}

func calendsWithDuration(p uint64, duration interface{}, calendar string) js.Value {
	c := instGet(p)
	out, err := c.SetDuration(duration, calendar)
	if err != nil {
		panic(err)
	}
	return js.ValueOf(instNum(out))
}

func CalendsWithDurationString(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	duration := args[1].String()
	calendar := args[2].String()

	return calendsWithDuration(p, duration, calendar)
}

func CalendsWithDurationInt64(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	duration := args[1].Int()
	calendar := args[2].String()

	return calendsWithDuration(p, duration, calendar)
}

func CalendsWithDurationDouble(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	duration := args[1].Float()
	calendar := args[2].String()

	return calendsWithDuration(p, duration, calendar)
}

func calendsWithDurationFromEnd(p uint64, duration interface{}, calendar string) js.Value {
	c := instGet(p)
	out, err := c.SetDurationFromEnd(duration, calendar)
	if err != nil {
		panic(err)
	}
	return js.ValueOf(instNum(out))
}

func CalendsWithDurationFromEndString(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	duration := args[1].String()
	calendar := args[2].String()

	return calendsWithDurationFromEnd(p, duration, calendar)
}

func CalendsWithDurationFromEndInt64(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	duration := args[1].Int()
	calendar := args[2].String()

	return calendsWithDurationFromEnd(p, duration, calendar)
}

func CalendsWithDurationFromEndDouble(this js.Value, args []js.Value) interface{} {
	p := uint64(args[0].Int())
	duration := args[1].Float()
	calendar := args[2].String()

	return calendsWithDurationFromEnd(p, duration, calendar)
}

func CalendsMerge(this js.Value, args []js.Value) interface{} {
	p1 := uint64(args[0].Int())
	p2 := uint64(args[1].Int())

	c := instGet(p1)
	z := instGet(p2)
	out, err := c.Merge(*z)
	if err != nil {
		panic(err)
	}
	return js.ValueOf(instNum(out))
}

func CalendsIntersect(this js.Value, args []js.Value) interface{} {
	p1 := uint64(args[0].Int())
	p2 := uint64(args[1].Int())

	c := instGet(p1)
	z := instGet(p2)
	out, err := c.Intersect(*z)
	if err != nil {
		panic(err)
	}
	return js.ValueOf(instNum(out))
}

func CalendsGap(this js.Value, args []js.Value) interface{} {
	p1 := uint64(args[0].Int())
	p2 := uint64(args[1].Int())

	c := instGet(p1)
	z := instGet(p2)
	out, err := c.Gap(*z)
	if err != nil {
		panic(err)
	}
	return js.ValueOf(instNum(out))
}
