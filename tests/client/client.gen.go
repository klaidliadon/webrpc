// Test v0.10.0 1e00d2fe909b6626e4156ac2a40ff08223d75224
// --
// Code generated by webrpc-gen@v0.15.0-dev with golang generator. DO NOT EDIT.
//
// webrpc-gen -schema=./schema/test.ridl -target=golang -pkg=client -client -out=./client/client.gen.go
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// WebRPC description and code-gen version
func WebRPCVersion() string {
	return "v1"
}

// Schema version of your RIDL schema
func WebRPCSchemaVersion() string {
	return "v0.10.0"
}

// Schema hash generated from your RIDL schema
func WebRPCSchemaHash() string {
	return "1e00d2fe909b6626e4156ac2a40ff08223d75224"
}

//
// Types
//

type Status uint32

const (
	Status_AVAILABLE     Status = 0
	Status_NOT_AVAILABLE Status = 1
)

var Status_name = map[uint32]string{
	0: "AVAILABLE",
	1: "NOT_AVAILABLE",
}

var Status_value = map[string]uint32{
	"AVAILABLE":     0,
	"NOT_AVAILABLE": 1,
}

func (x Status) String() string {
	return Status_name[uint32(x)]
}

func (x Status) MarshalText() ([]byte, error) {
	return []byte(Status_name[uint32(x)]), nil
}

func (x *Status) UnmarshalText(b []byte) error {
	*x = Status(Status_value[string(b)])
	return nil
}

func (x *Status) Is(values ...Status) bool {
	if x == nil {
		return false
	}
	for _, v := range values {
		if *x == v {
			return true
		}
	}
	return false
}

type Simple struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID       uint64 `json:"id" db:"id"`
	Username string `json:"USERNAME" db:"username"`
	Role     string `json:"role" db:"-"`
}

type Complex struct {
	Meta              map[string]interface{}       `json:"meta"`
	MetaNestedExample map[string]map[string]uint32 `json:"metaNestedExample"`
	NamesList         []string                     `json:"namesList"`
	NumsList          []int64                      `json:"numsList"`
	DoubleArray       [][]string                   `json:"doubleArray"`
	ListOfMaps        []map[string]uint32          `json:"listOfMaps"`
	ListOfUsers       []*User                      `json:"listOfUsers"`
	MapOfUsers        map[string]*User             `json:"mapOfUsers"`
	User              *User                        `json:"user"`
	Enum              Status                       `json:"enum"`
}

var WebRPCServices = map[string][]string{
	"TestApi": {
		"GetEmpty",
		"GetError",
		"GetOne",
		"SendOne",
		"GetMulti",
		"SendMulti",
		"GetComplex",
		"SendComplex",
		"GetSchemaError",
	},
}

//
// Server types
//

type TestApi interface {
	GetEmpty(ctx context.Context) error
	GetError(ctx context.Context) error
	GetOne(ctx context.Context) (*Simple, error)
	SendOne(ctx context.Context, one *Simple) error
	GetMulti(ctx context.Context) (*Simple, *Simple, *Simple, error)
	SendMulti(ctx context.Context, one *Simple, two *Simple, three *Simple) error
	GetComplex(ctx context.Context) (*Complex, error)
	SendComplex(ctx context.Context, complex *Complex) error
	GetSchemaError(ctx context.Context, code int) error
}

//
// Client types
//

type TestApiClient interface {
	GetEmpty(ctx context.Context) error
	GetError(ctx context.Context) error
	GetOne(ctx context.Context) (*Simple, error)
	SendOne(ctx context.Context, one *Simple) error
	GetMulti(ctx context.Context) (*Simple, *Simple, *Simple, error)
	SendMulti(ctx context.Context, one *Simple, two *Simple, three *Simple) error
	GetComplex(ctx context.Context) (*Complex, error)
	SendComplex(ctx context.Context, complex *Complex) error
	GetSchemaError(ctx context.Context, code int) error
}

//
// Client
//

const TestApiPathPrefix = "/rpc/TestApi/"

type testApiClient struct {
	client HTTPClient
	urls   [9]string
}

