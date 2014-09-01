package structs

type AvailablePhoneNumber struct {
	FriendlyName string `json:"friendly_name"`
	PhoneNumber  string `json:"phone_number"`
}

type AvailablePhoneNumbersResponse struct {
	Uri              string
	AvailableNumbers []AvailablePhoneNumber `json:"available_phone_numbers"`
}
