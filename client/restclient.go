package twiliogo

import (
	// "errors"
	// "appengine"
	// "appengine/urlfetch"
	"encoding/json"
	"fmt"
	"github.com/andrewwatson/TwilioGo/structs"
	// "io/ioutil"
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

	data := url.Values{}
	data.Add("AreaCode", areaCode)

	twilioUrl := fmt.Sprintf(
		"https://api.twilio.com/2010-04-01/Accounts/%s/AvailablePhoneNumbers/US/Local.json",
		t.AccountSid,
	)

	twilioRequest, err := http.NewRequest(
		"GET",
		twilioUrl,
		strings.NewReader(data.Encode()),
	)

	twilioRequest.SetBasicAuth(t.AccountSid, t.AuthToken)

	resp, clientError := client.Do(twilioRequest)
	defer resp.Body.Close()

	response := new(structs.AvailablePhoneNumbersResponse)
	json.NewDecoder(resp.Body).Decode(&response)

	// fmt.Printf("RESP %#v\n", response.AvailableNumbers[0])

	if clientError != nil {
		fmt.Printf("ERR %#v\n", clientError)
	}

	numbers = response.AvailableNumbers

	return
}