func NewTestApiClient(addr string, client HTTPClient) TestApiClient {
	prefix := urlBase(addr) + TestApiPathPrefix
	urls := [9]string{
		prefix + "GetEmpty",
		prefix + "GetError",
		prefix + "GetOne",
		prefix + "SendOne",
		prefix + "GetMulti",
		prefix + "SendMulti",
		prefix + "GetComplex",
		prefix + "SendComplex",
		prefix + "GetSchemaError",
	}
	return &testApiClient{
		client: client,
		urls:   urls,
	}
}

func (c *testApiClient) GetEmpty(ctx context.Context) error {

	resp, err := doHTTPRequest(ctx, c.client, c.urls[0], nil, nil)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return err
}

func (c *testApiClient) GetError(ctx context.Context) error {

	resp, err := doHTTPRequest(ctx, c.client, c.urls[1], nil, nil)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return err
}

func (c *testApiClient) GetOne(ctx context.Context) (*Simple, error) {
	out := struct {
		Ret0 *Simple `json:"one"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[2], nil, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return out.Ret0, err
}

func (c *testApiClient) SendOne(ctx context.Context, one *Simple) error {
	in := struct {
		Arg0 *Simple `json:"one"`
	}{one}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[3], in, nil)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return err
}

func (c *testApiClient) GetMulti(ctx context.Context) (*Simple, *Simple, *Simple, error) {
	out := struct {
		Ret0 *Simple `json:"one"`
		Ret1 *Simple `json:"two"`
		Ret2 *Simple `json:"three"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[4], nil, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return out.Ret0, out.Ret1, out.Ret2, err
}

func (c *testApiClient) SendMulti(ctx context.Context, one *Simple, two *Simple, three *Simple) error {
	in := struct {
		Arg0 *Simple `json:"one"`
		Arg1 *Simple `json:"two"`
		Arg2 *Simple `json:"three"`
	}{one, two, three}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[5], in, nil)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return err
}

