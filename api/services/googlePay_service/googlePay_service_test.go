package googlePay_service

import (
	"github.com/stretchr/testify/assert"
	"paymentservices/api/domain/button_domain"
	"paymentservices/api/providers"
	"strconv"
	"testing"
)

var (
	getServiceProviderFunc func(request button_domain.ButtonRequest, serviceURL string) (*button_domain.Button, *button_domain.ButtonError)
)

type getProviderMock struct{}

const (
	appStoreLink    = "https://apps.apple.com/us/app/headway-self-growth-challenge/id1457185832"
	googlePlayLink  = "https://play.google.com/store/apps/details?id=com.headway.books&hl=en&gl=US"
	resGooglePayUrl = "https://api.googlepay/5497"
)

func (c *getProviderMock) GetButton(request button_domain.ButtonRequest, serviceURL string) (*button_domain.Button, *button_domain.ButtonError) {
	return getServiceProviderFunc(request, googlePayUrl)
}

func TestGooglePayServiceNoProductID(t *testing.T) {
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
	result, err := GooglePayService.GetButton(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.AppStoreLink, appStoreLink)
	assert.EqualValues(t, err.GooglePlayLink, googlePlayLink)
}

func TestGooglePayServiceSuccess(t *testing.T) {
	getServiceProviderFunc = func(request button_domain.ButtonRequest, serviceURL string) (*button_domain.Button, *button_domain.ButtonError) {
		return &button_domain.Button{
			Link: "https://api.googlepay/5497",
		}, nil
	}

	providers.ServiceProvider = &getProviderMock{}

	request := button_domain.ButtonRequest{
		ProductID: strconv.FormatInt(15, 10),
	}
	result, err := GooglePayService.GetButton(request)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, result.Link, resGooglePayUrl)
}
