// Test v0.10.0 7007c2ec8ccd58e0d4e9451d42e35be10140b8eb
// --
// Code generated by webrpc-gen@v0.14.0-dev with golang generator. DO NOT EDIT.
//
// webrpc-gen -schema=./schema/test.ridl -target=golang -pkg=server -server -out=./server/server.gen.go
package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
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
	return "7007c2ec8ccd58e0d4e9451d42e35be10140b8eb"
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
	Enum              *Status                      `json:"enum"`
}

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
// Server
//

type WebRPCServer interface {
	http.Handler
}

type testApiServer struct {
	TestApi
}

func NewTestApiServer(svc TestApi) WebRPCServer {
	return &testApiServer{
		TestApi: svc,
	}
}

func (s *testApiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		// In case of a panic, serve a HTTP 500 error and then panic.
		if rr := recover(); rr != nil {
			RespondWithError(w, ErrorWithCause(ErrWebrpcServerPanic, fmt.Errorf("%v", rr)))
			panic(rr)
		}
	}()

	ctx := r.Context()
	ctx = context.WithValue(ctx, HTTPResponseWriterCtxKey, w)
	ctx = context.WithValue(ctx, HTTPRequestCtxKey, r)
	ctx = context.WithValue(ctx, ServiceNameCtxKey, "TestApi")

	var handler func(ctx context.Context, w http.ResponseWriter, r *http.Request)
	switch r.URL.Path {
	case "/rpc/TestApi/GetEmpty":
		handler = s.serveGetEmptyJSON
	case "/rpc/TestApi/GetError":
		handler = s.serveGetErrorJSON
	case "/rpc/TestApi/GetOne":
		handler = s.serveGetOneJSON
	case "/rpc/TestApi/SendOne":
		handler = s.serveSendOneJSON
	case "/rpc/TestApi/GetMulti":
		handler = s.serveGetMultiJSON
	case "/rpc/TestApi/SendMulti":
		handler = s.serveSendMultiJSON
	case "/rpc/TestApi/GetComplex":
		handler = s.serveGetComplexJSON
	case "/rpc/TestApi/SendComplex":
		handler = s.serveSendComplexJSON
	case "/rpc/TestApi/GetSchemaError":
		handler = s.serveGetSchemaErrorJSON
	default:
		err := ErrorWithCause(ErrWebrpcBadRoute, fmt.Errorf("no handler for path %q", r.URL.Path))
		RespondWithError(w, err)
		return
	}

	if r.Method != "POST" {
		w.Header().Add("Allow", "POST") // RFC 9110.
		err := ErrorWithCause(ErrWebrpcBadMethod, fmt.Errorf("unsupported method %q (only POST is allowed)", r.Method))
		RespondWithError(w, err)
		return
	}

	contentType := r.Header.Get("Content-Type")
	if i := strings.Index(contentType, ";"); i >= 0 {
		contentType = contentType[:i]
	}
	contentType = strings.TrimSpace(strings.ToLower(contentType))

	switch contentType {
	case "application/json":
		handler(ctx, w, r)
	default:
		err := ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("unexpected Content-Type: %q", r.Header.Get("Content-Type")))
		RespondWithError(w, err)
	}
}

