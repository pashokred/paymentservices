package base_service

import (
	"github.com/stretchr/testify/assert"
	"paymentservices/api/domain/button_domain"
	"paymentservices/api/services/applePay_service"
	"testing"
)

type getServiceMock struct{}

var getButtonServiceFunc func(input button_domain.ButtonRequest) (*button_domain.Button, *button_domain.ButtonError)

const (
	appStoreLink   = "https://apps.apple.com/us/app/headway-self-growth-challenge/id1457185832"
	googlePlayLink = "https://play.google.com/store/apps/details?id=com.headway.books&hl=en&gl=US"
)

func (c *getServiceMock) GetButton(input button_domain.ButtonRequest) (*button_domain.Button, *button_domain.ButtonError) {
	return getButtonServiceFunc(input)
}

func TestBaseServiceOneOfServicesError(t *testing.T) {
	getButtonServiceFunc = func(input button_domain.ButtonRequest) (*button_domain.Button, *button_domain.ButtonError) {
		return nil, &button_domain.ButtonError{
			AppStoreLink:   appStoreLink,
			GooglePlayLink: googlePlayLink,
		}
	}

	applePay_service.ApplePayService = &getServiceMock{}
	request := button_domain.ButtonRequest{
		ProductID: "1234",
	}
	result, err := BaseService.GetButtons(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, appStoreLink, err.AppStoreLink)
	assert.EqualValues(t, googlePlayLink, err.GooglePlayLink)
}

func TestBaseServiceSuccess(t *testing.T) {
	request := button_domain.ButtonRequest{
		ProductID: "124123",
	}
	result, err := BaseService.GetButtons(request)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	want := []button_domain.Button{
		{Link: "https://applepay.api"},
		{Link: "https://googlepay.api"},
		{Link: "https://stripe.api"},
		{Link: "https://paypal.api"},
	}
	assert.EqualValues(t, want, result)
}
