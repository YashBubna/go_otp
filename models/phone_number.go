package models

import (
	"fmt"

	"github.com/RaghavTheGreat1/go_otp/constants"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

type PhoneNumber struct {
	PhoneNumber int    `json:"phone_number" binding:"required"`
	CountryCode string `json:"country_code" binding:"required"`
	Code        string `json:"code"`
}

func (p *PhoneNumber) SendOtp() (string, error) {
	params := &openapi.CreateVerificationParams{}
	params.SetTo(fmt.Sprintf("%s%d", p.CountryCode, p.PhoneNumber))
	params.SetChannel("sms")

	resp, err := constants.Client.VerifyV2.CreateVerification(constants.VERIFY_SERVICE_SID, params)

	if err != nil {
		return "", fmt.Errorf(err.Error())
	} else {
		return *resp.Sid, nil
	}
}

func (p *PhoneNumber) CheckOtp() (bool, error) {

	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(fmt.Sprintf("%s%d", p.CountryCode, p.PhoneNumber))
	params.SetCode(p.Code)

	resp, err := constants.Client.VerifyV2.CreateVerificationCheck(constants.VERIFY_SERVICE_SID, params)

	if err != nil {
		return false, fmt.Errorf(err.Error())
	} else if *resp.Status == "approved" {
		return true, nil
	} else {
		return false, nil
	}
}
