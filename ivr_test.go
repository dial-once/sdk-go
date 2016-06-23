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

func TestIVR_IsMobilePhoneNumber_WithInternationalMobileNumber(t *testing.T) {
  c := testClient()

	r, err := c.IVR.IsMobilePhoneNumberRequest(&IsMobilePhoneNumberInput{
    Number: "+33651580955",
  })

	assert.NoError(t, err)
  assert.Equal(t, true, r.Mobile)
}

func TestIVR_IsMobilePhoneNumber_WithNationalMobileNumberAndCulture(t *testing.T) {
  c := testClient()

	r, err := c.IVR.IsMobilePhoneNumberRequest(&IsMobilePhoneNumberInput{
    Number: "0651580955",
    CultureISO: "fr",
  })

	assert.NoError(t, err)
  assert.Equal(t, true, r.Mobile)
}

func TestIVR_IsMobilePhoneNumber_WithNationalMobileNumberAndWithoutCulture(t *testing.T) {
  c := testClient()

	r, err := c.IVR.IsMobilePhoneNumberRequest(&IsMobilePhoneNumberInput{
    Number: "0651580955",
  })

	assert.NoError(t, err)
  assert.Equal(t, false, r.Mobile)
}


func TestIVR_IsMobilePhoneNumber_WithInternationalLandlineNumber(t *testing.T) {
  c := testClient()

	r, err := c.IVR.IsMobilePhoneNumberRequest(&IsMobilePhoneNumberInput{
    Number: "+33185086728",
  })

	assert.NoError(t, err)
  assert.Equal(t, false, r.Mobile)
}

func TestIVR_IsMobilePhoneNumber_WithNationalLandlineNumberAndCulture(t *testing.T) {
  c := testClient()

	r, err := c.IVR.IsMobilePhoneNumberRequest(&IsMobilePhoneNumberInput{
    Number: "0185086728",
    CultureISO: "fr",
  })

	assert.NoError(t, err)
  assert.Equal(t, false, r.Mobile)
}

func TestIVR_IsMobilePhoneNumber_WithNationalLandlineNumberAndWithoutCulture(t *testing.T) {
  c := testClient()

	r, err := c.IVR.IsMobilePhoneNumberRequest(&IsMobilePhoneNumberInput{
    Number: "0185086728",
  })

	assert.NoError(t, err)
  assert.Equal(t, false, r.Mobile)
}

func TestIVR_IsMobilePhoneNumber_InvalidNumber(t *testing.T) {
  c := testClient()

	r, err := c.IVR.IsMobilePhoneNumberRequest(&IsMobilePhoneNumberInput{
    Number: "0000",
  })

	assert.NoError(t, err)
	assert.Equal(t, false, r.Mobile)
}

func TestIVR_IsMobilePhoneNumber_InvalidResponse(t *testing.T) {
  c := testClient()

	r, err := c.IVR.IsMobilePhoneNumberRequest(&IsMobilePhoneNumberInput{})

	assert.Error(t, err)
	assert.NotEqual(t, true, r)
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
