package web

import (
	"math/rand"
	"net/http"
	"strconv"
)

// IssueSession attaches the session cookie to a response.
func IssueSession(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	})
}

// CSRFNonce is embedded into rendered forms.
func CSRFNonce() string {
	return strconv.FormatInt(rand.Int63(), 16)
}
