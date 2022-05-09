package http

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// ClientId := "775619592737-5o19qjbm0qfjj0nugh8hohs5ar49et0b.apps.googleusercontent.com"
// ClientSecret: "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",
// RedirectURL:  "http://localhost:8080/google_callback",
// Scopes: []string{
// 	"https://www.googleapis.com/auth/userinfo.email",
// 	"https://www.googleapis.com/auth/userinfo.profile",
// },
// Endpoint: google.Endpoint,

func SetupConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     "775619592737-5o19qjbm0qfjj0nugh8hohs5ar49et0b.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8080/google_callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}
