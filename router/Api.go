package router

import "MobieFund/controller"

func ApiV1(base string) {

	r := Router.Group("/" + base)

	r.POST("/createuser", controller.CreateUser)
	r.POST("/addfund", controller.AddFund)

}
