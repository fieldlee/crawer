package parser

import (
	"crawer/engine"
	"regexp"
)
const regexCity = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]*)</a>`
func ParseCityList(b []byte) engine.ParseRequest{

	reg := regexp.MustCompile(regexCity)
	matchs := reg.FindAllSubmatch(b,-1)

	parseReq := engine.ParseRequest{}
	for _,m := range  matchs{
		parseReq.Items = append(parseReq.Items,m[2])
		parseReq.Requests = append(parseReq.Requests,engine.Request{
			Url:string(m[1]),
			ParseFun:ParserUser,
		})
	}
	return  parseReq
}