func (s *testApiServer) serveGetEmptyJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	ctx = context.WithValue(ctx, MethodNameCtxKey, "GetEmpty")

	// Call service method implementation.
	err := s.TestApi.GetEmpty(ctx)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func (s *testApiServer) serveGetErrorJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	ctx = context.WithValue(ctx, MethodNameCtxKey, "GetError")

	// Call service method implementation.
	err := s.TestApi.GetError(ctx)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func (s *testApiServer) serveGetOneJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	ctx = context.WithValue(ctx, MethodNameCtxKey, "GetOne")

	// Call service method implementation.
	ret0, err := s.TestApi.GetOne(ctx)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	respPayload := struct {
		Ret0 *Simple `json:"one"`
	}{ret0}
	respBody, err := json.Marshal(respPayload)
	if err != nil {
		RespondWithError(w, ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to marshal json response: %w", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *testApiServer) serveSendOneJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		RespondWithError(w, ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to read request data: %w", err)))
		return
	}
	defer r.Body.Close()

	reqPayload := struct {
		Arg0 *Simple `json:"one"`
	}{}
	if err := json.Unmarshal(reqBody, &reqPayload); err != nil {
		RespondWithError(w, ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to unmarshal request data: %w", err)))
		return
	}

	ctx = context.WithValue(ctx, MethodNameCtxKey, "SendOne")

	// Call service method implementation.
	err = s.TestApi.SendOne(ctx, reqPayload.Arg0)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func (s *testApiServer) serveGetMultiJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	ctx = context.WithValue(ctx, MethodNameCtxKey, "GetMulti")

	// Call service method implementation.
	ret0, ret1, ret2, err := s.TestApi.GetMulti(ctx)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	respPayload := struct {
		Ret0 *Simple `json:"one"`
		Ret1 *Simple `json:"two"`
		Ret2 *Simple `json:"three"`
	}{ret0, ret1, ret2}
	respBody, err := json.Marshal(respPayload)
	if err != nil {
		RespondWithError(w, ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to marshal json response: %w", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *testApiServer) serveSendMultiJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		RespondWithError(w, ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to read request data: %w", err)))
		return
	}
	defer r.Body.Close()

	reqPayload := struct {
		Arg0 *Simple `json:"one"`
		Arg1 *Simple `json:"two"`
		Arg2 *Simple `json:"three"`
	}{}
	if err := json.Unmarshal(reqBody, &reqPayload); err != nil {
		RespondWithError(w, ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to unmarshal request data: %w", err)))
		return
	}

	ctx = context.WithValue(ctx, MethodNameCtxKey, "SendMulti")

	// Call service method implementation.
	err = s.TestApi.SendMulti(ctx, reqPayload.Arg0, reqPayload.Arg1, reqPayload.Arg2)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func (s *testApiServer) serveGetComplexJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	ctx = context.WithValue(ctx, MethodNameCtxKey, "GetComplex")

	// Call service method implementation.
	ret0, err := s.TestApi.GetComplex(ctx)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	respPayload := struct {
		Ret0 *Complex `json:"complex"`
	}{ret0}
	respBody, err := json.Marshal(respPayload)
	if err != nil {
		RespondWithError(w, ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to marshal json response: %w", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *testApiServer) serveSendComplexJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		RespondWithError(w, ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to read request data: %w", err)))
		return
	}
	defer r.Body.Close()

	reqPayload := struct {
		Arg0 *Complex `json:"complex"`
	}{}
	if err := json.Unmarshal(reqBody, &reqPayload); err != nil {
		RespondWithError(w, ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to unmarshal request data: %w", err)))
		return
	}

	ctx = context.WithValue(ctx, MethodNameCtxKey, "SendComplex")

	// Call service method implementation.
	err = s.TestApi.SendComplex(ctx, reqPayload.Arg0)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func (s *testApiServer) serveGetSchemaErrorJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		RespondWithError(w, ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to read request data: %w", err)))
		return
	}
	defer r.Body.Close()

	reqPayload := struct {
		Arg0 int `json:"code"`
	}{}
	if err := json.Unmarshal(reqBody, &reqPayload); err != nil {
		RespondWithError(w, ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to unmarshal request data: %w", err)))
		return
	}

	ctx = context.WithValue(ctx, MethodNameCtxKey, "GetSchemaError")

	// Call service method implementation.
	err = s.TestApi.GetSchemaError(ctx, reqPayload.Arg0)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func RespondWithError(w http.ResponseWriter, err error) {
	rpcErr, ok := err.(WebRPCError)
	if !ok {
		rpcErr = ErrorWithCause(ErrWebrpcEndpoint, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(rpcErr.HTTPStatus)

	respBody, _ := json.Marshal(rpcErr)
	w.Write(respBody)
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
	HTTPResponseWriterCtxKey = &contextKey{"HTTPResponseWriter"}

	HTTPRequestCtxKey = &contextKey{"HTTPRequest"}

	ServiceNameCtxKey = &contextKey{"ServiceName"}

	MethodNameCtxKey = &contextKey{"MethodName"}
)

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
	if rpcErr, ok := target.(WebRPCError); ok {
		return rpcErr.Code == e.Code
	}
	return errors.Is(e.cause, target)
}

func (e WebRPCError) Unwrap() error {
	return e.cause
}

func ErrorWithCause(rpcErr WebRPCError, cause error) WebRPCError {
	err := rpcErr
	err.cause = cause
	err.Cause = cause.Error()
	return err
}

// Webrpc errors
var (
	ErrWebrpcEndpoint      = WebRPCError{Code: 0, Name: "WebrpcEndpoint", Message: "endpoint error", HTTPStatus: 400}
	ErrWebrpcRequestFailed = WebRPCError{Code: -1, Name: "WebrpcRequestFailed", Message: "request failed", HTTPStatus: 400}
	ErrWebrpcBadRoute      = WebRPCError{Code: -2, Name: "WebrpcBadRoute", Message: "bad route", HTTPStatus: 404}
	ErrWebrpcBadMethod     = WebRPCError{Code: -3, Name: "WebrpcBadMethod", Message: "bad method", HTTPStatus: 405}
	ErrWebrpcBadRequest    = WebRPCError{Code: -4, Name: "WebrpcBadRequest", Message: "bad request", HTTPStatus: 400}
	ErrWebrpcBadResponse   = WebRPCError{Code: -5, Name: "WebrpcBadResponse", Message: "bad response", HTTPStatus: 500}
	ErrWebrpcServerPanic   = WebRPCError{Code: -6, Name: "WebrpcServerPanic", Message: "server panic", HTTPStatus: 500}
	ErrWebrpcInternalError = WebRPCError{Code: -7, Name: "WebrpcInternalError", Message: "internal error", HTTPStatus: 500}
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
