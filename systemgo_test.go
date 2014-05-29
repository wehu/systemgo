package systemgo

import "testing"

func TestSystemGo(t *testing.T) {
	Run(-1, func(){
		Always(On(Time(1)), func(){
			Write("aaa", 1)
		})
		Always(On(Signal("aaa")), func(){
			Info("eee %d", ReadB("aaa"))
		})
		Initial(func(){
			Info("hhh")
			Wait(Event("aaa"), Time(1))
			Info("bbb")
			Wait(Time(1))
			Info("ccc")
			WriteB("aaa", 1)
			Info("%d", Read("aaa"))
			Wait(Time(10))
			Finish()
		})
		Initial(func(){
			Info("ddd")
			Event("aaa").Notify()
		})
	})
}