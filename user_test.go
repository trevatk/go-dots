package dots

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputCreateUserParams{
		Email:       "",
		FirstName:   "",
		LastName:    "",
		CountryCode: "",
		PhoneNumber: "",
		Username:    "",
	}

	u, e := api.CreateUser(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	t.Logf("user verification_id %s", u.VerificationID)
}

func TestSendUserVerificationToken(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputSendVerificationTokenParams{
		VerificationID: "",
	}

	r, e := api.SendVerificationToken(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, true, r.Success)
}

func TestVerifyUserToken(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputVerifyUserTokenParams{
		VerificationID:    "",
		VerificationToken: "",
	}

	r, e := api.VerifyUserToken(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, true, r.Success)
}

func TestCreateUserPayout(t *testing.T) {

	ctx := context.TODO()

	api := New(clientID, apiKey, true)

	p := &InputProgramaticalPayoutParams{
		UserID:         "",
		PayoutMethod:   "",
		PayoutID:       "",
		ACHAccountType: "",
		SetDefault:     true,
	}

	po, e := api.CreateUserPayout(ctx, p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, true, po.Success)

	t.Logf("successfully created payout ach_account id %s", po.ACHAccountID)
}
