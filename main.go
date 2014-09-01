package main

import (
	"fmt"
	"github.com/andrewwatson/TwilioGo/client"
	"net/http"
)

const (
	twilioSID    = "<YOUR SID>"
	twilioSecret = "<YOUR TOKEN>"
)

func main() {

	client := http.Client{}
	twilio := twiliogo.NewTwilioClient(twilioSID, twilioSecret)

	numbers, err := twilio.SearchNumbers(client, "404")

	if err != nil {
		fmt.Printf("%#v", err)
	}

	firstFive := numbers[0:5]

	for _, v := range firstFive {
		fmt.Printf("NUMBERS: %#v\n", v.PhoneNumber)
	}

}
