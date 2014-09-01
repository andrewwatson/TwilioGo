package twiliogo

import (
	// "errors"
	// "appengine"
	// "appengine/urlfetch"
	"encoding/json"
	"fmt"
	"github.com/andrewwatson/TwilioGo/structs"
	ioutil "io/ioutil"
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
func (t *TwilioClient) SearchNumbers(client http.Client, areaCode string) (numbers []structs.AvailablePhoneNumber, err error) {

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
		fmt.Printf("ERR %#v\n", clientError)
	}

	numbers = response.AvailableNumbers

	return
}

func (t *TwilioClient) PurchaseNumber(client http.Client, phonenumber string) (number structs.PhoneNumber) {

	data := url.Values{}
	data.Add("PhoneNumber", phonenumber)

	fmt.Printf("DATA %#v\n", data)

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

	rawBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("RAW: %#v\n", rawBody)

	fmt.Printf("RESP %#v\n", response)

	if clientError != nil {
		fmt.Printf("ERR %#v\n", clientError)
	}

	// number = response.AvailableNumbers

	return

}
