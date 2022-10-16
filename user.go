package dots

import (
	"context"
	"encoding/json"
	"fmt"
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

// InputRefillWalletLinkParams
type InputRefillWalletLinkParams struct {
	UserID string `json:"user_id"`
}

// InputPayoutWalletLinkParams
type InputPayoutWalletLinkParams struct {
	UserID         string `json:"user_id"`
	VerificationID string `json:"verification_id,omitempty"` // optional
}

// InputProgramaticalPayoutParams
type InputProgramaticalPayoutParams struct {
	UserID           string `json:"user_id"`
	PayoutMethod     string `json:"payout_method"`
	PayoutID         string `json:"payout_id"`
	ACHRoutingNumber string `json:"ach_routing_number"`
	ACHAccountNumber string `json:"ach_account_number"`
	ACHAccountType   string `json:"ach_account_type"`
	SetDefault       bool   `json:"set_default"`
}

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

// AddUserKYCResponse
type AddUserKYCResponse struct {
	Success bool   `json:"success"` // boolean
	Message string `json:"message"`
}

// ListUserBankAccountResponse
type ListUserBankAccountResponse struct {
	Success  bool       `json:"success"`
	Accounts []*Account `json:"accounts"`
}

// GetUserWalletResponse
type GetUserWalletResponse struct {
	Success bool    `json:"success"`
	Wallet  *Wallet `json:"Wallet"`
}

// GetLimitedUserResponse
type GetLimitedUserResponse struct {
	Success   bool         `json:"success"`
	Connected bool         `json:"connected"`
	User      *UserLimited `json:"user"`
}

// RefillWalletLinkResponse
type RefillWalletLinkResponse struct {
	Success bool   `json:"success"`
	Link    string `json:"link"`
}

// PayoutWalletLinkResponse
type PayoutWalletLinkResponse struct {
	Success bool   `json:"success"`
	Link    string `json:"link"`
}

// ProgramaticalPayoutResponse
type ProgramaticalPayoutResponse struct {
	Success      bool   `json:"success"`
	ACHAccountID string `json:"ach_account_id"`
}

// CreateUser create/connect a new user
func (api *API) CreateUser(ctx context.Context, in *InputCreateUserParams) (*CreateUserResponse, error) {

	r := api.h + "/api/users/create"
	b, e := api.cl.post(ctx, r, in)
	if e != nil {
		return nil, e
	}

	var crb CreateUserResponse
	if e := json.Unmarshal(b, &crb); e != nil {
		return nil, fmt.Errorf("dots api create user json.Unmarshal err %v", e)
	}

	return &crb, nil
}

// SendVerificationToken send a verification token to the user
func (api *API) SendVerificationToken(ctx context.Context, in *InputSendVerificationTokenParams) (*SendVerificationResponse, error) {

	r := api.h + "/api/users/send_verification_token"
	b, e := api.cl.post(ctx, r, in)
	if e != nil {
		return nil, e
	}

	var svr SendVerificationResponse
	if e := json.Unmarshal(b, &svr); e != nil {
		return nil, fmt.Errorf("dots api send verification token json.Unmarshal err %v", e)
	}

	return &svr, nil
}

// VerifyUserToken verify a user with the token sent to them
func (api *API) VerifyUserToken(ctx context.Context, in *InputVerifyUserTokenParams) (*VerifyUserTokenResponse, error) {

	r := api.h + "/api/users/verify_user"
	b, e := api.cl.post(ctx, r, in)
	if e != nil {
		return nil, e
	}

	var vutp VerifyUserTokenResponse
	if e := json.Unmarshal(b, &vutp); e != nil {
		return nil, fmt.Errorf("dots api verify user token json.Unmarshal err %v", e)
	}

	return &vutp, nil
}

// RetrieveAppUserIDs retrieve and filter connected app user IDs
func (api *API) RetrieveAppUserIDs(ctx context.Context) ([]string, error) {

	r := api.h + "/api/users/get"
	b, e := api.cl.get(ctx, r)
	if e != nil {
		return nil, e
	}

	var re []string
	if e := json.Unmarshal(b, &re); e != nil {
		return nil, fmt.Errorf("dots api retrieve app user ids json.Unmarshal err %v", e)
	}

	return re, nil
}

// GetUserByID get the user by their id
func (api *API) GetUserByID(ctx context.Context, in *InputGetUserParams) (*GetUserByIDResponse, error) {

	r := api.h + "/api/users/get/" + in.UserID
	b, e := api.cl.get(ctx, r)
	if e != nil {
		return nil, fmt.Errorf("dots api get user by id http.NewRequest err %v", e)
	}

	var u GetUserByIDResponse
	if e := json.Unmarshal(b, &u); e != nil {
		return nil, fmt.Errorf("dots api get user by id json.Unmarshal err %v", e)
	}

	return &u, nil
}

// AddUserKYCWithContext add KYC or KYB information for user
func (api *API) AddUserKYCWithContext(ctx context.Context, in *InputAddUserKYCParams) (*AddUserKYCResponse, error) {

	r := api.h + "/api/users/add_kyc_information"
	b, e := api.cl.post(ctx, r, in)
	if e != nil {
		return nil, e
	}

	var ur AddUserKYCResponse
	if e := json.Unmarshal(b, &ur); e != nil {
		return nil, fmt.Errorf("dots api add user kyc json.Unmarshal err %v", e)
	}

	return &ur, nil
}

// ListUserBankAccounts list international bank accounts owned by a user
func (api *API) ListUserBankAccounts(ctx context.Context, ID string) (*ListUserBankAccountResponse, error) {

	r := fmt.Sprintf("/api/users/get/%s/intl_bank_accounts", ID)
	h := api.h + r

	b, e := api.cl.get(ctx, h)
	if e != nil {
		return nil, e
	}

	var l ListUserBankAccountResponse
	if e := json.Unmarshal(b, &l); e != nil {
		return nil, fmt.Errorf("dots api list user bank accounts json.Unmarshal err %v", e)
	}

	return &l, nil
}

// GetUserWallet get user wallet information
func (api *API) GetUserWallet(ctx context.Context, ID string) (*GetUserWalletResponse, error) {

	r := api.h + "/api/users/wallet/get/" + ID
	b, e := api.cl.get(ctx, r)
	if e != nil {
		return nil, e
	}

	var w GetUserWalletResponse
	if e := json.Unmarshal(b, &w); e != nil {
		return nil, fmt.Errorf("dots api get user wallet json.Unmarshal err %v", e)
	}

	return &w, nil
}

// GetLimitedUserByVerificationID retrieve limited information about a user given a verification id
func (api *API) GetLimitedUserByVerificationID(ctx context.Context, verificationID string) (*GetLimitedUserResponse, error) {

	r := api.h + "/api/users/get_by_verification_id/" + verificationID
	b, e := api.cl.get(ctx, r)
	if e != nil {
		return nil, e
	}

	var lu GetLimitedUserResponse
	if e := json.Unmarshal(b, &lu); e != nil {
		return nil, fmt.Errorf("dots api get limited user by verification id json.Unmarshal err %v", e)
	}

	return &lu, nil
}

// GenerateRefillUserWalletLink generate a link to refill the user's wallet
func (api *API) GenerateRefillUserWalletLink(ctx context.Context, in *InputRefillWalletLinkParams) (*RefillWalletLinkResponse, error) {

	r := api.h + "/api/users/wallet/refill"
	b, e := api.cl.post(ctx, r, in)
	if e != nil {
		return nil, e
	}

	var rw RefillWalletLinkResponse
	if e := json.Unmarshal(b, &rw); e != nil {
		return nil, fmt.Errorf("dots api generate refill user wallet link json.Unmarshal err %v", e)
	}

	return &rw, nil
}

// GeneratePayoutUserWalletLink generate a link to payout the user's wallet
func (api *API) GeneratePayoutUserWalletLink(ctx context.Context, in *InputPayoutWalletLinkParams) (*PayoutWalletLinkResponse, error) {

	r := api.h + "/api/users/wallet/payout"
	b, e := api.cl.post(ctx, r, in)
	if e != nil {
		return nil, e
	}

	var wl PayoutWalletLinkResponse
	if e := json.Unmarshal(b, &wl); e != nil {
		return nil, fmt.Errorf("dots api generate payout user wallet link json.Unmarshal err %v", e)
	}

	return &wl, nil
}

/*
CreateUserPayout programatically add a payout method for a user
* endpoint /api/users/wallet/add_payout_method
*/
func (api *API) CreateUserPayout(ctx context.Context, in *InputProgramaticalPayoutParams) (*ProgramaticalPayoutResponse, error) {

	r := api.h + "/api/users/wallet/add_payout_method"
	b, e := api.cl.post(ctx, r, in)
	if e != nil {
		return nil, e
	}

	fmt.Println("html response ", string(b))

	var p ProgramaticalPayoutResponse
	if e := json.Unmarshal(b, &p); e != nil {
		return nil, fmt.Errorf("dots api create user payout json.Unmarshal err %v html response %s", e, string(b))
	}

	return &p, nil
}
