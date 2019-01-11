// Package libcalends exports the calends library for use by C/C++, and thereby
// other programming languages.
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

func init() {
	C.Calends_version = C.CString(calends.Version)
	panicHandlers = newConcurrentMap()
	instances = newConcurrentMap()
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
