package dots

import (
	"context"
	"encoding/json"
	"fmt"
)

// InputCreateTransactionParams create transaction input
type InputCreateTransactionParams struct {
	UserID         string   `json:"user_id"`
	Amount         int      `json:"amount"`
	Receipt        *Receipt `json:"receipt"`
	Notes          string   `json:"notes"`
	IdempotencyKey string   `json:"idempotency_key"` // An optional idempotency key to prevent duplicate transactions, must a valid uuid
}

// InputCreateTransactionUnverifiedUserParams create transaction unverified user input
type InputCreateTransactionUnverifiedUserParams struct {
	VerificationID string      `json:"verification_id"`
	Amount         int         `json:"amount"`
	Receipt        *Receipt    `json:"receipt"`
	Notes          interface{} `json:"notes"`
	IdempotencyKey string      `json:"idempotency_key"` // An optional idempotency key to prevent duplicate transactions, must a valid uuid
}

// InputCreateTransactionsParams create multiple transactions input
type InputCreateTransactionsParams struct {
	Transactions   []*Transaction `json:"transactions"`
	IdempotencyKey string         `json:"idempotency_key"` // An optional idempotency key to prevent duplicate transactions, must a valid uuid
}

// InputListUserTransactionsParams list user transactions input
type InputListUserTransactions struct {
	UserID string
	Page   int    // page of transactions to retrieve
	Type   string // type of transactions to get 'wallet', 'credit', 'payout'
}

// InputCreditUserParams credit user input
type InputCreditUserParams struct {
	UserID  string      `json:"user_id"`
	Amount  int         `json:"amount"`
	Receipt *Receipt    `json:"receipt"`
	Notes   interface{} `json:"notes"`
}

// InputRemoveCreditParams remove credit from user input
type InputRemoveCreditParams struct {
	UserID  string      `json:"user_id"`
	Amount  int         `json:"amount"`
	Receipt *Receipt    `json:"receipt"`
	Notes   interface{} `json:"notes"`
}

// InputCreateACHPaymentParams create ACH payment input
type InputCreateACHPaymentParams struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Amount        int    `json:"amount"`
	RoutingNumber string `json:"routing_number"`
	AccountNumber string `json:"account_number"`
	AccountType   string `json:"account_type"`
	Plaid         *Plaid `json:"plaid"`
}

// CreateTransactionResponse create transaction output
type CreateTransactionResponse struct {
	Success     bool         `json:"success"`
	Message     string       `json:"message"`
	Transaction *Transaction `json:"transaction"`
}

// CreateTransactionUnverifiedUserResponse create transaction unverified user output
type CreateTransactionUnverifiedUserResponse struct {
	Success     bool         `json:"success"`
	Message     string       `json:"message"`
	Transaction *Transaction `json:"transaction"`
}

// CreateTransactionsResponse create multiple transactions output
type CreateTransactionsResponse struct {
	Success bool   `json:"success"`
	Status  string `json:"status"`   // The status of the batch request.
	BatchID string `json:"batch_id"` // batch ID for checking status. Use with /api/transactions/batch
}

// GetTransactionsBatchStatusResponse get transactions batch status output
type GetTransactionsBatchStatusResponse struct {
	Success bool   `json:"success"`
	Status  string `json:"status"`
}

// ListUserTransactionsResponse
type ListUserTransactionsResponse struct {
	Success      bool           `json:"success"`
	Transactions []*Transaction `json:"transactions"`
	Total        int            `json:"total"`
}

// GetTransactionByIDResponse
type GetTransactionByIDResponse struct {
	Success     bool         `json:"success"`
	Transaction *Transaction `json:"transaction"`
}

// CreditUserResponse
type CreditUserResponse struct {
	Success     bool         `json:"success"`
	Transaction *Transaction `json:"transaction"`
}

// RemoveCreditResponse
type RemoveCreditResponse struct {
	Success     bool         `json:"success"`
	Transaction *Transaction `json:"transaction"`
}

// CreateACHPaymentResponse
type CreateACHPaymentResponse struct {
	Success bool `json:"success"`
}

// CreateTransaction create a new transaction
func (api *API) CreateTransaction(ctx context.Context, in *InputCreateTransactionParams) (*CreateTransactionResponse, error) {

	r := api.h + "/api/transactions/create"
	bo, e := api.cl.post(ctx, r, in)
	if e != nil {
		return nil, e
	}

	var tr CreateTransactionResponse
	if e := json.Unmarshal(bo, &tr); e != nil {
		return nil, fmt.Errorf("dots api create transaction json.Unmarshal err %v html error %s", e, string(bo))
	}

	return &tr, nil
}

