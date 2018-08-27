package main

import (
	"crawer/engine"
	"crawer/scheduler"
	"crawer/zhenai/parser"
)

func main()  {
	//	并发
	//e := engine.ConcurrentEngine{
	//	WorkerCount:100,
	//	Scheduler:&scheduler.ConcurrentScheduler{},
	//}
	//
	//e.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun",
	//	ParseFun:parser.ParseCityList,
	//})

	//	单线程

	//engine.SimpleEngine{}.Run(
	//	engine.Request{
	//			Url:"http://www.zhenai.com/zhenghun",
	//			ParseFun:parser.ParseCityList,
	//		})

	//  队列

	e := engine.QueueEngine{
		WorkerCount:10,
		Scheduler:&scheduler.QueuedScheduler{},
	}

	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParseFun:parser.ParseCityList,
	})

}

