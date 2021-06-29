package base_service

import (
	"paymentservices/api/domain/button_domain"
	services2 "paymentservices/api/services"
	"paymentservices/api/services/applePay_service"
	"paymentservices/api/services/googlePay_service"
	"paymentservices/api/services/paypal_service"
	"paymentservices/api/services/stripe_service"
)

var services = []string{"ApplePay", "GooglePay", "Stripe", "PayPal"}

func GetButtons(input button_domain.ButtonRequest) ([]button_domain.Button, *button_domain.ButtonError) {
	var buttons []button_domain.Button

	var getButtonsHelper = func(s services2.PaymentServiceInterface) *button_domain.ButtonError {
		request := button_domain.ButtonRequest{
			ProductID: input.ProductID,
		}
		button, err := s.GetButton(request)
		if err != nil {
			return button_domain.NewButtonError()
		}
		buttons = append(buttons, button_domain.Button{
			Link: button.Link,
		})
		return nil
	}

	for _, v := range services {
		switch v {
		case "ApplePay":
			if err := getButtonsHelper(applePay_service.ApplePayService); err != nil {
				return nil, err
			}
		case "GooglePay":
			if err := getButtonsHelper(googlePay_service.GooglePayService); err != nil {
				return nil, err
			}
		case "Stripe":
			if err := getButtonsHelper(stripe_service.StripeService); err != nil {
				return nil, err
			}
		case "PayPal":
			if err := getButtonsHelper(paypal_service.PaypalService); err != nil {
				return nil, err
			}
		default:
			return nil, button_domain.NewButtonError()
		}
	}
	return buttons, nil
}
