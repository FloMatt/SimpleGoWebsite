package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received a request:", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// NoSurf is a middleware function that protects against Cross-Site Request Forgery (CSRF) attacks.
// It wraps the provided http.Handler with CSRF protection using the nosurf package.
//
// The function creates a new nosurf.CSRFHandler instance with the provided http.Handler as the next handler.
// It then sets the base cookie for the CSRF protection with the following properties:
// - HttpOnly: true, to prevent client-side scripts from accessing the cookie.
// - Path: "/", to apply the cookie to all paths.
// - Secure: app.InProduction, to enable the cookie only for HTTPS connections in production.
// - SameSite: http.SameSiteLaxMode, to restrict the cookie to same-site requests.
//
// Finally, it returns the CSRF-protected http.Handler.
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode})
	return csrfHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
