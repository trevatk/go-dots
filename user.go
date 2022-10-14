package dots

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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

// InputSendVerificationTokenParams
type InputSendVerificationTokenParams struct {
	VerificationID string `json:"verification_id"` // string <uuid>
}

// InputVerifyUserTokenParams
type InputVerifyUserTokenParams struct {
	VerificationID    string `json:"verification_id"`    // string <uuid>
	VerificationToken string `json:"verification_token"` // string^[0-9]{6,8}$
}

// InputGetUserParams
type InputGetUserParams struct {
	UserID string `json:"user_id"` // string <uuid>
}

// InputAddUserKYCParams
type InputAddUserKYCParams struct {
	UserID       string         `json:"user_id"`                 // string <uuid>
	EntityType   EntityTypeEnum `json:"entity_type"`             // string enum 'individual', 'business'
	BusinessName string         `json:"business_name,omitempty"` // string required if entity_type = business
	PostCode     string         `json:"post_code"`               // string
	City         string         `json:"city"`                    // string
	Country      CountryEnum    `json:"country"`                 // string enum
	State        StateEnum      `json:"state"`                   // string enum
	Line1        string         `json:"line1"`                   // string
	Line2        string         `json:"line2,omitempty"`         // string
	SSN          string         `json:"ssn,omitempty"`           // string required if entity_type = individual
	EIN          string         `json:"ein,omitempty"`           // string required if entity_type = business
}

// InputAddUserKYBParams
type InputAddUserKYBParams struct{}

// CreateUserResponse
type CreateUserResponse struct {
	Success        bool   `json:"success"`         // boolean
	Action         string `json:"action"`          // string
	VerificationID string `json:"verification_id"` // string <uuid>
}

// SendVerificationResponse
type SendVerificationResponse struct {
	Success bool `json:"success"`
}

// VerifyUserTokenResponse
type VerifyUserTokenResponse struct {
	Success bool        `json:"success"`        // boolean
	User    *VerifyUser `json:"user,omitempty"` // object
	Message string      `json:"mesage"`         // string the error message if there is one
}

// VerifyUser
type VerifyUser struct {
	ID string `json:"id"` // string <uuid>
}

