package goverify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// VerifySendSMS The function of sending SMS receives the recipient's number as a parameter.
func (t *TwilioClient) VerifySendSMS(to string) (smsResp *VerifySendSMSResponse, exception *Exception, err error) {
	msgData := url.Values{}
	client := &http.Client{}
	urlStr := t.VerifyURL + t.TwilioService + "/Verifications"

	msgData.Set("To", to)
	msgData.Set("Channel", "sms")

	msgDataReader := *strings.NewReader(msgData.Encode())

	req, err := http.NewRequest("POST", urlStr, &msgDataReader)
	if err != nil {
		return smsResp, exception, err
	}

	req.SetBasicAuth(t.AccountSid, t.AuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return smsResp, exception, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return smsResp, exception, err
	}

	if resp.StatusCode != http.StatusCreated {
		exception = new(Exception)
		err = json.Unmarshal(respBody, exception)
		return smsResp, exception, err
	}

	smsResp = new(VerifySendSMSResponse)
	err = json.Unmarshal(respBody, smsResp)

	return smsResp, exception, err
}

// VerifyCheckSMS Code verification function where we receive as parameters the recipient's number and the code sent by SMS.
func (t *TwilioClient) VerifyCheckSMS(to string, code string) (smsResp *VerifyCheckSMSResponse, exception *Exception, err error) {
	msgData := url.Values{}
	client := &http.Client{}
	urlStr := t.VerifyURL + t.TwilioService + "/VerificationCheck"

	msgData.Set("To", to)
	msgData.Set("Code", code)

	msgDataReader := *strings.NewReader(msgData.Encode())

	req, err := http.NewRequest("POST", urlStr, &msgDataReader)
	if err != nil {
		return smsResp, exception, err
	}

	req.SetBasicAuth(t.AccountSid, t.AuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return smsResp, exception, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return smsResp, exception, err
	}

	if resp.StatusCode != http.StatusCreated {
		exception = new(Exception)
		err = json.Unmarshal(respBody, exception)
		return smsResp, exception, err
	}

	smsResp = new(VerifyCheckSMSResponse)
	err = json.Unmarshal(respBody, smsResp)

	return smsResp, exception, err
}
