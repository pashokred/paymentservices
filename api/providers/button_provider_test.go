package providers

import (
	"github.com/stretchr/testify/assert"
	"paymentservices/api/domain/button_domain"
	"strconv"
	"testing"
)

const (
	applePayURL  = "https://api.applepay/%s"
	googlePayURL = "https://api.googlepay/%s"
	paypalURL    = "https://api.paypal/%s"
	stripeURL    = "https://api.stripe/%s"
)

const (
	resApplePayURL  = "https://applepay.api"
	resGooglePayURL = "https://googlepay.api"
	resStripeURL    = "https://stripe.api"
	resPaypalURL    = "https://paypal.api"
)

const (
	appStoreLink   = "https://apps.apple.com/us/app/headway-self-growth-challenge/id1457185832"
	googlePlayLink = "https://play.google.com/store/apps/details?id=com.headway.books&hl=en&gl=US"
)

func TestGetButtonNoErrorApplePay(t *testing.T) {
	response, err := ServiceProvider.GetButton(button_domain.ButtonRequest{ProductID: strconv.FormatInt(123, 10)}, applePayURL)
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.EqualValues(t, resApplePayURL, response.Link)
}

func TestGetButtonNoErrorGooglePay(t *testing.T) {
	response, err := ServiceProvider.GetButton(button_domain.ButtonRequest{ProductID: strconv.FormatInt(123, 10)}, googlePayURL)
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.EqualValues(t, resGooglePayURL, response.Link)
}

func TestGetButtonNoErrorPaypal(t *testing.T) {
	response, err := ServiceProvider.GetButton(button_domain.ButtonRequest{ProductID: strconv.FormatInt(123, 10)}, paypalURL)
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.EqualValues(t, resPaypalURL, response.Link)
}

func TestGetButtonNoErrorStripe(t *testing.T) {
	response, err := ServiceProvider.GetButton(button_domain.ButtonRequest{ProductID: strconv.FormatInt(123, 10)}, stripeURL)
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.EqualValues(t, resStripeURL, response.Link)
}

func TestGetButtonErrorReturnsAppLinks(t *testing.T) {
	response, err := ServiceProvider.GetButton(button_domain.ButtonRequest{ProductID: strconv.FormatInt(123, 10)}, "Incorrect value")
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, googlePlayLink, err.GooglePlayLink)
	assert.EqualValues(t, appStoreLink, err.AppStoreLink)
}
