package paypal_service

import (
	"paymentservices/api/domain/button_domain"
	"paymentservices/api/providers"
	"paymentservices/api/services"
)

type paypalService struct{}

var PaypalService services.PaymentServiceInterface = &paypalService{}

const paypalUrl = "https://api.paypal/%s"

func (s *paypalService) GetButton(input button_domain.ButtonRequest) (*button_domain.Button, *button_domain.ButtonError) {
	request := button_domain.ButtonRequest{
		ProductID: input.ProductID,
	}
	response, err := providers.ServiceProvider.GetButton(request, paypalUrl)
	if err != nil {
		return nil, button_domain.NewButtonError()
	}
	result := button_domain.Button{
		Link: response.Link,
	}
	return &result, nil
}
