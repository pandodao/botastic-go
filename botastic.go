package botastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	appID     string
	appSecret string
	host      string
	debug     bool
	logger    Logger
}

type Logger interface {
	Debugf(format string, args ...interface{})
}

type Option func(*Client)

func WithHost(host string) Option {
	return func(c *Client) {
		if host != "" {
			c.host = strings.TrimRight(host, "/") + "/api"
		}
	}
}

func WithDebug(debug bool) Option {
	return func(c *Client) {
		c.debug = debug
	}
}

func WithLogger(logger Logger) Option {
	return func(c *Client) {
		c.logger = logger
	}
}

func New(appID, appSecret string, opts ...Option) *Client {
	c := &Client{
		appID:     appID,
		appSecret: appSecret,
		host:      "https://botastic-api.pando.im/api",
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Client) debugLog(v string, args ...interface{}) {
	if !c.debug {
		return
	}
	if c.logger != nil {
		c.logger.Debugf(v, args...)
	} else {
		log.Printf(v, args...)
	}
}

func (c *Client) request(ctx context.Context, method string, uri string, body, result any) error {
	reqLog := fmt.Sprintf("[Request] %s %s", method, uri)
	start := time.Now()
	var r io.Reader
	if body != nil {
		data, _ := json.Marshal(body)
		reqLog += fmt.Sprintf(" %s", string(data))
		r = bytes.NewBuffer(data)
	}
	c.debugLog(reqLog)

	req, err := http.NewRequestWithContext(ctx, method, c.host+uri, r)
	if err != nil {
		return err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("X-BOTASTIC-APPID", c.appID)
	req.Header.Set("X-BOTASTIC-SECRET", c.appSecret)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	c.debugLog("[Response %s] %s %s %s", time.Since(start), method, uri, string(respData))

	var res struct {
		Error
		Data any `json:"data"`
	}
	res.Data = result
	res.StatusCode = resp.StatusCode

	if err := json.Unmarshal(respData, &res); err != nil {
		res.Msg = string(respData)
		return res.Error
	}
	if res.Code != 0 {
		return res.Error
	}

	return nil
}
