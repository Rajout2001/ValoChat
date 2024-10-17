package services

import (
	db "pc3r/prisma/db"
)

type SubscribeNewsletterSchema struct {
	Email string `json:"email"`
}

type NewsletterResponse struct {
	Code   string    `json:"code"`
	Message string `json:"message"`
}

// Ce découpage été inspiré par Tabellout Yanis.
type LoginSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginRes struct {
	User   UserRes    `json:"user"`
	Tokens AuthTokens `json:"tokens"`
}

type UserRes struct {
	*db.UserModel
}

type SignupSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
type SignupRes struct {
	Success bool   `json:"success"`
	Id      string `json:"id"`
	Message string `json:"message"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type AuthTokens struct {
	Access string `json:"access"`
}

type CtxAuthKey struct{}

type Error string

const (
	BAD_REQUEST           Error = "400 Bad Request"
	UNAUTHORIZED                = "401 UNAUTHORIZED"
	NOT_FOUND                   = "404 NOT_FOUND"
	INPUT_ERROR                 = "422 INPUT_ERROR"
	INTERNAL_SERVER_ERROR       = "500 INTERNAL_SERVER_ERROR"
)

type HTTPError struct {
	Message string `json:"message"`
	Code    Error  `json:"code"`
}
type HTTPErrorRes struct {
	Error HTTPError `json:"error"`
}

func CustomError(message string, code Error) HTTPErrorRes {
	e := HTTPError{
		Message: message, Code: code,
	}
	return HTTPErrorRes{
		Error: e,
	}

}