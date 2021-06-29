package providers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	restclient "paymentservices/api/clients"
	"paymentservices/api/domain/button_domain"
)

type serviceProvider struct{}

type providerInterface interface {
	GetButton(input button_domain.ButtonRequest, serviceURL string) (*button_domain.Button, *button_domain.ButtonError)
}

var ServiceProvider providerInterface = &serviceProvider{}

func (p *serviceProvider) GetButton(input button_domain.ButtonRequest, serviceURL string) (*button_domain.Button, *button_domain.ButtonError) {
	url := fmt.Sprintf(serviceURL, input.ProductID)
	response, err := restclient.ClientStruct.Get(url)
	if err != nil {
		log.Println(fmt.Sprintf("error when trying to get button from apple pay api %s", err.Error()))
		return nil, button_domain.NewButtonError()
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, button_domain.NewButtonError()
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	if response.StatusCode > 299 {
		var errResponse button_domain.ButtonError
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, button_domain.NewButtonError()
		}
		return nil, &errResponse
	}

	var button button_domain.Button
	if err := json.Unmarshal(bytes, &button); err != nil {
		log.Println(fmt.Sprintf("error when trying to unmarshal button successful response: %s", err.Error()))
		return nil, button_domain.NewButtonError()
	}
	return &button, nil
}
