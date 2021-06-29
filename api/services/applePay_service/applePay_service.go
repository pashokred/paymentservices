package applePay_service

import (
	"paymentservices/api/domain/button_domain"
	"paymentservices/api/providers"
	"paymentservices/api/services"
)

type applePayService struct{}

var ApplePayService services.PaymentServiceInterface = &applePayService{}

const applePayUrl = "https://api.applepay/%s"

func (s *applePayService) GetButton(input button_domain.ButtonRequest) (*button_domain.Button, *button_domain.ButtonError) {
	request := button_domain.ButtonRequest{
		ProductID: input.ProductID,
	}
	response, err := providers.ServiceProvider.GetButton(request, applePayUrl)
	if err != nil {
		return nil, button_domain.NewButtonError()
	}
	result := button_domain.Button{
		Link: response.Link,
	}
	return &result, nil
}
