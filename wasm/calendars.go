//go:build js && wasm

package main

import (
	"math/big"
	"strings"
	"syscall"
	"syscall/js"

	"github.com/danhunsaker/calends/calendars"
	"github.com/go-errors/errors"
)

func CalendsCalendarRegister(this js.Value, args []js.Value) interface{} {
	name := args[0].String()
	defaultFormat := args[1].String()

	toInternalString := args[2]
	toInternalInt64 := args[3]
	toInternalDouble := args[4]
	toInternalTai := args[5]

	fromInternal := args[6]

	offsetString := args[7]
	offsetInt64 := args[8]
	offsetDouble := args[9]
	offsetTai := args[10]

	calendars.RegisterElements(
		(name),
		wrapToInternal(name, toInternalString, toInternalInt64, toInternalDouble, toInternalTai),
		wrapFromInternal(name, fromInternal),
		wrapOffset(name, offsetString, offsetInt64, offsetDouble, offsetTai),
		(defaultFormat),
	)

	return js.Undefined()
}

func CalendsCalendarUnregister(this js.Value, args []js.Value) interface{} {
	name := args[0].String()

	calendars.Unregister((name))

	return js.Undefined()
}

func CalendsCalendarRegistered(this js.Value, args []js.Value) interface{} {
	calendar := args[0].String()

	return js.ValueOf(calendars.Registered(calendar))
}

func CalendsCalendarListRegistered(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(strings.Join(calendars.ListRegistered(), "\n"))
}

func Tai64TimeAdd(this js.Value, args []js.Value) interface{} {
	t := args[0]
	z := args[1]

	base := taiWASMToGo(t)
	addend := taiWASMToGo(z)
	return taiGoToWASM(base.Add(addend))
}

func Tai64TimeSub(this js.Value, args []js.Value) interface{} {
	t := args[0]
	z := args[1]

	base := taiWASMToGo(t)
	subtend := taiWASMToGo(z)
	return taiGoToWASM(base.Sub(subtend))
}

func Tai64TimeString(this js.Value, args []js.Value) interface{} {
	t := args[0]

	base := taiWASMToGo(t)
	return js.ValueOf(base.String())
}

func Tai64TimeFromString(this js.Value, args []js.Value) interface{} {
	in := args[0].String()

	return taiGoToWASM(calendars.TAI64NARUXTimeFromDecimalString((in)))
}

func Tai64TimeHexString(this js.Value, args []js.Value) interface{} {
	t := args[0]

	base := taiWASMToGo(t)
	return js.ValueOf(base.HexString())
}

func Tai64TimeFromHexString(this js.Value, args []js.Value) interface{} {
	in := args[0].String()

	return taiGoToWASM(calendars.TAI64NARUXTimeFromHexString((in)))
}

func Tai64TimeDouble(this js.Value, args []js.Value) interface{} {
	t := args[0]

	base := taiWASMToGo(t)
	out, _ := base.Float().Float64()
	return js.ValueOf(out)
}

func Tai64TimeFromDouble(this js.Value, args []js.Value) interface{} {
	in := args[0].Float()

	return taiGoToWASM(calendars.TAI64NARUXTimeFromFloat(*big.NewFloat(in)))
}

func Tai64TimeEncodeText(this js.Value, args []js.Value) interface{} {
	t := args[0]

	base := taiWASMToGo(t)
	out, err := base.MarshalText()
	if err != nil {
		panic(err)
	}
	return js.ValueOf(string(out))
}

func Tai64TimeDecodeText(this js.Value, args []js.Value) interface{} {
	in := args[0].String()

	time := calendars.TAI64NARUXTime{}
	time.UnmarshalText([]byte((in)))
	return taiGoToWASM(time)
}

func Tai64TimeUtcToTai(this js.Value, args []js.Value) interface{} {
	utc := args[0]

	return taiGoToWASM(calendars.UTCtoTAI(taiWASMToGo(utc)))
}

func Tai64TimeTaiToUtc(this js.Value, args []js.Value) interface{} {
	tai := args[0]

	return taiGoToWASM(calendars.TAItoUTC(taiWASMToGo(tai)))
}

