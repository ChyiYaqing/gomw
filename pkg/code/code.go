package code

// 通用: 基本错误
// Code must start with 15xxxxx
const (
	// ErrSuccess - 200: ok.
	ErrSuccess int = iota + 1500001001

	// ErrUnknown - 500: Internal server error.
	ErrUnknown

	// ErrBind - 400: Error occurred while binding the request body to the struct.
	ErrBind

	// ErrValidation - 400: Validation failed.
	ErrValidation

	// ErrTokenInvalid - 400: Token invalid
	ErrTokenInvalid
)

// 通用: 数据库类错误
const (
	// ErrDatabase - 500: Database error.
	ErrDatabase int = iota + 1500003001
)

// 通用: 认证授权类错误
const (
	// ErrEncrypt - 401: Error occurred while encrypting the user password.
	ErrEncrypt int = iota + 1500006001

	// ErrSignatureInvalid - 401: Signature is invalid.
	ErrSignatureInvalid

	// ErrExpired - 401: Token expired.
	ErrExpired

	// ErrInvalidAuthHeader - 401: Invalid authorization header.
	ErrInvalidAuthHeader

	// ErrMissingHeader - 401: The Authorization header was empty
	ErrMissingHeader

	// ErrPasswordIncorrect - 401: Password was incorrect.
	ErrPasswordIncorrect

	// PermissionDenied - 403: Permission denied.
	ErrPermissionDenied
)
