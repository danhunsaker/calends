package main

//char* Calends_version;
//typedef void (*Calends_panic_handler) (char*);
//
//static inline void wrap_Calends_panic_handler(Calends_panic_handler f, char* message) {
//  f(message);
//}
import "C"

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/danhunsaker/calends"
)

var SHARDS = uint64(64)

type ConcurrentMap []*ConcurrentMapShard

type ConcurrentMapShard struct {
	items map[uint64]interface{}
	sync.RWMutex
}

func NewConcurrentMap() ConcurrentMap {
	m := make(ConcurrentMap, SHARDS)
	for i := uint64(0); i < SHARDS; i++ {
		m[i] = &ConcurrentMapShard{items: make(map[uint64]interface{})}
	}
	return m
}

func (m ConcurrentMap) getShard(key uint64) *ConcurrentMapShard {
	return m[key%SHARDS]
}

func (m ConcurrentMap) Store(key uint64, value interface{}) {
	shard := m.getShard(key)
	shard.Lock()
	shard.items[key] = value
	shard.Unlock()
}

func (m ConcurrentMap) Load(key uint64) (interface{}, bool) {
	shard := m.getShard(key)
	shard.RLock()
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}

func (m ConcurrentMap) Length() (out uint64) {
	for _, shard := range m {
		shard.RLock()
		out += uint64(len(shard.items))
		shard.RUnlock()
	}
	return
}

func (m ConcurrentMap) All() (out []interface{}) {
	for _, shard := range m {
		shard.RLock()
		for _, item := range shard.items {
			out = append(out, item)
		}
		shard.RUnlock()
	}
	return
}

func (m ConcurrentMap) Delete(key uint64) {
	shard := m.getShard(key)
	shard.Lock()
	delete(shard.items, key)
	shard.Unlock()
}

type IdGenerator struct {
	id uint64
}

func (generator *IdGenerator) Id() uint64 {
	return atomic.AddUint64(&generator.id, 1)
}

var panicHandlers ConcurrentMap
var nextPanHandle IdGenerator
var instances ConcurrentMap
var nextInst IdGenerator

func init() {
	C.Calends_version = C.CString(calends.Version)
	panicHandlers = NewConcurrentMap()
	instances = NewConcurrentMap()
}

func instNum(c calends.Calends) C.ulonglong {
	p := nextInst.Id()
	instances.Store(p, &c)
	return C.ulonglong(p)
}

func instGet(p C.ulonglong) *calends.Calends {
	defer handlePanic()
	val, ok := instances.Load(uint64(p))
	if !ok {
		panic(fmt.Sprintf("Calends object #%d doesn't exist", uint64(p)))
	}
	return val.(*calends.Calends)
}

func calends_create(stamp interface{}, calendar, format *C.char) C.ulonglong {
	defer handlePanic()
	c, err := calends.Create(stamp, C.GoString(calendar), C.GoString(format))
	if err != nil {
		panic(err)
	}
	return instNum(c)
}

//export Calends_release
func Calends_release(p C.ulonglong) {
	instances.Delete(uint64(p))
}

//export Calends_create_string
func Calends_create_string(stamp, calendar, format *C.char) C.ulonglong {
	return calends_create(C.GoString(stamp), calendar, format)
}

//export Calends_create_string_range
func Calends_create_string_range(start, end, calendar, format *C.char) C.ulonglong {
	return calends_create(map[string]interface{}{
		"start": C.GoString(start),
		"end":   C.GoString(end),
	}, calendar, format)
}

//export Calends_create_string_start_period
func Calends_create_string_start_period(start, duration, calendar, format *C.char) C.ulonglong {
	return calends_create(map[string]interface{}{
		"start":    C.GoString(start),
		"duration": C.GoString(duration),
	}, calendar, format)
}

//export Calends_create_string_end_period
func Calends_create_string_end_period(duration, end, calendar, format *C.char) C.ulonglong {
	return calends_create(map[string]interface{}{
		"duration": C.GoString(duration),
		"end":      C.GoString(end),
	}, calendar, format)
}

//export Calends_create_long_long
func Calends_create_long_long(stamp C.longlong, calendar, format *C.char) C.ulonglong {
	return calends_create(int(stamp), calendar, format)
}

