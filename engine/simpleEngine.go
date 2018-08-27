package engine

import (
	"crawer/fetcher"
	"log"
)

type SimpleEngine struct {}

func (e SimpleEngine)Run(seeds ...Request)  {
	requests := []Request{}
	requests = append(requests,seeds...)

	for len(requests)>0 {
		r := requests[0]
		requests = requests[1:]
		parseRequest, err := e.workerFetch(r)
		if err != nil {
			continue
		}
		requests = append(requests,parseRequest.Requests...)

		for _,item := range parseRequest.Items{
			log.Printf("item:%s\n",item)
		}
	}

}

func (e SimpleEngine) workerFetch(r Request) (ParseRequest,error) {
	body,err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Println("err:",err.Error())
		return ParseRequest{} , err
	}
	return r.ParseFun(body),nil
}
