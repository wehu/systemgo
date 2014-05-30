package systemgo

// thread

type Cont struct {
	ch chan bool
}

var scheduler_ch   = make(chan *Cont, InitAllocSize)
var cont_q    = make(map[*Cont]*Cont, InitAllocSize)

var running = false

func GenCont() *Cont {
	return &Cont{make(chan bool)}
}

func Sleep(c *Cont) {
	scheduler_ch <- nil
	<- c.ch
}

func Wakeup(c *Cont) {
	if running {
		scheduler_ch <- c
	} else {
		cont_q[c] = c
	}
}

func terminated() {
	scheduler_ch <- nil
}

func resume(c *Cont) {
	c.ch <- true
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
	cont_q[c] = c
	go func() {
		<- c.ch
		body()
		terminated()
	}()
}

func Schedule() {
	running = true
	for g := true; g; {
		select {
			case c := <- scheduler_ch : cont_q[c] = c
			default : g = false
		}
	}
	for len(cont_q) > 0 {
		ql := len(cont_q)
		for _, c := range cont_q {
			resume(c)
		}
		cont_q = make(map[*Cont]*Cont, InitAllocSize)
		for i := 0; i < ql; {
			c := <- scheduler_ch
			if c == nil {
				i ++
			} else {
				cont_q[c] = c
			}
		}
	}
	running = false
}

