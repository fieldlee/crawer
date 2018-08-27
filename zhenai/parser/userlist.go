package parser

import (
	"crawer/engine"
	"regexp"
)

const userRegex  = `<th><a href="(http://album.zhenai.com/u/[0-9a-z]+)" target="_blank">([^<]+)</a></th>`

const nextPageRE = `<a class="next-page" href="(http://www.zhenai.com/zhenghun/[a-z]+/[0-9]+)">下一页</a>`


func ParserUser(b []byte)engine.ParseRequest  {
	reg := regexp.MustCompile(userRegex)
	matchs := reg.FindAllSubmatch(b,-1)

	parseReq := engine.ParseRequest{}
	for _,m := range  matchs{
		parseReq.Items = append(parseReq.Items,m[2])
		name := string(m[2])
		parseReq.Requests = append(parseReq.Requests,engine.Request{
			Url:string(m[1]),
			ParseFun: func(c []byte) engine.ParseRequest {
				return Profile(c,name)
			},
		})
	}

	//读取下一页
	nextReg := regexp.MustCompile(nextPageRE)
	match := nextReg.FindSubmatch(b)
	if match != nil {
		parseReq.Items = append(parseReq.Items,string(match[1]))
		parseReq.Requests = append(parseReq.Requests,engine.Request{
			Url:string(match[1]),
			ParseFun:ParserUser,
		})
	}
	return  parseReq
}
