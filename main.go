package main

import (
	"MobieFund/config"
	"MobieFund/model"
	"MobieFund/router"
	"MobieFund/util"
	"fmt"
)

func main() {
	config.NewConfig()
	model.Setup()
	//
	r := router.NewRouter()
	fund := util.Search("161725")
	fmt.Println(fund)
	r.Run("localhost:8080")

}
