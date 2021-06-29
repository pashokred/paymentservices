package button_domain

type Button struct {
	Link string `json:"link"`
}

type ButtonRequest struct {
	ProductID string `json:"productID"`
}
