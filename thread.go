package systemgo

import (
	"sync"
)

// thread

var tl = new(sync.Mutex)

type Cont struct {
	ch chan bool
}

var scheduler_ch   = make(chan bool, 1000)
var cont_q    = make(map[*Cont]*Cont)

func GenCont() *Cont {
	return &Cont{make(chan bool)}
}

func Sleep(c *Cont) {
	scheduler_ch <- true
	<- c.ch
}

func Wakeup(c *Cont) {
	tl.Lock()
	cont_q[c] = c
	tl.Unlock()
}

func terminated() {
	scheduler_ch <- true
}

func resume(c *Cont) {
	c.ch <- true
	tl.Lock()
	delete(cont_q, c)
	tl.Unlock()
}

func Wait(es ...EventIntf) {
	c := GenCont()
	var cb func()
	cb = func () {
		for _, e := range es {
			switch e.(type) {
				case TimeDT: addTime(e.(TimeDT)).UnSubscribeCallback(&cb)
				default: e.UnSubscribeCallback(&cb)
			}
			Wakeup(c)
		}
	}
	for _, e := range es {
		switch e.(type) {
			case TimeDT: addTime(e.(TimeDT)).Subscribe(&cb)
			default: e.Subscribe(&cb)
		}
	}
	Sleep(c)
}

func New(body func()) {
	c := GenCont()
	Wakeup(c)
	go func() {
		<- c.ch
		body()
		terminated()
	}()
}

func Schedule() {
	for len(cont_q) > 0 {
		q := make(map[*Cont]*Cont)
		for _, c := range cont_q {
			q[c] = c
		}
		ql := len(q)
		for _, c := range q {
			resume(c)
		}
		for i := 0; i < ql; i ++ {
			<- scheduler_ch
		}
	}
}

