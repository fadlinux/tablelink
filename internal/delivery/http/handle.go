package http

import (
	"context"
	"net/http"

	mUser "tablelink/internal/model/users"

	"github.com/julienschmidt/httprouter"
)

func (d Delivery) HandleLogin(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	ctx := context.Background()
	var exist bool

	requestLogin := mUser.RequestLogin{
		Email:    req.FormValue(`name`),
		Password: req.FormValue("password"),
	}

	exist, _ = d.userUC.Login(ctx, requestLogin)
	if !exist {

	}

	//response code
}

func (d Delivery) HandleGetAllUser(w http.ResponseWriter, req *http.Request, params httprouter.Params) {

	//response code
}

func (d Delivery) handleUpdateUser(w http.ResponseWriter, req *http.Request, params httprouter.Params) {

	//response code
}

func (d Delivery) HandleDeleteUser(w http.ResponseWriter, req *http.Request, params httprouter.Params) {

	//response code
}
