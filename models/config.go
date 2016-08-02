package models

import "github.com/hamed1soleimani/goauth"

var Secret = "1cfabccbf188251666dfa066a88864a7"
var MongoURI = "mongodb://localhost:27017"
var Host = "0.0.0.0"
var Port = "3000"

var Google = goauth.OauthConfig{
	ClientID:     "112530680305-3bh1cbm6a09sljbil5qb0b3lshmno1jp.apps.googleusercontent.com",
	ClientSecret: "ccD45Nkikl5pbtk9UB4oINef",
	CallbackURL:  "http://127.0.0.1:3000/auth/google/oauth2callback",
	AuthURL:      "https://accounts.google.com/o/oauth2/auth",
	TokenURL:     "https://accounts.google.com/o/oauth2/token",
	ApiURL:       "https://www.googleapis.com/oauth2/v2/userinfo?fields=email%2Cname%2Cpicture",
	Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile"},
}
