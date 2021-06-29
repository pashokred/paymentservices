package button_domain

const (
	appStoreLink   = "https://apps.apple.com/us/app/headway-self-growth-challenge/id1457185832"
	googlePlayLink = "https://play.google.com/store/apps/details?id=com.headway.books&hl=en&gl=US"
)

type ButtonError struct {
	AppStoreLink   string `json:"AppStore"`
	GooglePlayLink string `json:"GooglePlay"`
}

// Easily can add support of error messages or codes to ButtonError on demand

func NewButtonError() *ButtonError {
	return &ButtonError{
		AppStoreLink:   appStoreLink,
		GooglePlayLink: googlePlayLink,
	}
}
