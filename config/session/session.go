package session

import (
	"os"

	"github.com/gorilla/sessions"
)

// Store will hold all session data
var Store *sessions.CookieStore

func init() {

	secretKey := os.Getenv("SessionSecretKey")
	cookieStore := sessions.NewCookieStore([]byte(secretKey))

	cookieStore.Options = &sessions.Options{
		MaxAge:   60 * 100, //100 minutes   //try that 3 * time.Hour
		HttpOnly: false,
	}

	Store = cookieStore
}
