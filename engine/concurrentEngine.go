package engine

import (
	"log"
	"crawer/fetcher"
	"time"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	ConfigureWorkChan(chan Request)
	Submit(Request)
}

func (e *ConcurrentEngine)Run(seeds ...Request)  {

	requests := []Request{}
	requests = append(requests,seeds...)


	in := make(chan Request)
	out := make(chan ParseRequest)

	e.Scheduler.ConfigureWorkChan(in)
	//创建channel
	for i:=0;i< e.WorkerCount ;i++  {
		createWorker(in,out)
	}
	//发送chan
	for _,r := range  requests{
		e.Scheduler.Submit(r)
	}
	//从chan里取值
	for{
		pq := <- out
		for _,item := range pq.Items {
			log.Printf("item:%s\n",item)
		}

		for _,r := range  pq.Requests{
			e.Scheduler.Submit(r)
		}
	}

}


// 处理chan
func createWorker(in chan Request,out chan ParseRequest){
	go func() {
		for  {
			inRequest := <- in

			result , err := workerFetch(inRequest)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var timeLater = time.Tick(200 * time.Millisecond)

func workerFetch(r Request) (ParseRequest,error) {
	<- timeLater
	body,err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Println("err:",err.Error())
		return ParseRequest{} , err
	}
	return r.ParseFun(body),nil
}