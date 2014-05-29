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

var events = make(map[EventIntf]map[*func()]*func())

var el = new(sync.Mutex)

type EventT struct {
	name string
}

func (e EventT) GetName() string {
	return e.name
}

func (e EventT) initEventCallbacks() {
	el.Lock()
	_, ok := events[e]
	if !ok {
		events[e] = make(map[*func()]*func())
	}
	el.Unlock()
}

func (e EventT) Subscribe(cb *func()) *func() {
	e.initEventCallbacks()
	el.Lock()
	events[e][cb] = cb
	el.Unlock()
	return cb
}

func (e EventT) Notify() {
	e.initEventCallbacks()
	for _, cb := range events[e] {
		(*cb)()
	}
}

func (e EventT) UnSubscribe() {
	el.Lock()
	delete(events, e)
	el.Unlock()
}

func (e EventT) UnSubscribeCallback(cb *func()) {
	e.initEventCallbacks()
	el.Lock()
	delete(events[e], cb)
	el.Unlock()
}

func Event(e string) EventT {
	return EventT{e}
}

