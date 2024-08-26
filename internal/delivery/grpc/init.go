package grpc

import (
	uCustomer "tablelink/internal/usecase/user"
)

type Delivery struct {
	userUC uCustomer.Usecase
}

func NewUserGRPC(userUC uCustomer.Usecase) Delivery {
	return Delivery{userUC}

}
