package dialonce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestIVR_GetServiceStatus(t *testing.T) {
  c := testClient()

	r, err := c.IVR.GetServiceStatus()

	assert.NoError(t, err)
  assert.NotEqual(t, nil, r)
}

func TestIVR_IsEligible(t *testing.T) {
  c := testClient()

	r, err := c.IVR.IsEligible(&IsEligibleInput{
		Called: "+33185086728",
    Caller: "+33651580955",
  })

	assert.NoError(t, err)
  assert.Equal(t, true, r.Eligible)
}

func TestIVR_IsMobilePhoneNumberRequest_WithInternationalMobileNumber(t *testing.T) {
  c := testClient()

	r, err := c.IVR.IsMobilePhoneNumberRequest(&IsMobilePhoneNumberInput{
    Number: "+33651580955",
  })

	assert.NoError(t, err)
  assert.Equal(t, true, r.Mobile)
}

func TestIVR_IsMobilePhoneNumberRequest_WithNationalMobileNumberAndCulture(t *testing.T) {
  c := testClient()

	r, err := c.IVR.IsMobilePhoneNumberRequest(&IsMobilePhoneNumberInput{
    Number: "0651580955",
    CultureISO: "fr",
  })

	assert.NoError(t, err)
  assert.Equal(t, true, r.Mobile)
}

func TestIVR_IsMobilePhoneNumberRequest_WithNationalMobileNumberAndWithoutCulture(t *testing.T) {
  c := testClient()

	r, err := c.IVR.IsMobilePhoneNumberRequest(&IsMobilePhoneNumberInput{
    Number: "0651580955",
  })

	assert.NoError(t, err)
  assert.Equal(t, false, r.Mobile)
}

func TestIVR_IsMobilePhoneNumberRequest_WithInternationalLandlineNumber(t *testing.T) {
  c := testClient()

	r, err := c.IVR.IsMobilePhoneNumberRequest(&IsMobilePhoneNumberInput{
    Number: "+33185086728",
  })

	assert.NoError(t, err)
  assert.Equal(t, false, r.Mobile)
}

func TestIVR_IsMobilePhoneNumberRequest_WithNationalLandlineNumberAndCulture(t *testing.T) {
  c := testClient()

	r, err := c.IVR.IsMobilePhoneNumberRequest(&IsMobilePhoneNumberInput{
    Number: "0185086728",
    CultureISO: "fr",
  })

	assert.NoError(t, err)
  assert.Equal(t, false, r.Mobile)
}

func TestIVR_IsMobilePhoneNumberRequest_WithNationalLandlineNumberAndWithoutCulture(t *testing.T) {
  c := testClient()

	r, err := c.IVR.IsMobilePhoneNumberRequest(&IsMobilePhoneNumberInput{
    Number: "0185086728",
  })

	assert.NoError(t, err)
  assert.Equal(t, false, r.Mobile)
}

func TestIVR_IsMobilePhoneNumberRequest_WithInvalidNumber(t *testing.T) {
  c := testClient()

	r, err := c.IVR.IsMobilePhoneNumberRequest(&IsMobilePhoneNumberInput{
    Number: "0000",
  })

	assert.NoError(t, err)
	assert.Equal(t, false, r.Mobile)
}

func TestIVR_IsMobilePhoneNumberRequest_InvalidResponse(t *testing.T) {
  c := testClient()

	r, err := c.IVR.IsMobilePhoneNumberRequest(&IsMobilePhoneNumberInput{})

	assert.Error(t, err)
	assert.NotEqual(t, true, r)
}

func TestIVR_IsMobilePhoneNumber_WithoutParams(t *testing.T) {
  c := testClient()

	r, err := c.IVR.IsMobilePhoneNumber("+33651580955")

	assert.NoError(t, err)
	assert.Equal(t, true, r.Mobile)
}

func TestIVR_IsMobilePhoneNumber_WithParams(t *testing.T) {
  c := testClient()

	r, err := c.IVR.IsMobilePhoneNumber("0651580955", "fr")

	assert.NoError(t, err)
	assert.Equal(t, true, r.Mobile)
}

func TestIVR_SendIVRLog(t *testing.T) {
  c := testClient()

  r, err := c.IVR.SendIVRLog(&IVRLogInput{},	"test")

  assert.Error(t, err)
	assert.Nil(t, r)
}

func TestIVR_CallStart(t *testing.T) {
  c := testClient()

  r, err := c.IVR.CallStart(&IVRLogInput{
    Called: "+33185086728",
    Caller: "+33651580955",
  })

  assert.NoError(t, err)
	assert.Nil(t, r)
}

func TestIVR_CallEnd(t *testing.T) {
  c := testClient()

  r, err := c.IVR.CallEnd(&IVRLogInput{
    Called: "+33185086728",
    Caller: "+33651580955",
  })

  assert.NoError(t, err)
	assert.Nil(t, r)
}

func TestIVR_UserWantsToContinueWithSMS(t *testing.T) {
  c := testClient()

	r, err := c.IVR.UserWantsToContinueWithSMS(&IVRLogInput{
    Called: "+33185086728",
    Caller: "+33651580955",
  })

	assert.NoError(t, err)
  assert.Nil(t, r)
}

func TestIVR_UserPreferToContinueWithIVR(t *testing.T) {
  c := testClient()

	r, err := c.IVR.UserPreferToContinueWithIVR(&IVRLogInput{
    Called: "+33185086728",
    Caller: "+33651580955",
  })

	assert.NoError(t, err)
  assert.Nil(t, r)
}

func TestIVR_IVRServiceRequest(t *testing.T) {
  c := testClient()

	r, err := c.IVR.SendServiceRequest(&IVRServiceRequestInput{
    Called: "+33185086728",
    Caller: "+33651580955",
  })

	assert.NoError(t, err)
  assert.Equal(t, r.Success, true)
}

func TestIVR_IVRServiceRequest_WithMissingParameters(t *testing.T) {
  c := testClient()

	_, err := c.IVR.SendServiceRequest(&IVRServiceRequestInput{
    Caller: "+33651580955",
  })

	assert.Error(t, err)
}

func TestIVR_IVRServiceRequest_WithInvalidParameters(t *testing.T) {
  c := testClient()

	_, err := c.IVR.SendServiceRequest(&IVRServiceRequestInput{
		Called: "YYYYYYY",
    Caller: "XXXXXXX",
  })

	assert.Error(t, err)
}
