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

var signals = make(map[string]*SignalT, InitAllocSize)
var ssl = new(sync.Mutex)

func Signal(name string) *SignalT {
	s, ok := signals[name]
	if ok {
		return s
	}
	ssl.Lock()
	defer ssl.Unlock()
	s, ok = signals[name]
	if !ok {
	  s = &SignalT{0, 0, EventT{name}}
	  signals[name] = s
	}
	return s
}

func Read(name string) int {
	sv := Signal(name)
	//ssl.Lock()
	s := sv.old_var
	//ssl.Unlock()
	return s
}

func ReadB(name string) int {
	Wait(Time(0))
	return Read(name)
}

func Write(name string, v int) {
	sv := Signal(name)
	//ssl.Lock()
	sv.new_var = v
	//ssl.Unlock()
	if sv.old_var != sv.new_var {
		sv.Notify()
	}
}

func WriteB(name string, v int) {
	Write(name, v)
	Wait(Time(0))
}

func SyncSignals() {
	for _, s := range signals {
		s.old_var = s.new_var
	}
}
