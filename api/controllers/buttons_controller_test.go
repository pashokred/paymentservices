package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"paymentservices/api/domain/button_domain"
	"paymentservices/api/services/base_service"
	"reflect"
	"testing"
)

var getButtonsFunc func(input button_domain.ButtonRequest) ([]button_domain.Button, *button_domain.ButtonError)

type baseServiceMock struct{}

func (s *baseServiceMock) GetButtons(input button_domain.ButtonRequest) ([]button_domain.Button, *button_domain.ButtonError) {
	return getButtonsFunc(input)
}

func TestGetButtonsReturnStores(t *testing.T) {
	getButtonsFunc = func(input button_domain.ButtonRequest) ([]button_domain.Button, *button_domain.ButtonError) {
		return nil, button_domain.NewButtonError()
	}
	base_service.BaseService = &baseServiceMock{}
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/products/{id}", GetButtons)
	ts := httptest.NewServer(http.HandlerFunc(GetButtons))
	defer ts.Close()
	t.Run("Invalid url", func(t *testing.T) {
		response, err := http.Get("http://127.0.0.1:3000" + "/invalid")
		if err != nil {
			t.Errorf("Expected nil, received %s", err)
		}
		if response == nil {
			t.Errorf("Expected not nil, received nil response")
			return
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal(err)
				return
			}
		}(response.Body)

		got, err := ioutil.ReadAll(response.Body)
		if err != nil {
			t.Errorf(err.Error())
		}
		want := []byte(`{"AppStore": "https://apps.apple.com/us/app/headway-self-growth-challenge/id1457185832"}, {"GooglePlay": "https://play.google.com/store/apps/details?id=com.headway.books&hl=en&gl=US"}`)
		if eq, err := JSONBytesEqual(want, got); !eq || err != nil {
			t.Errorf("Expected store links, received %s, and error %s", string(got), err)
		}
	})
}

func TestGetButtonsSuccess(t *testing.T) {
	getButtonsFunc = func(input button_domain.ButtonRequest) ([]button_domain.Button, *button_domain.ButtonError) {
		return nil, button_domain.NewButtonError()
	}
	base_service.BaseService = &baseServiceMock{}
	r := mux.NewRouter()
	r.HandleFunc("/products/{id}", GetButtons)
	ts := httptest.NewServer(r)
	defer ts.Close()
	response, err := http.Get("http://127.0.0.1:3000" + "/products/5")
	if err != nil {
		t.Errorf("Expected not nil, received %s", err)
	}
	if response == nil {
		t.Errorf("Expected not nil, received nil response")
		return
	}
	got, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf(err.Error())
	}
	want := []byte(`[{"link": "https://applepay.api"}, {"link": "https://googlepay.api"}, {"link": "https://stripe.api"}, {"link": "https://paypal.api"}]`)
	if eq, err := JSONBytesEqual(want, got); !eq || err != nil {
		t.Errorf("Expected correct links, received %s, and error %s", string(got), err)
	}
}

// JSONBytesEqual compares the JSON in two byte slices.
func JSONBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}
