package model

type Restaurant struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PhoneNumber string `json:"phoneNumber"`
}
