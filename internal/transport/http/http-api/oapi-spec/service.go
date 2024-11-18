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
	Accrual    *int        `json:"accrual,omitempty"`
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
