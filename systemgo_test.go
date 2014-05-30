package systemgo

import "testing"
import "runtime/pprof"
import "os"

func TestSystemGo(t *testing.T) {
	f, _ := os.Create("cpu.prof")
	pprof.StartCPUProfile(f)
	Run(-1, func() {
		Always(On(Time(1)), func() {
			Write("aaa", Read("aaa")+1)
		})
		for i := 0; i < 10; i++ {
			Always(On(Signal("aaa")), func() {
				Info("eee %d", ReadB("aaa"))
			})
		}
		Initial(func() {
			Info("hhh")
			Wait(Event("aaa"), Time(1))
			Info("bbb")
			Wait(Time(1))
			Info("ccc")
			WriteB("aaa", 1)
			Info("%d", Read("aaa"))
			Wait(Time(100000))
			Finish()
		})
		Initial(func() {
			Info("ddd")
			Event("aaa").Notify()
		})
	})
	pprof.StopCPUProfile()
}
