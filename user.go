package dots

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// InputCreateUserParams
type InputCreateUserParams struct {
	Email       string `json:"email"`              // string <email"
	CountryCode string `json:"country_code"`       // string^[0-9]{1,3}$
	PhoneNumber string `json:"phone_number"`       // string^[0-9]{1,3}$
	FirstName   string `json:"first_name"`         // string [1..50]
	LastName    string `json:"last_name"`          // string [1..50]
	Username    string `json:"username,omitempty"` // string [1..50]
}

// CreateUserResponse
type CreateUserResponse struct {
	Success        bool   `json:"success"`         // boolean
	Action         string `json:"action"`          // string
	VerificationID string `json:"verification_id"` // string <uuid>
}

// User
type User struct {
	ID            string         `json:"id"`                       // string <uuid>
	Email         string         `json:"email"`                    // string <email>
	Username      string         `json:"username"`                 // string
	FirstName     string         `json:"first_name"`               // string
	LastName      string         `json:"last_name"`                // string
	DisplayName   string         `json:"display_name"`             // string
	PayoutMethods *PayoutMethods `json:"payout_methods,omitempty"` // object
	Wallet        *Wallet        `json:"wallet,omitempty"`         // wallet
}

// PayoutMethods
type PayoutMethods struct {
	ACHAccouns []string // string
	Paypal     string   // string
	Venmo      string   // string
}

// Wallet
type Wallet struct {
	Amount             int // user's balance in cents
	WithdrawableAmount int // user's balance they can withdraw
	CreditBalance      int // user's credit balance
}

// CreateUser
func (api *API) CreateUser(in *InputCreateUserParams) (*CreateUserResponse, error) {

	b, e := json.Marshal(in)
	if e != nil {
		return nil, fmt.Errorf("dots api create user json.Marshal err %v", e)
	}
	r := host + "/api/users/create"
	rq, e := http.NewRequest(http.MethodPost, r, bytes.NewBuffer(b))
	if e != nil {
		return nil, fmt.Errorf("dots api create user http.NewRequest err %v", e)
	}

	rq.Header.Add(headerAuthorization, headerBasic+api.token)
	rq.Header.Add(headerContentType, headerAppJSON)

	rp, e := api.cl.Do(rq)
	if e != nil {
		return nil, fmt.Errorf("dots api create user client.Do err %v", e)
	}
	defer rp.Body.Close()

	bo, e := io.ReadAll(rp.Body)
	if e != nil {
		return nil, fmt.Errorf("dots api create user io.ReadAll err %v", e)
	}

	var crb CreateUserResponse
	if e := json.Unmarshal(bo, &crb); e != nil {
		return nil, fmt.Errorf("dots api create user json.Unmarshal err %v", e)
	}

	return &crb, nil
}

// CreateUserWithContext

// SendVerificationToken

// SendVerificationTokenWithContext

// VerifyUser

// VerifyUserWithContext

// RetrieveAppUserIDs

// RetrieveAppUserIDsWithContext

// GetUserByID

// GetUserByIDWithContext

// AddUserKYC

// AddUserKYCWithContext

// ListUserBankAccounts

// ListUserBankAccountsWithContext

// GetUserWallet

// GetUserWalletWithContext

// GetLimitedUserByVerificationID

// GetLimitedUserByVerificationIDWithContext

// GenerateRefillUserWalletLink

// GenerateRefillUserWalletLinkWithContext

// GeneratePayoutUserWalletLink

// GeneratePayoutUserWalletLinkWithContext

// CreateUserPayout

// CreateUserPayoutWithContext
