// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// secured_service HTTP client types
//
// Command:
// $ goa gen goa.design/goa/examples/security/design -o
// $(GOPATH)/src/goa.design/goa/examples/security

package client

import (
	goa "goa.design/goa"
	securedservice "goa.design/goa/examples/security/gen/secured_service"
)

// SigninResponseBody is the type of the "secured_service" service "signin"
// endpoint HTTP response body.
type SigninResponseBody struct {
	// JWT token
	JWT *string `form:"jwt,omitempty" json:"jwt,omitempty" xml:"jwt,omitempty"`
	// API Key
	APIKey *string `form:"api_key,omitempty" json:"api_key,omitempty" xml:"api_key,omitempty"`
	// OAuth2 token
	OauthToken *string `form:"oauth_token,omitempty" json:"oauth_token,omitempty" xml:"oauth_token,omitempty"`
}

// SigninUnauthorizedResponseBody is the type of the "secured_service" service
// "signin" endpoint HTTP response body for the "unauthorized" error.
type SigninUnauthorizedResponseBody string

// SecureInvalidScopesResponseBody is the type of the "secured_service" service
// "secure" endpoint HTTP response body for the "invalid-scopes" error.
type SecureInvalidScopesResponseBody string

// SecureUnauthorizedResponseBody is the type of the "secured_service" service
// "secure" endpoint HTTP response body for the "unauthorized" error.
type SecureUnauthorizedResponseBody string

// DoublySecureInvalidScopesResponseBody is the type of the "secured_service"
// service "doubly_secure" endpoint HTTP response body for the "invalid-scopes"
// error.
type DoublySecureInvalidScopesResponseBody string

// DoublySecureUnauthorizedResponseBody is the type of the "secured_service"
// service "doubly_secure" endpoint HTTP response body for the "unauthorized"
// error.
type DoublySecureUnauthorizedResponseBody string

// AlsoDoublySecureInvalidScopesResponseBody is the type of the
// "secured_service" service "also_doubly_secure" endpoint HTTP response body
// for the "invalid-scopes" error.
type AlsoDoublySecureInvalidScopesResponseBody string

// AlsoDoublySecureUnauthorizedResponseBody is the type of the
// "secured_service" service "also_doubly_secure" endpoint HTTP response body
// for the "unauthorized" error.
type AlsoDoublySecureUnauthorizedResponseBody string

// NewSigninCredsOK builds a "secured_service" service "signin" endpoint result
// from a HTTP "OK" response.
func NewSigninCredsOK(body *SigninResponseBody) *securedservice.Creds {
	v := &securedservice.Creds{
		JWT:        *body.JWT,
		APIKey:     *body.APIKey,
		OauthToken: *body.OauthToken,
	}
	return v
}

// NewSigninUnauthorized builds a secured_service service signin endpoint
// unauthorized error.
func NewSigninUnauthorized(body SigninUnauthorizedResponseBody) securedservice.Unauthorized {
	v := securedservice.Unauthorized(body)
	return v
}

// NewSecureInvalidScopes builds a secured_service service secure endpoint
// invalid-scopes error.
func NewSecureInvalidScopes(body SecureInvalidScopesResponseBody) securedservice.InvalidScopes {
	v := securedservice.InvalidScopes(body)
	return v
}

// NewSecureUnauthorized builds a secured_service service secure endpoint
// unauthorized error.
func NewSecureUnauthorized(body SecureUnauthorizedResponseBody) securedservice.Unauthorized {
	v := securedservice.Unauthorized(body)
	return v
}

// NewDoublySecureInvalidScopes builds a secured_service service doubly_secure
// endpoint invalid-scopes error.
func NewDoublySecureInvalidScopes(body DoublySecureInvalidScopesResponseBody) securedservice.InvalidScopes {
	v := securedservice.InvalidScopes(body)
	return v
}

// NewDoublySecureUnauthorized builds a secured_service service doubly_secure
// endpoint unauthorized error.
func NewDoublySecureUnauthorized(body DoublySecureUnauthorizedResponseBody) securedservice.Unauthorized {
	v := securedservice.Unauthorized(body)
	return v
}

// NewAlsoDoublySecureInvalidScopes builds a secured_service service
// also_doubly_secure endpoint invalid-scopes error.
func NewAlsoDoublySecureInvalidScopes(body AlsoDoublySecureInvalidScopesResponseBody) securedservice.InvalidScopes {
	v := securedservice.InvalidScopes(body)
	return v
}

// NewAlsoDoublySecureUnauthorized builds a secured_service service
// also_doubly_secure endpoint unauthorized error.
func NewAlsoDoublySecureUnauthorized(body AlsoDoublySecureUnauthorizedResponseBody) securedservice.Unauthorized {
	v := securedservice.Unauthorized(body)
	return v
}

// ValidateSigninResponseBody runs the validations defined on signinResponseBody
func ValidateSigninResponseBody(body *SigninResponseBody) (err error) {
	if body.JWT == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("jwt", "body"))
	}
	if body.APIKey == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("api_key", "body"))
	}
	if body.OauthToken == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("oauth_token", "body"))
	}
	return
}
