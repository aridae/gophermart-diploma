//go:build go1.22

// Package apispec provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package apispec

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for OrderStatus.
const (
	INVALID    OrderStatus = "INVALID"
	NEW        OrderStatus = "NEW"
	PROCESSED  OrderStatus = "PROCESSED"
	PROCESSING OrderStatus = "PROCESSING"
)

// Balance defines model for Balance.
type Balance struct {
	Current   float32 `json:"current"`
	Withdrawn float32 `json:"withdrawn"`
}

// Order defines model for Order.
type Order struct {
	Accrual    *float32    `json:"accrual,omitempty"`
	Number     string      `json:"number"`
	Status     OrderStatus `json:"status"`
	UploadedAt time.Time   `json:"uploaded_at"`
}

// OrderStatus defines model for Order.Status.
type OrderStatus string

// Withdrawal defines model for Withdrawal.
type Withdrawal struct {
	Order       string    `json:"order"`
	ProcessedAt time.Time `json:"processed_at"`
	Sum         float32   `json:"sum"`
}

// PostUserBalanceJSONBody defines parameters for PostUserBalance.
type PostUserBalanceJSONBody struct {
	Order string  `json:"order"`
	Sum   float32 `json:"sum"`
}

// PostUserLoginJSONBody defines parameters for PostUserLogin.
type PostUserLoginJSONBody struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// PostUserOrdersTextBody defines parameters for PostUserOrders.
type PostUserOrdersTextBody = string

// PostUserRegisterJSONBody defines parameters for PostUserRegister.
type PostUserRegisterJSONBody struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// PostUserBalanceJSONRequestBody defines body for PostUserBalance for application/json ContentType.
type PostUserBalanceJSONRequestBody PostUserBalanceJSONBody

// PostUserLoginJSONRequestBody defines body for PostUserLogin for application/json ContentType.
type PostUserLoginJSONRequestBody PostUserLoginJSONBody

// PostUserOrdersTextRequestBody defines body for PostUserOrders for text/plain ContentType.
type PostUserOrdersTextRequestBody = PostUserOrdersTextBody

// PostUserRegisterJSONRequestBody defines body for PostUserRegister for application/json ContentType.
type PostUserRegisterJSONRequestBody PostUserRegisterJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get current balance
	// (GET /user/balance)
	GetUserBalance(w http.ResponseWriter, r *http.Request)
	// Request withdrawal
	// (POST /user/balance)
	PostUserBalance(w http.ResponseWriter, r *http.Request)
	// Authenticate an existing user
	// (POST /user/login)
	PostUserLogin(w http.ResponseWriter, r *http.Request)
	// Get list of submitted orders
	// (GET /user/orders)
	GetUserOrders(w http.ResponseWriter, r *http.Request)
	// Submit order number for processing
	// (POST /user/orders)
	PostUserOrders(w http.ResponseWriter, r *http.Request)
	// Register a new user
	// (POST /user/register)
	PostUserRegister(w http.ResponseWriter, r *http.Request)
	// Get withdrawal history
	// (GET /user/withdrawals)
	GetUserWithdrawals(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetUserBalance operation middleware
func (siw *ServerInterfaceWrapper) GetUserBalance(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUserBalance(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostUserBalance operation middleware
func (siw *ServerInterfaceWrapper) PostUserBalance(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostUserBalance(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostUserLogin operation middleware
func (siw *ServerInterfaceWrapper) PostUserLogin(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostUserLogin(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetUserOrders operation middleware
func (siw *ServerInterfaceWrapper) GetUserOrders(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUserOrders(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostUserOrders operation middleware
func (siw *ServerInterfaceWrapper) PostUserOrders(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostUserOrders(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostUserRegister operation middleware
func (siw *ServerInterfaceWrapper) PostUserRegister(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostUserRegister(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetUserWithdrawals operation middleware
func (siw *ServerInterfaceWrapper) GetUserWithdrawals(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUserWithdrawals(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{})
}

// ServeMux is an abstraction of http.ServeMux.
type ServeMux interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type StdHTTPServerOptions struct {
	BaseURL          string
	BaseRouter       ServeMux
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, m ServeMux) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseRouter: m,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, m ServeMux, baseURL string) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseURL:    baseURL,
		BaseRouter: m,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options StdHTTPServerOptions) http.Handler {
	m := options.BaseRouter

	if m == nil {
		m = http.NewServeMux()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	m.HandleFunc("GET "+options.BaseURL+"/user/balance", wrapper.GetUserBalance)
	m.HandleFunc("POST "+options.BaseURL+"/user/balance", wrapper.PostUserBalance)
	m.HandleFunc("POST "+options.BaseURL+"/user/login", wrapper.PostUserLogin)
	m.HandleFunc("GET "+options.BaseURL+"/user/orders", wrapper.GetUserOrders)
	m.HandleFunc("POST "+options.BaseURL+"/user/orders", wrapper.PostUserOrders)
	m.HandleFunc("POST "+options.BaseURL+"/user/register", wrapper.PostUserRegister)
	m.HandleFunc("GET "+options.BaseURL+"/user/withdrawals", wrapper.GetUserWithdrawals)

	return m
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
