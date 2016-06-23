package dialonce

import (
	"encoding/json"
)


// IVR ...
type IVR struct {
	*Client
}

// GetServiceStatusOutput request output.
type GetServiceStatusOutput struct {
	Status           bool             `json:"status"`
}

// GetServiceStatus ...
func (c *IVR) GetServiceStatus() (out *GetServiceStatusOutput, err error) {
	body, err := c.call("GET", "ivr/status", nil)
  if err != nil {
		return
	}
  defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

// IsMobilePhoneNumberInput request input.
type IsMobilePhoneNumberInput struct {
	Number           string             `json:"number"`
  CultureISO       string             `json:"cultureISO"`
}

// IsMobilePhoneNumberOutput request output.
type IsMobilePhoneNumberOutput struct {
	Mobile           bool             `json:"mobile"`
}

// IsMobilePhoneNumberRequest ...
func (c *IVR) IsMobilePhoneNumberRequest(in *IsMobilePhoneNumberInput) (out *IsMobilePhoneNumberOutput, err error) {
  body, err := c.call("GET", "phoneNumbers/isMobile", in)
  if err != nil {
    return
  }
  defer body.Close()

  err = json.NewDecoder(body).Decode(&out)
  return
}

// IsMobilePhoneNumber ...
func (c *IVR) IsMobilePhoneNumber(number string, params ...string) (out *IsMobilePhoneNumberOutput, err error) {
	isMobilePhoneNumberInput := &IsMobilePhoneNumberInput{
		Number: number,
	}

	if len(params) > 0 {
		isMobilePhoneNumberInput.CultureISO = params[0]
	}

	return c.IsMobilePhoneNumberRequest(isMobilePhoneNumberInput)
}

// IVRLogInputData request input.
type IVRLogInputData struct {
	Culture           string             `json:"culture"`
  SDA      	 				string             `json:"sda"`
	DID         			string						 `json:"did"`
}

// IVRLogInput request input.
type IVRLogInput struct {
	Called           	string             	`json:"called"`
  Caller      	 		string            	`json:"caller"`
	Type         			string						 	`json:"type"`
	Data						 	IVRLogInputData  		`json:"data"`
}

// IVRLogOutput request output.
type IVRLogOutput struct {
	Success           string             	`json:"success"`
}


// SendIVRLog ...
func (c *IVR) SendIVRLog(in *IVRLogInput, ivrLogType string) (out *IVRLogOutput, err error) {
	in.Type = ivrLogType

	body, err := c.call("POST", "ivrlogs", in)
  if err != nil {
		return
	}
	defer body.Close()
	err = json.NewDecoder(body).Decode(&out)
	return
}

// CallStart ...
func (c *IVR) CallStart(in *IVRLogInput) (out *IVRLogOutput, err error) {
	c.SendIVRLog(in, "call-start")
	return
}

// CallEnd ...
func (c *IVR) CallEnd(in *IVRLogInput) (out *IVRLogOutput, err error) {
	c.SendIVRLog(in, "call-end")
	return
}

// UserWantsToContinueWithSMS ...
func (c *IVR) UserWantsToContinueWithSMS(in *IVRLogInput) (out *IVRLogOutput, err error) {
	c.SendIVRLog(in, "answer-get-sms")
	return
}

// UserPreferToContinueWithIVR ...
func (c *IVR) UserPreferToContinueWithIVR(in *IVRLogInput) (out *IVRLogOutput, err error) {
	c.SendIVRLog(in, "answer-no-sms")
	return
}
