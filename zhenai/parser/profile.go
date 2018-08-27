package parser

import (
	"crawer/engine"
	"regexp"
	"crawer/model"
	"strconv"
	"fmt"
)

var nameRe  = regexp.MustCompile(`<h1 class="ceiling-name ib fl fs24 lh32 blue">([^<]+)</h1>`)
var ageRe  = regexp.MustCompile(`<td><span class="label">年龄：</span>([0-9]+)岁</td>`)
var heightRe  = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var weightRe  = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
var incomeRe  = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var marrageRe  = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var workplaceRe  = regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+)</td>`)
var workerRe  = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var hometownRe  = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)


func Profile(b []byte,name string) engine.ParseRequest  {
	profile := model.Profile{}

	profile.Name = name
	match := ageRe.FindSubmatch(b)
	if match != nil {
		//log.Println(string(match[1]))
		age , err := strconv.Atoi(string(match[1]))
		if err != nil {
			profile.Age = 0
		}else{
			profile.Age = age
		}

	}else{
		profile.Age = 0
	}

	match = heightRe.FindSubmatch(b)
	if match != nil {
		//log.Println(string(match[1]))
		age , err := strconv.Atoi(string(match[1]))
		if err != nil {
			profile.Height = 0
		}else{
			profile.Height = age
		}

	}else{
		profile.Height = 0
	}


	match = weightRe.FindSubmatch(b)
	if match != nil {
		//log.Println(string(match[1]))
		age , err := strconv.Atoi(string(match[1]))
		if err != nil {
			profile.Weight = 0
		}else{
			profile.Weight = age
		}

	}else{
		profile.Weight = 0
	}

	match = incomeRe.FindSubmatch(b)
	if match != nil {
		//log.Println(string(match[1]))
		profile.InCome = string(match[1])
	}else{
		profile.InCome = ""
	}

	match = marrageRe.FindSubmatch(b)
	if match != nil {
		//log.Println(string(match[1]))
		profile.Marrage = string(match[1])
	}else{
		profile.Marrage = ""
	}

	match = workplaceRe.FindSubmatch(b)
	if match != nil {
		profile.WorkPlace = string(match[1])
	}else{
		profile.WorkPlace = ""
	}


	match = workerRe.FindSubmatch(b)
	if match != nil {
		profile.Worker = string(match[1])
	}else{
		profile.Worker = ""
	}


	match = hometownRe.FindSubmatch(b)
	if match != nil {
		profile.HomeTown = string(match[1])
	}else{
		profile.HomeTown = ""
	}
	parseReq := engine.ParseRequest{}

	parseReq.Items = append(parseReq.Items,profile)

	//parseReq.Requests = append(parseReq.Requests,engine.Request{
	//	Url:string(m[1]),
	//	ParseFun:engine.NilParseFunc,
	//})

	fmt.Println(profile)

	return  parseReq
}
