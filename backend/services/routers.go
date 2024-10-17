package services

import (
	"net/http"
)

func UseHttpRouter(mux *http.ServeMux) {
	mux.Handle("/api/subscribeNewsletter", JsonContentMiddleware(http.HandlerFunc(subscribeNewsletter)))
	mux.Handle("/api/matchesResult", JsonContentMiddleware(http.HandlerFunc(GetMatchesResultsFromDB)))
	mux.Handle("/api/auth/login", JsonContentMiddleware(http.HandlerFunc(login)))
	mux.Handle("/api/auth/signup", JsonContentMiddleware(http.HandlerFunc(signup)))
	mux.Handle("/api/chat/sendMessage", JsonContentMiddleware(http.HandlerFunc(SendMessageService)))
	mux.Handle("/api/chat", JsonContentMiddleware(AuthGate(http.HandlerFunc(GetChat))))
	mux.Handle("/api/chat/addUser", JsonContentMiddleware(AuthGate(http.HandlerFunc(AddUserToChat))))
	mux.Handle("/api/chat/removeUser", JsonContentMiddleware(AuthGate(http.HandlerFunc(RemoveUserFromChat))))
}