// CreateTransactionUnverifiedUser create a new transaction with an unverified user
func (api *API) CreateTransactionUnverifiedUser(ctx context.Context, in *InputCreateTransactionUnverifiedUserParams) (*CreateTransactionUnverifiedUserResponse, error) {

	r := api.h + "/api/transactions/create_unverified"
	bo, e := api.cl.post(ctx, r, in)
	if e != nil {
		return nil, e
	}

	var ur CreateTransactionUnverifiedUserResponse
	if e := json.Unmarshal(bo, &ur); e != nil {
		return nil, fmt.Errorf("dots api create transaction unverified user json.Unmarshal err %v html response %s", e, string(bo))
	}

	return &ur, nil
}

// CreateTransactions creates multiple transactions asynchronously errors will be delivered by email
func (api *API) CreateTransactions(ctx context.Context, in *InputCreateTransactionsParams) (*CreateTransactionsResponse, error) {

	r := api.h + "/api/transactions/create_batch"
	bo, e := api.cl.post(ctx, r, in)
	if e != nil {
		return nil, e
	}

	var ctr CreateTransactionsResponse
	if e := json.Unmarshal(bo, &ctr); e != nil {
		return nil, fmt.Errorf("dots api create transactions json.Unmarshal err %v html response %s", e, string(bo))
	}

	return &ctr, nil
}

// GetTransactionsBatchStatus gets the completion status of a transaction batch
func (api *API) GetTransactionsBatchStatus(ctx context.Context, batchID string) (*GetTransactionsBatchStatusResponse, error) {

	r := api.h + "/api/transactions/batch/" + batchID
	bo, e := api.cl.post(ctx, r, nil)
	if e != nil {
		return nil, e
	}

	var bsr GetTransactionsBatchStatusResponse
	if e := json.Unmarshal(bo, &bsr); e != nil {
		return nil, fmt.Errorf("dots api get transactions batch status json.Unmarshal err %v html response %s", e, string(bo))
	}

	return &bsr, nil
}

// ListUserTransactions get transactions for a user
func (api *API) ListUserTransactions(ctx context.Context, in *InputListUserTransactions) (*ListUserTransactionsResponse, error) {

	r := fmt.Sprintf(api.h+"/api/transactions/get/user/%s&page=%d?type=%s", in.UserID, in.Page, in.Type)
	bo, e := api.cl.get(ctx, r)
	if e != nil {
		return nil, e
	}

	var l ListUserTransactionsResponse
	if e := json.Unmarshal(bo, &l); e != nil {
		return nil, fmt.Errorf("dots api list user transactions json.Unmarshal err %v html response %s", e, string(bo))
	}

	return &l, nil
}

// GetTransactionByID get a transaction by its id
func (api *API) GetTransactionByID(ctx context.Context, transactionID string) (*GetTransactionByIDResponse, error) {

	r := fmt.Sprintf(api.h+"/api/transactions/get/transactions/%s", transactionID)
	bo, e := api.cl.get(ctx, r)
	if e != nil {
		return nil, e
	}

	var t GetTransactionByIDResponse
	if e := json.Unmarshal(bo, &t); e != nil {
		return nil, fmt.Errorf("dots api get transaction by id json.Unmarshal err %v html response %s", e, string(bo))
	}

	return &t, nil
}

// CreditUser add credit to a user
func (api *API) CreditUser(ctx context.Context, in *InputCreditUserParams) (*CreditUserResponse, error) {

	r := api.h + "/api/transactions/add_credit"
	bo, e := api.cl.post(ctx, r, in)
	if e != nil {
		return nil, e
	}

	var c CreditUserResponse
	if e := json.Unmarshal(bo, &c); e != nil {
		return nil, fmt.Errorf("dots api credit user json.Unmarshal err %v html response %s", e, string(bo))
	}

	return &c, nil
}

// RemoveCredit remote credit from a user
func (api *API) RemoveCredit(ctx context.Context, in *InputRemoveCreditParams) (*RemoveCreditResponse, error) {

	r := api.h + "/api/transactions/remove_credit"
	bo, e := api.cl.post(ctx, r, in)
	if e != nil {
		return nil, e
	}

	var c RemoveCreditResponse
	if e := json.Unmarshal(bo, &c); e != nil {
		return nil, fmt.Errorf("dots api remove credit json.Unmarshal err %v html response %s", e, string(bo))
	}

	return &c, nil
}

// CreateACHPayment create an ACH payment from a user's bank account with their Plaid information
func (api *API) CreateACHPayment(ctx context.Context, in *InputCreateACHPaymentParams) (*CreateACHPaymentResponse, error) {

	r := api.h + "/api/transactions/create_ach_payment"
	bo, e := api.cl.post(ctx, r, in)
	if e != nil {
		return nil, e
	}

	var p CreateACHPaymentResponse
	if e := json.Unmarshal(bo, &p); e != nil {
		return nil, fmt.Errorf("dots api create ach payment json.Unmarshal err %v html response %s", e, string(bo))
	}

	return &p, nil
}
