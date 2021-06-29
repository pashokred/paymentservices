package stripe_service

import (
	"github.com/stretchr/testify/assert"
	"paymentservices/api/domain/button_domain"
	"paymentservices/api/providers"
	"strconv"
	"testing"
)

var getServiceProviderFunc func(request button_domain.ButtonRequest, serviceURL string) (*button_domain.Button, *button_domain.ButtonError)

type getProviderMock struct{}

const (
	appStoreLink   = "https://apps.apple.com/us/app/headway-self-growth-challenge/id1457185832"
	googlePlayLink = "https://play.google.com/store/apps/details?id=com.headway.books&hl=en&gl=US"
	resApplePayUrl = "https://api.applepay/15"
)

func (c *getProviderMock) GetButton(request button_domain.ButtonRequest, serviceURL string) (*button_domain.Button, *button_domain.ButtonError) {
	return getServiceProviderFunc(request, serviceURL)
}

func TestApplePayServiceNoProductID(t *testing.T) {
	getServiceProviderFunc = func(request button_domain.ButtonRequest, serviceURL string) (*button_domain.Button, *button_domain.ButtonError) {
		return nil, &button_domain.ButtonError{
			AppStoreLink:   appStoreLink,
			GooglePlayLink: googlePlayLink,
		}
	}
	providers.ServiceProvider = &getProviderMock{}

	request := button_domain.ButtonRequest{
		ProductID: "",
	}
	result, err := StripeService.GetButton(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.AppStoreLink, appStoreLink)
	assert.EqualValues(t, err.GooglePlayLink, googlePlayLink)
}

func TestApplePayServiceSuccess(t *testing.T) {
	getServiceProviderFunc = func(request button_domain.ButtonRequest, serviceURL string) (*button_domain.Button, *button_domain.ButtonError) {
		return &button_domain.Button{
			Link: "https://api.stripe/754",
		}, nil
	}
	providers.ServiceProvider = &getProviderMock{}

	request := button_domain.ButtonRequest{
		ProductID: strconv.FormatInt(754, 10),
	}
	result, err := StripeService.GetButton(request)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, result.Link, resApplePayUrl)
}
