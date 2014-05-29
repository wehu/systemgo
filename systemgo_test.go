package systemgo

import "testing"

func TestSystemgo(t *testing.T) {
	Run(-1, func(){
		//Always(On(Time(1)), func(){
		//	Info("aaaaaaaa")
		//})
		Initial(func(){
			Info("hhh")
			Wait(Event("aaa"), Time(1))
			Info("bbb")
			Wait(Time(1))
			Info("ccc")
			WriteB("aaa", 1)
			Info("%d", Read("aaa"))
		})
		Initial(func(){
			Info("ddd")
			Event("aaa").Notify()
		})
	})
}