// GetUserByIDResponse
type GetUserByIDResponse struct {
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
func (api *API) CreateUserWithContext(ctx context.Context, in *InputCreateUserParams) (*CreateUserResponse, error) {

	b, e := json.Marshal(in)
	if e != nil {
		return nil, fmt.Errorf("dots api create user json.Marshal err %v", e)
	}

	to, ca := context.WithTimeout(ctx, time.Second*3)
	defer ca()

	r := host + "/api/users/create"
	rq, e := http.NewRequestWithContext(to, http.MethodPost, r, bytes.NewBuffer(b))
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

// SendVerificationToken
func (api *API) SendVerificationToken(in *InputSendVerificationTokenParams) (*SendVerificationResponse, error) {

	b, e := json.Marshal(in)
	if e != nil {
		return nil, fmt.Errorf("dots api send verification token json.Marshal err %v", e)
	}

	r := host + "/api/users/send_verification_token"
	rq, e := http.NewRequest(http.MethodPost, r, bytes.NewBuffer(b))
	if e != nil {
		return nil, fmt.Errorf("dots api send verification token http.NewRequest err %v", e)
	}

	rq.Header.Add(headerAuthorization, headerBasic+api.token)
	rq.Header.Add(headerContentType, headerAppJSON)

	rp, e := api.cl.Do(rq)
	if e != nil {
		return nil, fmt.Errorf("dots api send verification token client.Do err %v", e)
	}
	defer rp.Body.Close()

	bo, e := io.ReadAll(rp.Body)
	if e != nil {
		return nil, fmt.Errorf("dots api send verification token io.ReadAll err %v", e)
	}

	var svr SendVerificationResponse
	if e := json.Unmarshal(bo, &svr); e != nil {
		return nil, fmt.Errorf("dots api send verification token json.Unmarshal err %v", e)
	}

	return &svr, nil
}

// SendVerificationTokenWithContext
func (api *API) SendVerificationTokenWithContext(ctx context.Context, in *InputSendVerificationTokenParams) (*SendVerificationResponse, error) {

	b, e := json.Marshal(in)
	if e != nil {
		return nil, fmt.Errorf("dots api send verification token json.Marshal err %v", e)
	}

	r := host + "/api/users/send_verification_token"
	rq, e := http.NewRequestWithContext(ctx, http.MethodPost, r, bytes.NewBuffer(b))
	if e != nil {
		return nil, fmt.Errorf("dots api send verification token http.NewRequest err %v", e)
	}

	rp, e := api.cl.Do(rq)
	if e != nil {
		return nil, fmt.Errorf("dots api send verification token client.Do err %v", e)
	}
	defer rp.Body.Close()

	rq.Header.Add(headerAuthorization, headerBasic+api.token)
	rq.Header.Add(headerContentType, headerAppJSON)

	bo, e := io.ReadAll(rp.Body)
	if e != nil {
		return nil, fmt.Errorf("dots api send verification token io.ReadAll err %v", e)
	}

	var svr SendVerificationResponse
	if e := json.Unmarshal(bo, &svr); e != nil {
		return nil, fmt.Errorf("dots api send verification token json.Unmarshal err %v", e)
	}

	return &svr, nil
}

// VerifyUserToken
func (api *API) VerifyUserToken(in *InputVerifyUserTokenParams) (*VerifyUserTokenResponse, error) {

	b, e := json.Marshal(in)
	if e != nil {
		return nil, fmt.Errorf("dots api verify user token json.Marshal err %v", e)
	}

	r := host + "/api/users/verify_user"
	rq, e := http.NewRequest(http.MethodPost, r, bytes.NewBuffer(b))
	if e != nil {
		return nil, fmt.Errorf("dots api verify user token http.NewRequest err %v", e)
	}

	rq.Header.Add(headerAuthorization, headerBasic+api.token)
	rq.Header.Add(headerContentType, headerAppJSON)

	rp, e := api.cl.Do(rq)
	if e != nil {
		return nil, fmt.Errorf("dots api verify user token client.Do err %v", e)
	}
	defer rp.Body.Close()

	bo, e := io.ReadAll(rp.Body)
	if e != nil {
		return nil, fmt.Errorf("dots api verify user token io.ReadAll %v", e)
	}

	var vutp VerifyUserTokenResponse
	if e := json.Unmarshal(bo, &vutp); e != nil {
		return nil, fmt.Errorf("dots api verify user token json.Unmarshal err %v", e)
	}

	return &vutp, nil
}

// VerifyUserWithContext
func (api *API) VerifyUserTokenWithContext(ctx context.Context, in *InputVerifyUserTokenParams) (*VerifyUserTokenResponse, error) {

	b, e := json.Marshal(in)
	if e != nil {
		return nil, fmt.Errorf("dots api verify user token json.Marshal err %v", e)
	}

	to, ca := context.WithTimeout(ctx, time.Second*3)
	defer ca()

	r := host + "/api/users/verify_user"
	rq, e := http.NewRequestWithContext(to, http.MethodPost, r, bytes.NewBuffer(b))
	if e != nil {
		return nil, fmt.Errorf("dots api verify user token http.NewRequest err %v", e)
	}

	rq.Header.Add(headerAuthorization, headerBasic+api.token)
	rq.Header.Add(headerContentType, headerAppJSON)

	rp, e := api.cl.Do(rq)
	if e != nil {
		return nil, fmt.Errorf("dots api verify user token client.Do err %v", e)
	}
	defer rp.Body.Close()

	bo, e := io.ReadAll(rp.Body)
	if e != nil {
		return nil, fmt.Errorf("dots api verify user token io.ReadAll %v", e)
	}

	var vutp VerifyUserTokenResponse
	if e := json.Unmarshal(bo, &vutp); e != nil {
		return nil, fmt.Errorf("dots api verify user token json.Unmarshal err %v", e)
	}

	return &vutp, nil
}

// RetrieveAppUserIDs
func (api *API) RetrieveAppUserIDs() ([]string, error) {

	r := host + "/api/users/get"
	rq, e := http.NewRequest(http.MethodGet, r, nil)
	if e != nil {
		return nil, fmt.Errorf("dots api retrieve app user ids http.NewRequest err %v", e)
	}

	rp, e := api.cl.Do(rq)
	if e != nil {
		return nil, fmt.Errorf("dots api retrieve app user ids client.Do err %v", e)
	}
	defer rp.Body.Close()

	rq.Header.Add(headerAuthorization, headerBasic+api.token)
	rq.Header.Add(headerContentType, headerAppJSON)

	bo, e := io.ReadAll(rp.Body)
	if e != nil {
		return nil, fmt.Errorf("dots api retrieve app user ids io.ReadAll err %v", e)
	}

	var re []string
	if e := json.Unmarshal(bo, &re); e != nil {
		return nil, fmt.Errorf("dots api retrieve app user ids json.Unmarshal err %v", e)
	}

	return re, nil
}

// RetrieveAppUserIDsWithContext
func (api *API) RetrieveAppUserIDsWithContext(ctx context.Context) ([]string, error) {

	to, ca := context.WithTimeout(ctx, time.Second*3)
	defer ca()

	r := host + "/api/users/get"
	rq, e := http.NewRequestWithContext(to, http.MethodGet, r, nil)
	if e != nil {
		return nil, fmt.Errorf("dots api retrieve app user ids http.NewRequest err %v", e)
	}

	rp, e := api.cl.Do(rq)
	if e != nil {
		return nil, fmt.Errorf("dots api retrieve app user ids client.Do err %v", e)
	}
	defer rp.Body.Close()

	rq.Header.Add(headerAuthorization, headerBasic+api.token)
	rq.Header.Add(headerContentType, headerAppJSON)

	bo, e := io.ReadAll(rp.Body)
	if e != nil {
		return nil, fmt.Errorf("dots api retrieve app user ids io.ReadAll err %v", e)
	}

	var re []string
	if e := json.Unmarshal(bo, &re); e != nil {
		return nil, fmt.Errorf("dots api retrieve app user ids json.Unmarshal err %v", e)
	}

	return re, nil
}

// GetUserByID
func (api *API) GetUserByID(in *InputGetUserParams) (*GetUserByIDResponse, error) {

	r := host + "/api/users/get/" + in.UserID
	rq, e := http.NewRequest(http.MethodGet, r, nil)
	if e != nil {
		return nil, fmt.Errorf("dots api get user by id http.NewRequest err %v", e)
	}

	rp, e := api.cl.Do(rq)
	if e != nil {
		return nil, fmt.Errorf("dots api get user by id client.Do err %v", e)
	}
	defer rp.Body.Close()

	bo, e := io.ReadAll(rp.Body)
	if e != nil {
		return nil, fmt.Errorf("dots api get user by id io.ReadAll err %v", e)
	}

	var u GetUserByIDResponse
	if e := json.Unmarshal(bo, &u); e != nil {
		return nil, fmt.Errorf("dots api get user by id json.Unmarshal err %v", e)
	}

	return &u, nil
}

// GetUserByIDWithContext
func (api *API) GetUserByIDWithContext(ctx context.Context, in *InputGetUserParams) (*GetUserByIDResponse, error) {

	to, ca := context.WithTimeout(ctx, time.Second*3)
	defer ca()

	r := host + "/api/users/get/" + in.UserID
	rq, e := http.NewRequestWithContext(to, http.MethodGet, r, nil)
	if e != nil {
		return nil, fmt.Errorf("dots api get user by id http.NewRequest err %v", e)
	}

	rp, e := api.cl.Do(rq)
	if e != nil {
		return nil, fmt.Errorf("dots api get user by id client.Do err %v", e)
	}
	defer rp.Body.Close()

	bo, e := io.ReadAll(rp.Body)
	if e != nil {
		return nil, fmt.Errorf("dots api get user by id io.ReadAll err %v", e)
	}

	var u GetUserByIDResponse
	if e := json.Unmarshal(bo, &u); e != nil {
		return nil, fmt.Errorf("dots api get user by id json.Unmarshal err %v", e)
	}

	return &u, nil
}

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
