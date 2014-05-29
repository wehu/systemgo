package systemgo

import (
	"sync"
)

// signal

type SignalT struct {
	old_var int
	new_var int
	EventT
}

var signals = make(map[string]*SignalT)
var ssl = new(sync.Mutex)

func Signal(name string) *SignalT {
	ssl.Lock()
	s, ok := signals[name]
	if !ok {
	  s = &SignalT{0, 0, EventT{name}}
	  signals[name] = s
	}
	ssl.Unlock()
	return s
}

func Read(name string) int {
	sv := Signal(name)
	ssl.Lock()
	s := sv.old_var
	ssl.Unlock()
	return s
}

func WriteNB(name string, v int) {
	sv := Signal(name)
	ssl.Lock()
	sv.new_var = v
	ssl.Unlock()
}

func WriteB(name string, v int) {
	WriteNB(name, v)
	Wait(Time(0))
}

func SyncSignals() {
	for _, s := range signals {
		s.old_var = s.new_var
	}
}