//export Calends_create_long_long_range
func Calends_create_long_long_range(start, end C.longlong, calendar, format *C.char) C.ulonglong {
	return calends_create(map[string]interface{}{
		"start": int(start),
		"end":   int(end),
	}, calendar, format)
}

//export Calends_create_long_long_start_period
func Calends_create_long_long_start_period(start, duration C.longlong, calendar, format *C.char) C.ulonglong {
	return calends_create(map[string]interface{}{
		"start":    int(start),
		"duration": int(duration),
	}, calendar, format)
}

//export Calends_create_long_long_end_period
func Calends_create_long_long_end_period(duration, end C.longlong, calendar, format *C.char) C.ulonglong {
	return calends_create(map[string]interface{}{
		"duration": int(duration),
		"end":      int(end),
	}, calendar, format)
}

//export Calends_create_double
func Calends_create_double(stamp C.double, calendar, format *C.char) C.ulonglong {
	return calends_create(float64(stamp), calendar, format)
}

//export Calends_create_double_range
func Calends_create_double_range(start, end C.double, calendar, format *C.char) C.ulonglong {
	return calends_create(map[string]interface{}{
		"start": float64(start),
		"end":   float64(end),
	}, calendar, format)
}

//export Calends_create_double_start_period
func Calends_create_double_start_period(start, duration C.double, calendar, format *C.char) C.ulonglong {
	return calends_create(map[string]interface{}{
		"start":    float64(start),
		"duration": float64(duration),
	}, calendar, format)
}

//export Calends_create_double_end_period
func Calends_create_double_end_period(duration, end C.double, calendar, format *C.char) C.ulonglong {
	return calends_create(map[string]interface{}{
		"duration": float64(duration),
		"end":      float64(end),
	}, calendar, format)
}

//export Calends_date
func Calends_date(p C.ulonglong, calendar, format *C.char) *C.char {
	defer handlePanic()
	c := instGet(p)
	out, err := c.Date(C.GoString(calendar), C.GoString(format))
	if err != nil {
		panic(err)
	}
	return C.CString(out)
}

//export Calends_duration
func Calends_duration(p C.ulonglong) *C.char {
	c := instGet(p)
	return C.CString(c.Duration().String())
}

//export Calends_end_date
func Calends_end_date(p C.ulonglong, calendar, format *C.char) *C.char {
	defer handlePanic()
	c := instGet(p)
	out, err := c.EndDate(C.GoString(calendar), C.GoString(format))
	if err != nil {
		panic(err)
	}
	return C.CString(out)
}

//export Calends_string
func Calends_string(p C.ulonglong) *C.char {
	c := instGet(p)
	return C.CString(c.String())
}

//export Calends_encode_text
func Calends_encode_text(p C.ulonglong) *C.char {
	defer handlePanic()
	c := instGet(p)
	out, err := c.MarshalText()
	if err != nil {
		panic(err)
	}
	return C.CString(string(out))
}

//export Calends_decode_text
func Calends_decode_text(in *C.char) C.ulonglong {
	defer handlePanic()
	c := calends.Calends{}
	err := c.UnmarshalText([]byte(C.GoString(in)))
	if err != nil {
		panic(err)
	}
	return instNum(c)
}

//export Calends_encode_json
func Calends_encode_json(p C.ulonglong) *C.char {
	defer handlePanic()
	c := instGet(p)
	out, err := c.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return C.CString(string(out))
}

//export Calends_decode_json
func Calends_decode_json(in *C.char) C.ulonglong {
	defer handlePanic()
	c := calends.Calends{}
	err := c.UnmarshalJSON([]byte(C.GoString(in)))
	if err != nil {
		panic(err)
	}
	return instNum(c)
}

//export Calends_register_panic_handler
func Calends_register_panic_handler(handler C.Calends_panic_handler) {
	id := nextPanHandle.Id()
	panicHandlers.Store(id, handler)
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
		switch r.(type) {
		case string:
			err = r.(string)
		case error:
			err = r.(error).Error()
		default:
			err = fmt.Sprintf("%#v", r)
		}

		// Call the handler(s)!
		for _, handler := range panicHandlers.All() {
			C.wrap_Calends_panic_handler(handler.(C.Calends_panic_handler), C.CString(err))
		}
	}
}

func main() {
	panic(fmt.Sprintf("Calends %s\nThis shouldn't ever be called!", C.GoString(C.Calends_version)))
}
