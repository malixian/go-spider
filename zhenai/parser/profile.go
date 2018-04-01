package parser

import (

	"regexp"
	"go-spider/engine"
	"go-spider/model"
)

const ageRe = `<td><span class="label">年龄：</span>(.*?)</td>`
const heighRe = `<td><span class="label">身高：</span>(.*?)</td>`
const incomeRe= `<td><span class="label">月收入：</span>(.*?)</td>`
const marryRe = `<td><span class="label">婚况：</span>(.*?)</td>`
const educationRe = `<td><span class="label">学历：</span>(.*?)</td>`
const workLocationRe = `<td><span class="label">工作地：</span>(.*?)</td>`
const idRe = `http://album.zhenai.com/u/([\d]+)`
var ageCom = regexp.MustCompile(ageRe)
var heighCom = regexp.MustCompile(heighRe)
var incomeCom = regexp.MustCompile(incomeRe)
var marryCom = regexp.MustCompile(marryRe)
var educationCom = regexp.MustCompile(educationRe)
var workLocationCom = regexp.MustCompile(workLocationRe)
var idCom = regexp.MustCompile(idRe)



func ParseProfile(contents []byte,url string, name string) engine.ParseResult{
	age := extractString(contents, ageCom)
	heigh := extractString(contents, heighCom)
	income := extractString(contents, incomeCom)
	marry := extractString(contents, marryCom)
	education := extractString(contents, educationCom)
	workLocation := extractString(contents, workLocationCom)
	id := extractString(contents, idCom)
	var profile model.Profile
	profile.Name = name
	profile.Age = age
	profile.Heigh = heigh
	profile.Income = income
	profile.Marry = marry
	profile.Education = education
	profile.WorkLocation = workLocation
	item := engine.Item{url,id,"zhenai",profile}
	res := engine.ParseResult{Items: []engine.Item{item}}
	return res
}

func extractString(contents []byte, re *regexp.Regexp) string{
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	}else{
		return ""
	}
}
