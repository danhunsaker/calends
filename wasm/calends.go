//go:build js && wasm

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"syscall/js"

	"github.com/danhunsaker/calends"
)

var maxShards = uint64(64)

type concurrentMap []*concurrentMapShard

type concurrentMapShard struct {
	items map[uint64]interface{}
	sync.RWMutex
}

func newConcurrentMap() concurrentMap {
	m := make(concurrentMap, maxShards)
	for i := uint64(0); i < maxShards; i++ {
		m[i] = &concurrentMapShard{items: make(map[uint64]interface{})}
	}
	return m
}

func (m concurrentMap) getShard(key uint64) *concurrentMapShard {
	return m[key%maxShards]
}

func (m concurrentMap) Store(key uint64, value interface{}) {
	shard := m.getShard(key)
	shard.Lock()
	shard.items[key] = value
	shard.Unlock()
}

func (m concurrentMap) Load(key uint64) (interface{}, bool) {
	shard := m.getShard(key)
	shard.RLock()
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}

func (m concurrentMap) Length() (out uint64) {
	for _, shard := range m {
		shard.RLock()
		out += uint64(len(shard.items))
		shard.RUnlock()
	}
	return
}

func (m concurrentMap) All() (out []interface{}) {
	for _, shard := range m {
		shard.RLock()
		for _, item := range shard.items {
			out = append(out, item)
		}
		shard.RUnlock()
	}
	return
}

func (m concurrentMap) Delete(key uint64) {
	shard := m.getShard(key)
	shard.Lock()
	delete(shard.items, key)
	shard.Unlock()
}

type idGenerator struct {
	id uint64
}

func (generator *idGenerator) Id() uint64 {
	return atomic.AddUint64(&generator.id, 1)
}

var panicHandlers concurrentMap
var nextPanHandle idGenerator
var instances concurrentMap
var nextInst idGenerator

func instNum(c calends.Calends) uint64 {
	p := nextInst.Id()
	instances.Store(p, &c)
	return p
}

func instGet(p uint64) *calends.Calends {
	val, ok := instances.Load(p)
	if !ok {
		panic(fmt.Sprintf("Calends object #%d doesn't exist", p))
	}
	return val.(*calends.Calends)
}

func calends_create(stamp interface{}, calendar, format string) uint64 {
	c, err := calends.Create(stamp, calendar, format)
	if err != nil {
		panic(err)
	}
	return instNum(c)
}

func handlePanic() {
	var err string

	if r := recover(); r != nil {
		// If nothing has been registered to handle
		// these error states, resume panicking!
		if panicHandlers.Length() < 1 {
			panic(r)
		}

		// We need a string to pass to the handler(s)
		switch r := r.(type) {
		case string:
			err = r
		case error:
			err = r.Error()
			// err = errors.Wrap(r, 2).ErrorStack()
		default:
			err = fmt.Sprintf("%#v", r)
		}

		// Call the handler(s)!
		for _, handler := range panicHandlers.All() {
			handler.(js.Value).Invoke([]js.Value{js.ValueOf(err)})
		}
	}
}

