package stripe_service

import (
	"paymentservices/api/domain/button_domain"
	"paymentservices/api/providers"
	"paymentservices/api/services"
)

type stripeService struct{}

var StripeService services.PaymentServiceInterface = &stripeService{}

const stripeUrl = "https://api.stripe/%s"

func (s *stripeService) GetButton(input button_domain.ButtonRequest) (*button_domain.Button, *button_domain.ButtonError) {
	request := button_domain.ButtonRequest{
		ProductID: input.ProductID,
	}
	response, err := providers.ServiceProvider.GetButton(request, stripeUrl)
	if err != nil {
		return nil, button_domain.NewButtonError()
	}
	result := button_domain.Button{
		Link: response.Link,
	}
	return &result, nil
}
