package googlePay_service

import (
	"paymentservices/api/domain/button_domain"
	"paymentservices/api/providers"
	"paymentservices/api/services"
)

type googlePayService struct{}

var GooglePayService services.PaymentServiceInterface = &googlePayService{}

const googlePayUrl = "https://api.googlepay/%s"

func (s *googlePayService) GetButton(input button_domain.ButtonRequest) (*button_domain.Button, *button_domain.ButtonError) {
	request := button_domain.ButtonRequest{
		ProductID: input.ProductID,
	}
	response, err := providers.ServiceProvider.GetButton(request, googlePayUrl)
	if err != nil {
		return nil, button_domain.NewButtonError()
	}
	result := button_domain.Button{
		Link: response.Link,
	}
	return &result, nil
}
