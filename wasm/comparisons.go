//go:build js && wasm

package main

import "syscall/js"

func CalendsDifference(this js.Value, args []js.Value) interface{} {
	p1 := uint64(args[0].Int())
	p2 := uint64(args[1].Int())
	mode := args[2].String()

	c := instGet(p1)
	z := instGet(p2)
	out := c.Difference(*z, mode)
	return js.ValueOf(out.String())
}

func CalendsCompare(this js.Value, args []js.Value) interface{} {
	p1 := uint64(args[0].Int())
	p2 := uint64(args[1].Int())
	mode := args[2].String()

	c := instGet(p1)
	z := instGet(p2)
	return js.ValueOf(int8(c.Compare(*z, mode)))
}

func CalendsIsSame(this js.Value, args []js.Value) interface{} {
	p1 := uint64(args[0].Int())
	p2 := uint64(args[1].Int())

	c := instGet(p1)
	z := instGet(p2)
	return js.ValueOf(c.IsSame(*z))
}

func CalendsIsSameDuration(this js.Value, args []js.Value) interface{} {
	p1 := uint64(args[0].Int())
	p2 := uint64(args[1].Int())

	c := instGet(p1)
	z := instGet(p2)
	return js.ValueOf(c.IsSameDuration(*z))
}

func CalendsIsShorter(this js.Value, args []js.Value) interface{} {
	p1 := uint64(args[0].Int())
	p2 := uint64(args[1].Int())

	c := instGet(p1)
	z := instGet(p2)
	return js.ValueOf(c.IsShorter(*z))
}

func CalendsIsLonger(this js.Value, args []js.Value) interface{} {
	p1 := uint64(args[0].Int())
	p2 := uint64(args[1].Int())

	c := instGet(p1)
	z := instGet(p2)
	return js.ValueOf(c.IsLonger(*z))
}

func CalendsContains(this js.Value, args []js.Value) interface{} {
	p1 := uint64(args[0].Int())
	p2 := uint64(args[1].Int())

	c := instGet(p1)
	z := instGet(p2)
	return js.ValueOf(c.Contains(*z))
}

func CalendsOverlaps(this js.Value, args []js.Value) interface{} {
	p1 := uint64(args[0].Int())
	p2 := uint64(args[1].Int())

	c := instGet(p1)
	z := instGet(p2)
	return js.ValueOf(c.Overlaps(*z))
}

func CalendsAbuts(this js.Value, args []js.Value) interface{} {
	p1 := uint64(args[0].Int())
	p2 := uint64(args[1].Int())

	c := instGet(p1)
	z := instGet(p2)
	return js.ValueOf(c.Abuts(*z))
}

func CalendsIsBefore(this js.Value, args []js.Value) interface{} {
	p1 := uint64(args[0].Int())
	p2 := uint64(args[1].Int())

	c := instGet(p1)
	z := instGet(p2)
	return js.ValueOf(c.IsBefore(*z))
}

func CalendsStartsBefore(this js.Value, args []js.Value) interface{} {
	p1 := uint64(args[0].Int())
	p2 := uint64(args[1].Int())

	c := instGet(p1)
	z := instGet(p2)
	return js.ValueOf(c.StartsBefore(*z))
}

func CalendsEndsBefore(this js.Value, args []js.Value) interface{} {
	p1 := uint64(args[0].Int())
	p2 := uint64(args[1].Int())

	c := instGet(p1)
	z := instGet(p2)
	return js.ValueOf(c.EndsBefore(*z))
}

func CalendsIsDuring(this js.Value, args []js.Value) interface{} {
	p1 := uint64(args[0].Int())
	p2 := uint64(args[1].Int())

	c := instGet(p1)
	z := instGet(p2)
	return js.ValueOf(c.IsDuring(*z))
}

func CalendsStartsDuring(this js.Value, args []js.Value) interface{} {
	p1 := uint64(args[0].Int())
	p2 := uint64(args[1].Int())

	c := instGet(p1)
	z := instGet(p2)
	return js.ValueOf(c.StartsDuring(*z))
}

func CalendsEndsDuring(this js.Value, args []js.Value) interface{} {
	p1 := uint64(args[0].Int())
	p2 := uint64(args[1].Int())

	c := instGet(p1)
	z := instGet(p2)
	return js.ValueOf(c.EndsDuring(*z))
}

func CalendsIsAfter(this js.Value, args []js.Value) interface{} {
	p1 := uint64(args[0].Int())
	p2 := uint64(args[1].Int())

	c := instGet(p1)
	z := instGet(p2)
	return js.ValueOf(c.IsAfter(*z))
}

func CalendsStartsAfter(this js.Value, args []js.Value) interface{} {
	p1 := uint64(args[0].Int())
	p2 := uint64(args[1].Int())

	c := instGet(p1)
	z := instGet(p2)
	return js.ValueOf(c.StartsAfter(*z))
}

func CalendsEndsAfter(this js.Value, args []js.Value) interface{} {
	p1 := uint64(args[0].Int())
	p2 := uint64(args[1].Int())

	c := instGet(p1)
	z := instGet(p2)
	return js.ValueOf(c.EndsAfter(*z))
}
