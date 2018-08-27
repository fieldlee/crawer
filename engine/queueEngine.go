package engine

import (
	"log"
	"crawer/fetcher"
)

type QueueEngine struct {
	Scheduler QueueScheduler
	WorkerCount int
}

type QueueScheduler interface {
	ConfigureWorkChan(chan Request)
	Submit(Request)
	WorkReady( chan Request)
	Run()
}

func (q *QueueEngine)Run(seeds ...Request)  {

	out := make(chan ParseRequest)
	q.Scheduler.Run()
	//创建channel
	for i:=0;i< q.WorkerCount ;i++  {
		queueCreateWorker(out,q.Scheduler)
	}
	//发送chan
	for _,r := range seeds{
		q.Scheduler.Submit(r)
	}
	log.Println("1")
	//从chan里取值
	for{
		pq := <- out
		for _,item := range pq.Items {
			log.Printf("item:%s\n",item)
		}
		for _,r := range pq.Requests{
			q.Scheduler.Submit(r)
		}
	}
}

// 处理chan
func queueCreateWorker(out chan ParseRequest,s QueueScheduler){
	in := make(chan Request)
	go func() {
		for  {
			s.WorkReady(in)
			inRequest := <- in
			result , err := queueWorkerFetch(inRequest)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func queueWorkerFetch(r Request) (ParseRequest,error) {
	body,err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Println("err:",err.Error())
		return ParseRequest{} , err
	}
	return r.ParseFun(body),nil
}