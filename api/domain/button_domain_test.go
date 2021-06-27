package domain

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestButton(t *testing.T){
	request := Button{
		ResponseTime: 0,
		ServerName: "applepay.com",
	}

	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var result Button
	err = json.Unmarshal(bytes, &result)

	assert.Nil(t, err)
	assert.EqualValues(t, result.ResponseTime, request.ResponseTime)
	assert.EqualValues(t, result.ServerName, request.ServerName)
}

func TestButtonError(t *testing.T) {
	request := ButtonError{
		Code:         400,
		ErrorMessage: "Bad Request Error",
	}
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var errResult ButtonError
	err = json.Unmarshal(bytes, &errResult)
	assert.Nil(t, err)
	assert.EqualValues(t, errResult.Code, request.Code)
	assert.EqualValues(t, errResult.ErrorMessage, request.ErrorMessage)

}
