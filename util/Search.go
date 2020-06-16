package util

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

/**
 * @Author 王睿
 * @Description //TODO 净值估算
 * @Date 20:35 2020/6/16
 * @Param
 * @return
 **/
type evaluate struct {
	date       string `json:"date"`
	value      string `json:"value"`
	add        string `json:"add"`
	addPercent string `json:"add_percent"`
}

/**
 * @Author 王睿
 * @Description //TODO 单位净值
 * @Date 20:36 2020/6/16
 * @Param
 * @return
 **/
type trueValue struct {
	date  string `json:"date"`
	value string `json:"value"`
	add   string `json:"add"`
}

type found struct {
	id   string `json:"id"`
	name string `json:"name"`
	evaluate
	trueValue
}

var fund = found{}

func Search(fundId string) found {

	fund.id = fundId

	c := colly.NewCollector(
		colly.AllowedDomains("fund.eastmoney.com"),
	)

	//基金名称
	c.OnHTML(".fundDetail-tit", func(e *colly.HTMLElement) {
		fund.name = e.Text
		// Print link
		//fmt.Printf("Link found : %q \n", e.Text)
	})

	//基金净值估算 date
	c.OnHTML("#gz_gztime", func(e *colly.HTMLElement) {

		str := e.Text
		str = strings.ReplaceAll(str, "(", "")
		str = strings.ReplaceAll(str, ")", "")
		str = "20" + str
		fund.evaluate.date = str

		//fmt.Printf("Link found : %v \n", str)
	})

	//基金净值估算 value
	c.OnHTML("#gz_gsz", func(e *colly.HTMLElement) {

		fund.evaluate.value = e.Text
		//fmt.Printf("Link found : %v \n", e.Text)
	})

	//基金净值估算
	c.OnHTML("#gz_gszze", func(e *colly.HTMLElement) {

		fund.evaluate.add = e.Text
		//fmt.Printf("Link found : %v \n", e.Text)
	})

	//基金净值估算
	c.OnHTML("#gz_gszzl", func(e *colly.HTMLElement) {

		fund.evaluate.addPercent = e.Text
		//fmt.Printf("Link found : %v \n", e.Text)
	})

	//基金净值
	c.OnHTML(".dataItem02 p", func(e *colly.HTMLElement) {

		str := e.Text
		str = strings.ReplaceAll(str, "单位净值 (", "")
		str = strings.ReplaceAll(str, ")", "")
		fund.trueValue.date = str
		//fmt.Printf("Link found : %v \n", str)
	})

	//基金净值
	c.OnHTML(".dataItem02 .dataNums span", func(e *colly.HTMLElement) {

		index := strings.Index(e.Text, "%")
		if index > -1 {
			fund.trueValue.add = e.Text
		} else {
			fund.trueValue.value = e.Text
		}

		//fmt.Printf("Link found : %v \n", e.Text)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("http://fund.eastmoney.com/" + fundId + ".html")

	return fund
}
