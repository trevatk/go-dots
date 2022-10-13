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

// InputCreateFlowParams
type InputCreateFlowParams struct {
	Steps  []string `json:"steps"`             // array of strings enum: 'compliance', 'manage-payements', 'manage-payouts', 'payout'
	UserID *string  `json:"user_id,omitempty"` // string <uuid>
}

// InputGetFlowParams
type InputGetFlowParams struct {
	FlowID string `json:"flow_id"` // string <uuid>
}

// Flow
type Flow struct {
	ID             string   `json:"id"`              // string <uuid>
	Created        string   `json:"created"`         // string <date-time>
	Updated        string   `json:"updated"`         // string <date-time>
	UserID         string   `json:"user_id"`         // string <uuid>
	Steps          []string `json:"steps"`           // array of strings enum: 'compliance', 'manage-payments', 'manage-payouts', 'payout'
	CompletedSteps []string `json:"completed_steps"` // array of strings enum: 'compliance', 'manage-payments', 'manage-payouts', 'payout'
	PayoutLinkID   string   `json:"payout_link_id"`  // string <uuid>
	Link           string   `json:"link"`            // string <url>
}

// FlowResponse
type FlowResponse struct {
	Success bool  `json:"success"`
	Flow    *Flow `json:"flow"`
}

// CreateFlow
func (api *API) CreateFlow(in *InputCreateFlowParams) (*Flow, error) {

	b, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("dots api create flow json.Marshal err %v", err)
	}

	r := host + "/api/flow/create"
	rq, err := http.NewRequest(http.MethodPost, r, bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("dots api create flow http.NewRequest err %v", err)
	}

	rq.Header.Add(headerAuthorization, headerBasic+api.token)
	rq.Header.Add(headerContentType, headerAppJSON)

	rp, err := api.cl.Do(rq)
	if err != nil {
		return nil, fmt.Errorf("dots api create flow client.Do err %v", err)
	}
	defer rp.Body.Close()

	bo, err := io.ReadAll(rp.Body)
	if err != nil {
		return nil, fmt.Errorf("dots api create flow io.Readall err %v", err)
	}

	var fr FlowResponse
	if err := json.Unmarshal(bo, &fr); err != nil {
		return nil, fmt.Errorf("dots api create flow json.Unmarshal err %v", err)
	}

	if fr.Success == false {

	}

	return fr.Flow, nil
}

// CreateFlowWithContext
func (api *API) CreateFlowWithContext(ctx context.Context, in *InputCreateFlowParams) (*Flow, error) {

	b, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("dots api create flow json.Marshal err %v", err)
	}

	to, ca := context.WithTimeout(ctx, time.Second*3)
	defer ca()

	r := host + "/api/flow/create"
	rq, err := http.NewRequestWithContext(to, http.MethodPost, r, bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("dots api create flow http.NewRequest err %v", err)
	}

	rq.Header.Add(headerAuthorization, headerBasic+api.token)
	rq.Header.Add(headerContentType, headerAppJSON)

	rp, err := api.cl.Do(rq)
	if err != nil {
		return nil, fmt.Errorf("dots api create flow client.Do err %v", err)
	}
	defer rp.Body.Close()

	bo, err := io.ReadAll(rp.Body)
	if err != nil {
		return nil, fmt.Errorf("dots api create flow io.Readall err %v", err)
	}

	var fr FlowResponse
	if err := json.Unmarshal(bo, &fr); err != nil {
		return nil, fmt.Errorf("dots api create flow json.Unmarshal err %v", err)
	}

	return fr.Flow, nil
}

// GetFlowByID
func (api *API) GetFlowByID(in *InputGetFlowParams) (*Flow, error) {

	r := host + "/api/flow/get/" + in.FlowID
	rq, err := http.NewRequest(http.MethodGet, r, nil)
	if err != nil {
		return nil, fmt.Errorf("dots api GetFlowByID http.NewRequest err %v", err)
	}

	rq.Header.Add(headerAuthorization, headerBasic+api.token)
	rq.Header.Add(headerContentType, headerAppJSON)

	rp, err := api.cl.Do(rq)
	if err != nil {
		return nil, fmt.Errorf("dots api create flow client.Do err %v", err)
	}
	defer rp.Body.Close()

	bo, err := io.ReadAll(rp.Body)
	if err != nil {
		return nil, fmt.Errorf("dots api create flow io.Readall err %v", err)
	}

	var fr FlowResponse
	if err := json.Unmarshal(bo, &fr); err != nil {
		return nil, fmt.Errorf("dots api create flow json.Unmarshal err %v", err)
	}

	return fr.Flow, nil
}

// GetFlowByIDWithContext
func (api *API) GetFlowByIdWithContext(ctx context.Context, in *InputGetFlowParams) (*Flow, error) {

	to, ca := context.WithTimeout(ctx, time.Second*3)
	defer ca()

	r := host + "/api/flow/get/" + in.FlowID
	rq, err := http.NewRequestWithContext(to, http.MethodGet, r, nil)
	if err != nil {
		return nil, fmt.Errorf("dots api create flow http.NewRequest err %v", err)
	}

	rq.Header.Add(headerAuthorization, headerBasic+api.token)
	rq.Header.Add(headerContentType, headerAppJSON)

	rp, err := api.cl.Do(rq)
	if err != nil {
		return nil, fmt.Errorf("dots api create flow client.Do err %v", err)
	}
	defer rp.Body.Close()

	bo, err := io.ReadAll(rp.Body)
	if err != nil {
		return nil, fmt.Errorf("dots api create flow io.Readall err %v", err)
	}

	var fr FlowResponse
	if err := json.Unmarshal(bo, &fr); err != nil {
		return nil, fmt.Errorf("dots api create flow json.Unmarshal err %v", err)
	}

	return fr.Flow, nil
}
