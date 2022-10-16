package dots

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	api := New(clientID, apiKey, true)

	p := &InputCreateUserParams{
		Email:       "joshua.forgy@gmail.com",
		FirstName:   "joshua",
		LastName:    "forgy",
		CountryCode: "1",
		PhoneNumber: "6268262883",
		Username:    "jforgy",
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
		VerificationID: "44b0f285-1399-4007-9cb9-75e5b614dfee",
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
		VerificationID:    "44b0f285-1399-4007-9cb9-75e5b614dfee",
		VerificationToken: "226096",
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
		UserID:         "f259a622-0929-4ce8-bb88-e28f38b2b2d6",
		PayoutMethod:   "venmo",
		PayoutID:       "@Joshua-Forgy",
		ACHAccountType: "checking",
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
