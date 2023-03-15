package botastic

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type CreateIndexesItem struct {
	ObjectID   string `json:"object_id"`
	Category   string `json:"category"`
	Data       string `json:"data"`
	Properties string `json:"properties"`
}

type CreateIndexesRequest struct {
	Items []*CreateIndexesItem `json:"items"`
}

type SearchIndexesRequest struct {
	Keywords string
	N        int
}

type Index struct {
	Data       string  `json:"data"`
	ObjectID   string  `json:"object_id"`
	Category   string  `json:"category"`
	Properties string  `json:"properties"`
	CreatedAt  int64   `json:"created_at"`
	Score      float32 `json:"score"`
}

type SearchIndexesResponse struct {
	Items []*Index `json:"items"`
}

type Error struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}

func (e Error) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code, e.Msg)
}

func (c *Client) CreateIndexes(ctx context.Context, req CreateIndexesRequest) error {
	return c.request(ctx, http.MethodPost, "/indexes", req, nil)
}

func (c *Client) DeleteIndex(ctx context.Context, objectId string) error {
	return c.request(ctx, http.MethodDelete, "/indexes/"+objectId, nil, nil)
}

func (c *Client) SearchIndexes(ctx context.Context, req SearchIndexesRequest) (*SearchIndexesResponse, error) {
	values := url.Values{}
	values.Add("keywords", req.Keywords)
	if req.N != 0 {
		values.Add("n", strconv.Itoa(req.N))
	}

	result := &SearchIndexesResponse{}
	if err := c.request(ctx, http.MethodGet, "/indexes/search?"+values.Encode(), nil, &result.Items); err != nil {
		return nil, err
	}
	return result, nil
}