func (c *testApiClient) GetComplex(ctx context.Context) (*Complex, error) {
	out := struct {
		Ret0 *Complex `json:"complex"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[6], nil, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return out.Ret0, err
}

func (c *testApiClient) SendComplex(ctx context.Context, complex *Complex) error {
	in := struct {
		Arg0 *Complex `json:"complex"`
	}{complex}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[7], in, nil)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return err
}

func (c *testApiClient) GetSchemaError(ctx context.Context, code int) error {
	in := struct {
		Arg0 int `json:"code"`
	}{code}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[8], in, nil)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return err
}

// HTTPClient is the interface used by generated clients to send HTTP requests.
// It is fulfilled by *(net/http).Client, which is sufficient for most users.
// Users can provide their own implementation for special retry policies.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// urlBase helps ensure that addr specifies a scheme. If it is unparsable
// as a URL, it returns addr unchanged.
func urlBase(addr string) string {
	// If the addr specifies a scheme, use it. If not, default to
	// http. If url.Parse fails on it, return it unchanged.
	url, err := url.Parse(addr)
	if err != nil {
		return addr
	}
	if url.Scheme == "" {
		url.Scheme = "http"
	}
	return url.String()
}

// newRequest makes an http.Request from a client, adding common headers.
func newRequest(ctx context.Context, url string, reqBody io.Reader, contentType string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", contentType)
	req.Header.Set("Content-Type", contentType)
	if headers, ok := HTTPRequestHeaders(ctx); ok {
		for k := range headers {
			for _, v := range headers[k] {
				req.Header.Add(k, v)
			}
		}
	}
	return req, nil
}

// doHTTPRequest is common code to make a request to the remote service.
func doHTTPRequest(ctx context.Context, client HTTPClient, url string, in, out interface{}) (*http.Response, error) {
	reqBody, err := json.Marshal(in)
	if err != nil {
		return nil, ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to marshal JSON body: %w", err))
	}
	if err = ctx.Err(); err != nil {
		return nil, ErrWebrpcRequestFailed.WithCause(fmt.Errorf("aborted because context was done: %w", err))
	}

	req, err := newRequest(ctx, url, bytes.NewBuffer(reqBody), "application/json")
	if err != nil {
		return nil, ErrWebrpcRequestFailed.WithCause(fmt.Errorf("could not build request: %w", err))
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, ErrWebrpcRequestFailed.WithCause(err)
	}

	if resp.StatusCode != 200 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to read server error response body: %w", err))
		}

		var rpcErr WebRPCError
		if err := json.Unmarshal(respBody, &rpcErr); err != nil {
			return nil, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to unmarshal server error: %w", err))
		}
		if rpcErr.Cause != "" {
			rpcErr.cause = errors.New(rpcErr.Cause)
		}
		return nil, rpcErr
	}

	if out != nil {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to read response body: %w", err))
		}

		err = json.Unmarshal(respBody, &out)
		if err != nil {
			return nil, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to unmarshal JSON response body: %w", err))
		}
	}

	return resp, nil
}

func WithHTTPRequestHeaders(ctx context.Context, h http.Header) (context.Context, error) {
	if _, ok := h["Accept"]; ok {
		return nil, errors.New("provided header cannot set Accept")
	}
	if _, ok := h["Content-Type"]; ok {
		return nil, errors.New("provided header cannot set Content-Type")
	}

	copied := make(http.Header, len(h))
	for k, vv := range h {
		if vv == nil {
			copied[k] = nil
			continue
		}
		copied[k] = make([]string, len(vv))
		copy(copied[k], vv)
	}

	return context.WithValue(ctx, HTTPClientRequestHeadersCtxKey, copied), nil
}

func HTTPRequestHeaders(ctx context.Context) (http.Header, bool) {
	h, ok := ctx.Value(HTTPClientRequestHeadersCtxKey).(http.Header)
	return h, ok
}

//
// Helpers
//

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "webrpc context value " + k.name
}

var (
	HTTPClientRequestHeadersCtxKey = &contextKey{"HTTPClientRequestHeaders"}
	HTTPRequestCtxKey              = &contextKey{"HTTPRequest"}

	ServiceNameCtxKey = &contextKey{"ServiceName"}

	MethodNameCtxKey = &contextKey{"MethodName"}
)

func ServiceNameFromContext(ctx context.Context) string {
	service, _ := ctx.Value(ServiceNameCtxKey).(string)
	return service
}

func MethodNameFromContext(ctx context.Context) string {
	method, _ := ctx.Value(MethodNameCtxKey).(string)
	return method
}

func RequestFromContext(ctx context.Context) *http.Request {
	r, _ := ctx.Value(HTTPRequestCtxKey).(*http.Request)
	return r
}

//
// Errors
//

type WebRPCError struct {
	Name       string `json:"error"`
	Code       int    `json:"code"`
	Message    string `json:"msg"`
	Cause      string `json:"cause,omitempty"`
	HTTPStatus int    `json:"status"`
	cause      error
}

var _ error = WebRPCError{}

func (e WebRPCError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s %d: %s: %v", e.Name, e.Code, e.Message, e.cause)
	}
	return fmt.Sprintf("%s %d: %s", e.Name, e.Code, e.Message)
}

func (e WebRPCError) Is(target error) bool {
	if target == nil {
		return false
	}
	if rpcErr, ok := target.(WebRPCError); ok {
		return rpcErr.Code == e.Code
	}
	return errors.Is(e.cause, target)
}

func (e WebRPCError) Unwrap() error {
	return e.cause
}

func (e WebRPCError) WithCause(cause error) WebRPCError {
	err := e
	err.cause = cause
	err.Cause = cause.Error()
	return err
}

// Deprecated: Use .WithCause() method on WebRPCError.
func ErrorWithCause(rpcErr WebRPCError, cause error) WebRPCError {
	return rpcErr.WithCause(cause)
}

// Webrpc errors
var (
	ErrWebrpcEndpoint           = WebRPCError{Code: 0, Name: "WebrpcEndpoint", Message: "endpoint error", HTTPStatus: 400}
	ErrWebrpcRequestFailed      = WebRPCError{Code: -1, Name: "WebrpcRequestFailed", Message: "request failed", HTTPStatus: 400}
	ErrWebrpcBadRoute           = WebRPCError{Code: -2, Name: "WebrpcBadRoute", Message: "bad route", HTTPStatus: 404}
	ErrWebrpcBadMethod          = WebRPCError{Code: -3, Name: "WebrpcBadMethod", Message: "bad method", HTTPStatus: 405}
	ErrWebrpcBadRequest         = WebRPCError{Code: -4, Name: "WebrpcBadRequest", Message: "bad request", HTTPStatus: 400}
	ErrWebrpcBadResponse        = WebRPCError{Code: -5, Name: "WebrpcBadResponse", Message: "bad response", HTTPStatus: 500}
	ErrWebrpcServerPanic        = WebRPCError{Code: -6, Name: "WebrpcServerPanic", Message: "server panic", HTTPStatus: 500}
	ErrWebrpcInternalError      = WebRPCError{Code: -7, Name: "WebrpcInternalError", Message: "internal error", HTTPStatus: 500}
	ErrWebrpcClientDisconnected = WebRPCError{Code: -8, Name: "WebrpcClientDisconnected", Message: "client disconnected", HTTPStatus: 400}
	ErrWebrpcStreamLost         = WebRPCError{Code: -9, Name: "WebrpcStreamLost", Message: "stream lost", HTTPStatus: 400}
	ErrWebrpcStreamFinished     = WebRPCError{Code: -10, Name: "WebrpcStreamFinished", Message: "stream finished", HTTPStatus: 200}
)

// Schema errors
var (
	ErrUnauthorized    = WebRPCError{Code: 1, Name: "Unauthorized", Message: "unauthorized", HTTPStatus: 401}
	ErrExpiredToken    = WebRPCError{Code: 2, Name: "ExpiredToken", Message: "expired token", HTTPStatus: 401}
	ErrInvalidToken    = WebRPCError{Code: 3, Name: "InvalidToken", Message: "invalid token", HTTPStatus: 401}
	ErrDeactivated     = WebRPCError{Code: 4, Name: "Deactivated", Message: "account deactivated", HTTPStatus: 403}
	ErrConfirmAccount  = WebRPCError{Code: 5, Name: "ConfirmAccount", Message: "confirm your email", HTTPStatus: 403}
	ErrAccessDenied    = WebRPCError{Code: 6, Name: "AccessDenied", Message: "access denied", HTTPStatus: 403}
	ErrMissingArgument = WebRPCError{Code: 7, Name: "MissingArgument", Message: "missing argument", HTTPStatus: 400}
	ErrUnexpectedValue = WebRPCError{Code: 8, Name: "UnexpectedValue", Message: "unexpected value", HTTPStatus: 400}
	ErrRateLimited     = WebRPCError{Code: 100, Name: "RateLimited", Message: "too many requests", HTTPStatus: 429}
	ErrDatabaseDown    = WebRPCError{Code: 101, Name: "DatabaseDown", Message: "service outage", HTTPStatus: 503}
	ErrElasticDown     = WebRPCError{Code: 102, Name: "ElasticDown", Message: "search is degraded", HTTPStatus: 503}
	ErrNotImplemented  = WebRPCError{Code: 103, Name: "NotImplemented", Message: "not implemented", HTTPStatus: 501}
	ErrUserNotFound    = WebRPCError{Code: 200, Name: "UserNotFound", Message: "user not found", HTTPStatus: 400}
	ErrUserBusy        = WebRPCError{Code: 201, Name: "UserBusy", Message: "user busy", HTTPStatus: 400}
	ErrInvalidUsername = WebRPCError{Code: 202, Name: "InvalidUsername", Message: "invalid username", HTTPStatus: 400}
	ErrFileTooBig      = WebRPCError{Code: 300, Name: "FileTooBig", Message: "file is too big (max 1GB)", HTTPStatus: 400}
	ErrFileInfected    = WebRPCError{Code: 301, Name: "FileInfected", Message: "file is infected", HTTPStatus: 400}
	ErrFileType        = WebRPCError{Code: 302, Name: "FileType", Message: "unsupported file type", HTTPStatus: 400}
)
