package middleware

import (
	"net/http"
	"github.com/gorilla/securecookie"
)

var s = securecookie.New([]byte("your-secret-key-here"), nil) // Initialize the securecookie instance with a secret key

func CSRFMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            next.ServeHTTP(w, r)
            return
        }

        // Extract the CSRF token from the form
        clientToken := r.FormValue("csrf_token")

        // Retrieve the CSRF token from the cookie
        cookie, err := r.Cookie("csrf_token")
        if err != nil {
            http.Error(w, "CSRF token not found", http.StatusForbidden)
            return
        }

        // Decode the CSRF token from the cookie
        var decodedToken string
        err = s.Decode("csrf_token", cookie.Value, &decodedToken)
        if err != nil || decodedToken != clientToken {
            http.Error(w, "Invalid CSRF token", http.StatusForbidden)
            return
        }

        // If token is valid, proceed to the next handler
        next.ServeHTTP(w, r)
    })
}
