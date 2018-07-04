package zabbix

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"strings"
	"sync"
)

// ErrNotFound describes an empty result set for an API call.
var ErrNotFound = errors.New("No results were found matching the given search parameters")

// A Session is an authenticated Zabbix JSON-RPC API client. It must be
// initialized and connected with NewSession.
type Session struct {
	// URL of the Zabbix JSON-RPC API (ending in `/api_jsonrpc.php`).
	URL string `json:"url"`

	// Token is the cached authentication token returned by `user.login` and
	// used to authenticate all API calls in this Session.
	Token string `json:"token"`

	// APIVersion is the software version string of the connected Zabbix API.
	APIVersion string `json:"apiVersion"`

	username string

	password string

	// requests timeout
	timeout time.Duration

	// Mutex for exclusive access to session
	sync.RWMutex
}

// NewSession returns a new Session given an API connection URL and an API
// username and password.
//
// An error is returned if there was an HTTP protocol error, the API credentials
// are incorrect or if the API version is indeterminable.
//
// The authentication token returned by the Zabbix API server is cached to
// authenticate all subsequent requests in this Session.
func NewSession(url string, username string, password string, timeout time.Duration) (session *Session, err error) {
	// create session
	session = &Session{URL: url, username: username, password: password, timeout: timeout}

	// get Zabbix API version
	res, err := session.Do(NewRequest("apiinfo.version", nil))
	if err != nil {
		return nil, fmt.Errorf("Error getting Zabbix API version: %v", err)
	}

	err = res.Bind(&session.APIVersion)
	if err != nil {
		return
	}

	if err = session.Login(""); err != nil {
		return
	}

	return
}

// Login to API
func (c *Session) Login(currentToken string) error {
	c.Lock()
	defer c.Unlock()

	if currentToken != c.Token { // no need re-login
		return nil
	}

	params := map[string]string{
		"user":     c.username,
		"password": c.password,
	}

	res, err := c.Do(NewRequest("user.login", params))
	if err != nil {
		return fmt.Errorf("Error logging in to Zabbix API: %v", err)
	}

	err = res.Bind(&c.Token)

	return err
}

// Version returns the software version string of the connected Zabbix API.
func (c *Session) Version() string {
	return c.APIVersion
}

// AuthToken returns the authentication token used by this session to
// authentication all API calls.
func (c *Session) AuthToken() string {
	return c.Token
}

// Do sends a JSON-RPC request and returns an API Response, using connection
// configuration defined in the parent Session.
//
// An error is returned if there was an HTTP protocol error, a non-200 response
// is received, or if an error code is set is the JSON response body.
//
// When err is nil, resp always contains a non-nil resp.Body.
//
// Generally Get or a wrapper function will be used instead of Do.
func (c *Session) Do(req *Request) (resp *Response, err error) {
	// configure request
	if req.Method != "user.login" && req.Method != "apiinfo.version" {
		c.RLock()
		req.AuthToken = c.Token
		c.RUnlock()
	}

	// encode request as json
	b, err := json.Marshal(req)
	if err != nil {
		return
	}

	dprintf("Call [m:%s,r:%d]: %s\n", req.Method, req.RequestID, b)

	// create HTTP request
	r, err := http.NewRequest("POST", c.URL, bytes.NewReader(b))
	if err != nil {
		return
	}
	r.ContentLength = int64(len(b))
	r.Header.Add("Content-Type", "application/json-rpc")

	// send request
	client := http.Client{ Timeout: c.timeout }
	res, err := client.Do(r)
	if err != nil {
		return
	}

	defer res.Body.Close()

	// read response body
	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response: %v", err)
	}

	dprintf("Response [m:%s,r:%d]: %s\n", req.Method, req.RequestID, b)

	// map HTTP response to Response struct
	resp = &Response{
		StatusCode: res.StatusCode,
	}

	// unmarshal response body
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, fmt.Errorf("Error decoding JSON response body: %v", err)
	}

	// check for API errors
	if err = resp.Err(); err != nil {
		if resp.Error.Code == -32602 && strings.Contains(resp.Error.Data, "re-login") {
			if err = c.Login(req.AuthToken); err != nil {
				return nil, err
			}
			return c.Do(NewRequest(req.Method, req.Params))
		}
		return
	}

	return
}

// Get calls the given Zabbix API method with the given query parameters and
// unmarshals the JSON response body into the given interface.
//
// An error is return if a transport, marshalling or API error happened.
func (c *Session) Get(method string, params interface{}, v interface{}) error {
	req := NewRequest(method, params)
	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	err = resp.Bind(v)
	if err != nil {
		return err
	}

	return nil
}
