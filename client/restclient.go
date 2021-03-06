package client

import (
	"errors"
	// "appengine"
	// "appengine/urlfetch"
	"encoding/json"
	"fmt"
	"github.com/andrewwatson/TwilioGo/structs"
	// ioutil "io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type TwilioClient struct {
	AccountSid string
	AuthToken  string
}

func NewTwilioClient(account, token string) *TwilioClient {
	t := TwilioClient{account, token}

	return &t
}

// Takes an http.Client as an agrument because AppEngine makes you use their URL fetcher instead of
// the normal http.Client
func (t *TwilioClient) SearchNumbers(client http.Client, areaCode string, results int) (numbers []structs.AvailablePhoneNumber, err error) {

	twilioUrl := fmt.Sprintf(
		"https://api.twilio.com/2010-04-01/Accounts/%s/AvailablePhoneNumbers/US/Local.json?AreaCode=%s",
		t.AccountSid,
		areaCode,
	)

	twilioRequest, err := http.NewRequest(
		"GET",
		twilioUrl,
		nil,
	)

	twilioRequest.SetBasicAuth(t.AccountSid, t.AuthToken)
	resp, clientError := client.Do(twilioRequest)
	defer resp.Body.Close()

	response := new(structs.AvailablePhoneNumbersResponse)
	json.NewDecoder(resp.Body).Decode(&response)

	if clientError != nil {
		err = clientError
	} else {

		numResults := len(response.AvailableNumbers)
		if numResults < results {
			results = numResults
		}

		numbers = response.AvailableNumbers[0:results]
	}

	return
}

func (t *TwilioClient) SendMessage(client http.Client, toNumber, fromNumber, message string) (err error) {

	data := url.Values{}
	data.Add("From", fromNumber)
	data.Add("To", toNumber)
	data.Add("Body", message)

	twilioUrl := fmt.Sprintf(
		"https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json",
		t.AccountSid,
	)

	twilioRequest, err := http.NewRequest(
		"POST",
		twilioUrl,
		strings.NewReader(data.Encode()),
	)

	if err != nil {
		fmt.Printf("ERR %#v", err)
	}

	twilioRequest.SetBasicAuth(t.AccountSid, t.AuthToken)
	twilioRequest.Header.Add("Content-type", "application/x-www-form-urlencoded")
	resp, clientError := client.Do(twilioRequest)
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		err = errors.New(fmt.Sprintf("TWILIO ERROR: %d", resp.StatusCode))
		// fmt.Printf("ERROR: %d", resp.StatusCode)
	}

	return clientError
}

func (t *TwilioClient) PurchaseNumber(client http.Client, phonenumber string, messageurl string) (number structs.PhoneNumber, err error) {

	data := url.Values{}
	data.Add("PhoneNumber", phonenumber)
	data.Add("SmsUrl", messageurl)

	// fmt.Printf("DATA %#v\n", data)

	twilioUrl := fmt.Sprintf(
		"https://api.twilio.com/2010-04-01/Accounts/%s/IncomingPhoneNumbers.json",
		t.AccountSid,
	)

	twilioRequest, err := http.NewRequest(
		"POST",
		twilioUrl,
		strings.NewReader(data.Encode()),
	)

	if err != nil {
		fmt.Printf("ERR %#v", err)
	}

	twilioRequest.SetBasicAuth(t.AccountSid, t.AuthToken)
	twilioRequest.Header.Add("Content-type", "application/x-www-form-urlencoded")
	resp, clientError := client.Do(twilioRequest)
	defer resp.Body.Close()

	if resp.StatusCode > 299 {

	}

	response := new(structs.PhoneNumber)
	json.NewDecoder(resp.Body).Decode(&response)

	// rawBody, _ := ioutil.ReadAll(resp.Body)
	// fmt.Printf("RAW: %#v\n", rawBody)

	// fmt.Printf("RESP %#v\n", response)

	if clientError != nil {
		fmt.Printf("ERR %#v\n", clientError)
	}

	// number = response.AvailableNumbers

	return

}
