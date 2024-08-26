package main

import (
	"fmt"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/paytm/grace.v1"

	userCmd "tablelink/internal/cmd"
)

func main() {
	initModule()
	initGRPCServer()

	//set handler
	router := httprouter.New()
	initRoute(router)
	grace.Serve(":8080", router)

	fmt.Println("Init Setup")

}

func initGRPCServer() {
	//
}

func initRoute(router *httprouter.Router) {
	//get all
	router.GET("/users/user", userCmd.HTTPDelivery.HandleGetAllUser)

	//login
	router.POST("/auth/login", userCmd.HTTPDelivery.HandleLogin)

	//add user
	router.POST("/users/user", userCmd.HTTPDelivery.HandleGetAllUser)

	//update user
	router.PUT("/users/user", userCmd.HTTPDelivery.HandleLogin)

	//delete user
	router.DELETE("/users/user/", userCmd.HTTPDelivery.HandleLogin)

}

func initModule() {
	//init module initialze
	userCmd.Initialize()
}