func wrapToInternal(
	name string,
	toInternalString js.Value,
	toInternalInt64 js.Value,
	toInternalDouble js.Value,
	toInternalTai js.Value,
) func(interface{}, string) (calendars.TAI64NARUXTime, error) {
	return func(in interface{}, format string) (out calendars.TAI64NARUXTime, err error) {
		var raw js.Value
		switch in := in.(type) {
		case string:
			raw = toInternalString.Invoke(name, (in), (format))
			if err != nil {
				err = errors.Wrap(err, 0)
			}
		case int:
			raw = toInternalInt64.Invoke(name, int64(in), (format))
			if err != nil {
				err = errors.Wrap(err, 0)
			}
		case float64:
			raw = toInternalDouble.Invoke(name, in, (format))
			if err != nil {
				err = errors.Wrap(err, 0)
			}
		case calendars.TAI64NARUXTime:
			pass := taiGoToWASM(in)
			raw = toInternalTai.Invoke(name, pass)
			if err != nil {
				err = errors.Wrap(err, 0)
			}
		default:
			err = errors.Wrap(calendars.ErrUnsupportedInput, 1)
		}
		if err != nil && !errors.Is(err, syscall.Errno(2)) {
			panic(err)
		}
		out = taiWASMToGo(raw)
		err = nil
		return
	}
}

func wrapFromInternal(name string, fromInternal js.Value) func(calendars.TAI64NARUXTime, string) (string, error) {
	return func(in calendars.TAI64NARUXTime, format string) (out string, err error) {
		var raw js.Value
		pass := taiGoToWASM(in)
		raw = fromInternal.Invoke(name, pass, (format))
		if err != nil && !errors.Is(err, syscall.Errno(2)) {
			panic(errors.Wrap(err, 0))
		}
		out = (raw.String())
		err = nil
		return
	}
}

func wrapOffset(
	name string,
	offsetString js.Value,
	offsetInt64 js.Value,
	offsetDouble js.Value,
	offsetTai js.Value,
) func(calendars.TAI64NARUXTime, interface{}) (calendars.TAI64NARUXTime, error) {
	return func(in calendars.TAI64NARUXTime, offset interface{}) (out calendars.TAI64NARUXTime, err error) {
		var raw js.Value
		pass := taiGoToWASM(in)
		switch offset := offset.(type) {
		case string:
			raw = offsetString.Invoke(name, pass, (offset))
			if err != nil {
				err = errors.Wrap(err, 0)
			}
		case int:
			raw = offsetInt64.Invoke(name, pass, int64(offset))
			if err != nil {
				err = errors.Wrap(err, 0)
			}
		case float64:
			raw = offsetDouble.Invoke(name, pass, offset)
			if err != nil {
				err = errors.Wrap(err, 0)
			}
		case calendars.TAI64NARUXTime:
			pass2 := taiGoToWASM(offset)
			raw = offsetTai.Invoke(name, pass, pass2)
			if err != nil {
				err = errors.Wrap(err, 0)
			}
		default:
			err = errors.Wrap(calendars.ErrUnsupportedInput, 1)
		}
		if err != nil && !errors.Is(err, syscall.Errno(2)) {
			panic(err)
		}
		out = taiWASMToGo(raw)
		err = nil
		return
	}
}

func taiGoToWASM(in calendars.TAI64NARUXTime) js.Value {
	return js.ValueOf(map[string]interface{}{
		"seconds":  js.ValueOf(in.Seconds),
		"nano":     js.ValueOf(in.Nano),
		"atto":     js.ValueOf(in.Atto),
		"ronto":    js.ValueOf(in.Ronto),
		"udecto":   js.ValueOf(in.Udecto),
		"xindecto": js.ValueOf(in.Xindecto),
	})
}

func taiWASMToGo(in js.Value) calendars.TAI64NARUXTime {
	return calendars.TAI64NARUXTime{
		Seconds:  int64(in.Get("seconds").Int()),
		Nano:     uint32(in.Get("nano").Int()),
		Atto:     uint32(in.Get("atto").Int()),
		Ronto:    uint32(in.Get("ronto").Int()),
		Udecto:   uint32(in.Get("udecto").Int()),
		Xindecto: uint32(in.Get("xindecto").Int()),
	}
}