func main() {
	defer handlePanic()

	quit := make(chan struct{})
	panicHandlers = newConcurrentMap()
	instances = newConcurrentMap()

	js.Global().Set("CalendsFuncs", js.ValueOf(map[string]interface{}{
		// Calends Exports
		"release": js.FuncOf(CalendsRelease),

		"createString":            js.FuncOf(CalendsCreateString),
		"createStringRange":       js.FuncOf(CalendsCreateStringRange),
		"createStringStartPeriod": js.FuncOf(CalendsCreateStringStartPeriod),
		"createStringEndPeriod":   js.FuncOf(CalendsCreateStringEndPeriod),

		"createInt64":            js.FuncOf(CalendsCreateInt64),
		"createInt64Range":       js.FuncOf(CalendsCreateInt64Range),
		"createInt64StartPeriod": js.FuncOf(CalendsCreateInt64StartPeriod),
		"createInt64EndPeriod":   js.FuncOf(CalendsCreateInt64EndPeriod),

		"createDouble":            js.FuncOf(CalendsCreateDouble),
		"createDoubleRange":       js.FuncOf(CalendsCreateDoubleRange),
		"createDoubleStartPeriod": js.FuncOf(CalendsCreateDoubleStartPeriod),
		"createDoubleEndPeriod":   js.FuncOf(CalendsCreateDoubleEndPeriod),

		"date":     js.FuncOf(CalendsDate),
		"duration": js.FuncOf(CalendsDuration),
		"endDate":  js.FuncOf(CalendsEndDate),

		"string":     js.FuncOf(CalendsString),
		"decodeText": js.FuncOf(CalendsDecodeText),
		"encodeText": js.FuncOf(CalendsEncodeText),
		"decodeJson": js.FuncOf(CalendsDecodeJson),
		"encodeJson": js.FuncOf(CalendsEncodeJson),

		"registerPanicHandler": js.FuncOf(CalendsRegisterPanicHandler),

		// Calends Offsets
		"addString": js.FuncOf(CalendsAddString),
		"addInt64":  js.FuncOf(CalendsAddInt64),
		"addDouble": js.FuncOf(CalendsAddDouble),

		"subtractString": js.FuncOf(CalendsSubtractString),
		"subtractInt64":  js.FuncOf(CalendsSubtractInt64),
		"subtractDouble": js.FuncOf(CalendsSubtractDouble),

		"addFromEndString": js.FuncOf(CalendsAddFromEndString),
		"addFromEndInt64":  js.FuncOf(CalendsAddFromEndInt64),
		"addFromEndDouble": js.FuncOf(CalendsAddFromEndDouble),

		"subtractFromEndString": js.FuncOf(CalendsSubtractFromEndString),
		"subtractFromEndInt64":  js.FuncOf(CalendsSubtractFromEndInt64),
		"subtractFromEndDouble": js.FuncOf(CalendsSubtractFromEndDouble),

		"nextString": js.FuncOf(CalendsNextString),
		"nextInt64":  js.FuncOf(CalendsNextInt64),
		"nextDouble": js.FuncOf(CalendsNextDouble),

		"previousString": js.FuncOf(CalendsPreviousString),
		"previousInt64":  js.FuncOf(CalendsPreviousInt64),
		"previousDouble": js.FuncOf(CalendsPreviousDouble),

		"withDateString": js.FuncOf(CalendsWithDateString),
		"withDateInt64":  js.FuncOf(CalendsWithDateInt64),
		"withDateDouble": js.FuncOf(CalendsWithDateDouble),

		"withEndDateString": js.FuncOf(CalendsWithEndDateString),
		"withEndDateInt64":  js.FuncOf(CalendsWithEndDateInt64),
		"withEndDateDouble": js.FuncOf(CalendsWithEndDateDouble),

		"withDurationString": js.FuncOf(CalendsWithDurationString),
		"withDurationInt64":  js.FuncOf(CalendsWithDurationInt64),
		"withDurationDouble": js.FuncOf(CalendsWithDurationDouble),

		"withDurationFromEndString": js.FuncOf(CalendsWithDurationFromEndString),
		"withDurationFromEndInt64":  js.FuncOf(CalendsWithDurationFromEndInt64),
		"withDurationFromEndDouble": js.FuncOf(CalendsWithDurationFromEndDouble),

		"merge":     js.FuncOf(CalendsMerge),
		"intersect": js.FuncOf(CalendsIntersect),
		"gap":       js.FuncOf(CalendsGap),

		// Calends Comparisons
		"difference": js.FuncOf(CalendsDifference),

		"compare": js.FuncOf(CalendsCompare),
		"isSame":  js.FuncOf(CalendsIsSame),

		"isShorter":      js.FuncOf(CalendsIsShorter),
		"isSameDuration": js.FuncOf(CalendsIsSameDuration),
		"isLonger":       js.FuncOf(CalendsIsLonger),

		"isBefore": js.FuncOf(CalendsIsBefore),
		"isDuring": js.FuncOf(CalendsIsDuring),
		"isAfter":  js.FuncOf(CalendsIsAfter),

		"startsBefore": js.FuncOf(CalendsStartsBefore),
		"startsDuring": js.FuncOf(CalendsStartsDuring),
		"startsAfter":  js.FuncOf(CalendsStartsAfter),

		"endsBefore": js.FuncOf(CalendsEndsBefore),
		"endsDuring": js.FuncOf(CalendsEndsDuring),
		"endsAfter":  js.FuncOf(CalendsEndsAfter),

		"contains": js.FuncOf(CalendsContains),
		"overlaps": js.FuncOf(CalendsOverlaps),
		"abuts":    js.FuncOf(CalendsAbuts),

		// Calendar Definition
		"calendarRegister":       js.FuncOf(CalendsCalendarRegister),
		"calendarUnregister":     js.FuncOf(CalendsCalendarUnregister),
		"calendarRegistered":     js.FuncOf(CalendsCalendarRegistered),
		"calendarListRegistered": js.FuncOf(CalendsCalendarListRegistered),

		// TAI64Time
		"taiAdd": js.FuncOf(Tai64TimeAdd),
		"taiSub": js.FuncOf(Tai64TimeSub),

		"taiString":     js.FuncOf(Tai64TimeString),
		"taiHexString":  js.FuncOf(Tai64TimeHexString),
		"taiDouble":     js.FuncOf(Tai64TimeDouble),
		"taiEncodeText": js.FuncOf(Tai64TimeEncodeText),

		"taiFromString":    js.FuncOf(Tai64TimeFromString),
		"taiFromHexString": js.FuncOf(Tai64TimeFromHexString),
		"taiFromDouble":    js.FuncOf(Tai64TimeFromDouble),
		"taiDecodeText":    js.FuncOf(Tai64TimeDecodeText),

		"taiFromUtc": js.FuncOf(Tai64TimeUtcToTai),
		"taiToUtc":   js.FuncOf(Tai64TimeTaiToUtc),

		// Cleanup
		"stop": js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			quit <- struct{}{}
			return js.Undefined()
		}),
	}))

	<-quit

	js.Global().Set(`CalendsFuncs`, js.Undefined())
}
