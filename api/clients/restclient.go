package restclient

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"unicode"
)

type clientStruct struct{}

type ClientInterface interface {
	Get(string) (*http.Response, error)
}

var (
	ClientStruct ClientInterface = &clientStruct{}
)

const (
	applePayUrl  = "https://api.applepay/"
	googlePayUrl = "https://api.googlepay/"
	stripeUrl    = "https://api.stripe/"
	paypalUrl    = "https://api.paypal/"
)

// Get method mock
func (ci *clientStruct) Get(url string) (*http.Response, error) {
	switch DigitPrefix(url) {
	case applePayUrl:
		return &http.Response{
			Body:       ioutil.NopCloser(strings.NewReader(`{"Link": "https://applepay.api"}`)),
			StatusCode: http.StatusOK,
		}, nil
	case googlePayUrl:
		return &http.Response{
			Body:       ioutil.NopCloser(strings.NewReader(`{"Link": "https://googlepay.api"}`)),
			StatusCode: http.StatusOK,
		}, nil
	case stripeUrl:
		return &http.Response{
			Body:       ioutil.NopCloser(strings.NewReader(`{"Link": "https://stripe.api"}`)),
			StatusCode: http.StatusOK,
		}, nil
	case paypalUrl:
		return &http.Response{
			Body:       ioutil.NopCloser(strings.NewReader(`{"Link": "https://paypal.api"}`)),
			StatusCode: http.StatusOK,
		}, nil
	default:
		return nil, errors.New("link is not correct to get button")
	}
}

func DigitPrefix(s string) string {
	for i, r := range s {
		if unicode.IsDigit(r) {
			return s[:i]
		}
	}
	return s
}
