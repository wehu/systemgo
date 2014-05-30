package systemgo

import (
	"sync"
)

// event

type EventIntf interface {
	GetName() string
	initEventCallbacks()
	Subscribe(*func()) *func()
	Notify()
	UnSubscribe()
	UnSubscribeCallback(*func())
}

var events = make(map[EventIntf]map[*func()]*func(), InitAllocSize)

var el = new(sync.RWMutex)

type EventT struct {
	name string
}

func (e EventT) GetName() string {
	return e.name
}

func (e EventT) initEventCallbacks() {
	el.Lock()
	defer el.Unlock()
	_, ok := events[e]
	if !ok {
		events[e] = make(map[*func()]*func(), 10)
	}
}

func (e EventT) Subscribe(cb *func()) *func() {
	e.initEventCallbacks()
	el.Lock()
	defer el.Unlock()
	events[e][cb] = cb
	return cb
}

func (e EventT) Notify() {
	e.initEventCallbacks()
	el.RLock()
	cbs := events[e]
	el.RUnlock()
	for _, cb := range cbs {
		(*cb)()
	}
}

func (e EventT) UnSubscribe() {
	el.Lock()
	defer el.Unlock()
	delete(events, e)
}

func (e EventT) UnSubscribeCallback(cb *func()) {
	e.initEventCallbacks()
	el.Lock()
	defer el.Unlock()
	delete(events[e], cb)
}

func Event(e string) EventT {
	return EventT{e}
}
