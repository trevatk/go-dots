package dots

import (
	"context"
	"encoding/json"
	"fmt"
)

// InputCreatePayoutParams
type InputCreatePayoutParams struct {
	UserID        string `json:"user_id"`
	Amount        int    `json:"amount"`
	PayoutMethod  string `json:"payout_method"`
	ACHAccountID  string `json:"ach_account_id"`
	IntlAccountID string `json:"intl_account_id"`
	Fund          bool   `json:"fund"`
}

// InputCreatePayoutLinkParams
type InputCreatePayoutLinkParams struct {
	Delivery  Delivery `json:"delivery"`
	Amount    int      `json:"amount"`
	Notes     string   `json:"notes"`
	TaxExempt bool     `json:"tax_exempt"`
	Payee     *Payee   `json:"payee,omitempty"`
}

// CreatePayoutResponse
type CreatePayoutResponse struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code"`
}

// CreatePayoutLinkResponse
type CreatePayoutLinkResponse struct {
	Success    bool        `json:"success"`
	PayoutLink *PayoutLink `json:"payout_link"`
}

// CreatePayout
func (api *API) CreatePayout(in *InputCreatePayoutParams) (*CreatePayoutResponse, error) {

	r := host + "/api/payouts/create"
	b, e := api.cl.Post(r, in)
	if e != nil {
		return nil, e
	}

	var pr CreatePayoutResponse
	if e := json.Unmarshal(b, &pr); e != nil {
		return nil, fmt.Errorf("dots api create payout json.Unmarshal err %v", e)
	}

	return &pr, nil
}

// CreatePayoutWithContext
func (api *API) CreatePayoutWithContext(ctx context.Context, in *InputCreatePayoutParams) (*CreatePayoutResponse, error) {

	r := host + "/api/payouts/create"
	b, e := api.cl.PostWithContext(ctx, r, in)
	if e != nil {
		return nil, e
	}

	var pr CreatePayoutResponse
	if e := json.Unmarshal(b, &pr); e != nil {
		return nil, fmt.Errorf("dots api create payout json.Unmarshal err %v", e)
	}

	return &pr, nil
}

// CreatePayoutLink
func (api *API) CreatePayoutLink(in *InputCreatePayoutLinkParams) (*CreatePayoutLinkResponse, error) {

	r := host + "/api/payouts/create_payout_link"
	b, e := api.cl.Post(r, in)
	if e != nil {
		return nil, e
	}

	var pl CreatePayoutLinkResponse
	if e := json.Unmarshal(b, &pl); e != nil {
		return nil, fmt.Errorf("dots api create payout link json.Unmarshal err %v", e)
	}

	return &pl, nil
}

// CreatePayoutLinkWithContext
func (api *API) CreatePayoutLinkWithContext(ctx context.Context, in *InputCreatePayoutLinkParams) (*CreatePayoutLinkResponse, error) {

	r := host + "/api/payouts/create_payout_link"
	b, e := api.cl.PostWithContext(ctx, r, in)
	if e != nil {
		return nil, e
	}

	var pl CreatePayoutLinkResponse
	if e := json.Unmarshal(b, &pl); e != nil {
		return nil, fmt.Errorf("dots api create payout link json.Unmarshal err %v", e)
	}

	return &pl, nil
}
