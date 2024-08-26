package http

import (
	uCustomer "tablelink/internal/usecase/user"
)

type Delivery struct {
	userUC uCustomer.Usecase
}

func NewUserHTTP(userUC uCustomer.Usecase) Delivery {
	return Delivery{userUC}

}
