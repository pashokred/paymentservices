package services

import "paymentservices/api/domain/button_domain"

type PaymentServiceInterface interface {
	GetButton(input button_domain.ButtonRequest) (*button_domain.Button, *button_domain.ButtonError)
}
