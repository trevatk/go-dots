package dots

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	clientID := os.Getenv("DOTS_CLIENT_ID")
	apiKey := os.Getenv("DOTS_API_KEY")

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

	clientID := os.Getenv("DOTS_CLIENT_ID")
	apiKey := os.Getenv("DOTS_API_KEY")

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

	clientID := os.Getenv("DOTS_CLIENT_ID")
	apiKey := os.Getenv("DOTS_API_KEY")

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

func TestRetrieveAppUserIDs(t *testing.T) {

	clientID := os.Getenv("DOTS_CLIENT_ID")
	apiKey := os.Getenv("DOTS_API_KEY")

	api := New(clientID, apiKey, true)

	r, e := api.RetrieveAppUserIDs(context.TODO())
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, true, r.Success)

	for _, id := range r.Users {
		t.Logf("retrieve app id %s", id)
	}
}

func TestGetUserByID(t *testing.T) {

	clientID := os.Getenv("DOTS_CLIENT_ID")
	apiKey := os.Getenv("DOTS_API_KEY")

	uID := "changeme"

	api := New(clientID, apiKey, true)

	p := &InputGetUserParams{
		UserID: uID,
	}

	u, e := api.GetUserByID(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, uID, u.ID)
}

func TestAddUserKYC(t *testing.T) {

	clientID := os.Getenv("DOTS_CLIENT_ID")
	apiKey := os.Getenv("DOTS_API_KEY")

	api := New(clientID, apiKey, true)

	p := &InputAddUserKYCParams{}

	r, e := api.AddUserKYC(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, true, r.Success)
}

func TestListUserBankAccounts(t *testing.T) {

	clientID := os.Getenv("DOTS_CLIENT_ID")
	apiKey := os.Getenv("DOTS_API_KEY")

	api := New(clientID, apiKey, true)

	r, e := api.ListUserBankAccounts(context.TODO(), "")
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, true, r.Success)
}

func TestGetUserWallet(t *testing.T) {

	clientID := os.Getenv("DOTS_CLIENT_ID")
	apiKey := os.Getenv("DOTS_API_KEY")

	api := New(clientID, apiKey, true)

	r, e := api.GetUserWallet(context.TODO(), "")
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, true, r.Success)
}

func TestGetLimitedUserByVerificationID(t *testing.T) {

	clientID := os.Getenv("DOTS_CLIENT_ID")
	apiKey := os.Getenv("DOTS_API_KEY")

	api := New(clientID, apiKey, true)

	r, e := api.GetLimitedUserByVerificationID(context.TODO(), "")
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, true, r.Success)
}

func TestGenerateRefillUserWalletLink(t *testing.T) {

	clientID := os.Getenv("DOTS_CLIENT_ID")
	apiKey := os.Getenv("DOTS_API_KEY")

	api := New(clientID, apiKey, true)

	p := &InputRefillWalletLinkParams{}

	r, e := api.GenerateRefillUserWalletLink(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, true, r.Success)
}

func TestGeneratePayoutUserWalletLink(t *testing.T) {

	clientID := os.Getenv("DOTS_CLIENT_ID")
	apiKey := os.Getenv("DOTS_API_KEY")

	api := New(clientID, apiKey, true)

	p := &InputPayoutWalletLinkParams{}

	r, e := api.GeneratePayoutUserWalletLink(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, true, r.Success)
}

func TestCreateUserPayout(t *testing.T) {

	clientID := os.Getenv("DOTS_CLIENT_ID")
	apiKey := os.Getenv("DOTS_API_KEY")

	api := New(clientID, apiKey, true)

	p := &InputCreateUserPayoutParams{
		UserID:         "",
		PayoutMethod:   "",
		PayoutID:       "",
		ACHAccountType: "",
		SetDefault:     true,
	}

	po, e := api.CreateUserPayout(context.TODO(), p)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}

	assert.Equal(t, true, po.Success)

	t.Logf("successfully created payout ach_account id %s", po.ACHAccountID)
}
