package structs

type PhoneNumber struct {
	FriendlyName string `json:"friendly_name"`
	PhoneNumber  string `json:"phone_number"`
	Sid          string `json:"sid"`
	VoiceUrl     string `json:"voice_url"`
	VoiceMethod  string `json:"voice_method"`
	SMSUrl       string `json:"sms_url"`
}
