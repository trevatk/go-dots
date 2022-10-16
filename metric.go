package dots

import (
	"context"
	"encoding/json"
	"fmt"
)

// InputGetTransactionVolumeMetricsParams
type InputGetTransactionVolumeMetricsParams struct {
	Period   string `json:"period"` // enum 'day', 'week', 'month', 'year', '30d'
	Type     string `json:"type"`   // enum 'receive', 'send'
	Timezone string `json:"tz"`     // three letter timezone code
}

// GetTransactionVolumeMetricsResponse
type GetTransactionVolumeMetricsResponse struct {
	Success bool `json:"success"`
	Volume  int  `json:"volume"` // volume in cents
}

// GetTransactionVolumeMetrics retrieve filtered transaction volume metrics
func (api *API) GetTransactionVolumeMetrics(ctx context.Context, in *InputGetTransactionVolumeMetricsParams) (*GetTransactionVolumeMetricsResponse, error) {

	r := api.h + "/api/metrics/transaction/volume"
	bo, e := api.cl.get(ctx, r)
	if e != nil {
		return nil, e
	}

	var t GetTransactionVolumeMetricsResponse
	if e := json.Unmarshal(bo, &t); e != nil {
		return nil, fmt.Errorf("dots api get transaction volume metrics json.Unmarshal err %v html response %s", e, string(bo))
	}

	return &t, nil
}
