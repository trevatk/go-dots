package dots

import (
	"context"
	"testing"
)

func TestCreateUser(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputCreateUserParams{
		Email:       "trevor@badappsdevelopment.com",
		FirstName:   "trevor",
		LastName:    "atkinson",
		CountryCode: "1",
		PhoneNumber: "2083714060",
		Username:    "tatkinson",
	}

	u, e := api.CreateUser(p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	t.Logf("user verification_id %s", u.VerificationID)
}

func TestCreateUserWithContext(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputCreateUserParams{
		Email:       "trevor@badappsdevelopment.com",
		FirstName:   "trevor",
		LastName:    "atkinson",
		CountryCode: "1",
		PhoneNumber: "2083714060",
		Username:    "tatkinson",
	}

	u, e := api.CreateUserWithContext(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	t.Logf("user verification_id %s", u.VerificationID)
}

func TestSendUserVerificationToken(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputSendVerificationTokenParams{
		VerificationID: "2d6b0173-2f5b-449b-9ce8-b153b8f5084b",
	}

	r, e := api.SendVerificationToken(p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	t.Log("user verified ", r.Success)
}
