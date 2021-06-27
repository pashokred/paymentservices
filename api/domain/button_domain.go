package domain

type Button struct {
	ResponseTime int `json:"X-Response-Time"`
	ServerName string `json:"X-Server-Name"`
}

type ButtonRequest struct {
	ProductID string `json:"productID"`
}
