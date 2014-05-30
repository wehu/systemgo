package systemgo

import (
	"fmt"
	"sync"
)

// sim time

var current_simtime = 0

var stop = true

type TimeDT struct {
	delay int
	EventT
}

type TimeT struct {
	EventT
}

var simtimes = make(map[int]TimeT, InitAllocSize)

var sl = new(sync.Mutex)

func Time(t int) TimeDT {
	te := TimeDT{t, EventT{fmt.Sprintf("simtimed%d", t)}}
	return te
}

func addTime(t TimeDT) TimeT {
	nt := current_simtime + t.delay
	te := TimeT{EventT{fmt.Sprintf("simtime%d", nt)}}
	sl.Lock()
	defer sl.Unlock()
	simtimes[nt] = te
	return te
}

func CurrentTime() int {
	return current_simtime
}

func Simulate(delta int) {
	stop = false
	max_time := current_simtime + delta
	for {
		Schedule()
		SyncSignals()
		if stop || len(simtimes) == 0 || (delta >= 0 && max_time <= current_simtime) {
			break
		} else {
			mt := -1
			for t, _ := range simtimes {
				if t < current_simtime {
					delete(simtimes, t)
				} else {
					if mt == -1 || t < mt {
						mt = t
					}
				}
			}
			te := simtimes[mt]
			te.Notify()
			te.UnSubscribe()
			delete(simtimes, mt)
			current_simtime = mt
		}
	}
}

func Finish() {
	stop = true
}
