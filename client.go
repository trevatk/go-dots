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

type client struct {
	h *http.Client
	t string
}

func newClient(token string) *client {

	h := &http.Client{
		Timeout: time.Second * 5,
	}

	return &client{
		h: h,
		t: token,
	}
}

func (c *client) get(ctx context.Context, route string) ([]byte, error) {

	to, ca := context.WithTimeout(ctx, time.Second)
	defer ca()

	rq, e := http.NewRequestWithContext(to, http.MethodGet, route, nil)
	if e != nil {
		return []byte{}, fmt.Errorf("client Get http.NewRequest err %v", e)
	}

	rq.Header.Add(headerAuthorization, headerBasic+c.t)
	rq.Header.Add(headerContentType, headerAppJSON)

	rp, e := c.h.Do(rq)
	if e != nil {
		return []byte{}, fmt.Errorf("client Get client.Do err %v", e)
	}
	defer rp.Body.Close()

	bo, e := io.ReadAll(rp.Body)
	if e != nil {
		return []byte{}, fmt.Errorf("client Get io.ReadAll err %v", e)
	}

	return bo, nil
}

func (c *client) post(ctx context.Context, route string, body interface{}) ([]byte, error) {

	b, e := json.Marshal(body)
	if e != nil {
		return []byte{}, fmt.Errorf("client Post json.Marshal err %v", e)
	}

	to, ca := context.WithTimeout(ctx, time.Second*3)
	defer ca()

	rq, e := http.NewRequestWithContext(to, http.MethodPost, route, bytes.NewBuffer(b))
	if e != nil {
		return []byte{}, fmt.Errorf("client Post http.NewQuest err %v", e)
	}

	rq.Header.Add(headerAuthorization, headerBasic+c.t)
	rq.Header.Add(headerContentType, headerAppJSON)

	rp, e := c.h.Do(rq)
	if e != nil {
		return []byte{}, fmt.Errorf("client Post client.Do err %v", e)
	}
	defer rp.Body.Close()

	fmt.Println("status code: ", rp.StatusCode)

	bo, e := io.ReadAll(rp.Body)
	if e != nil {
		return []byte{}, fmt.Errorf("client Post io.ReadAll err %v", e)
	}

	return bo, nil
}
