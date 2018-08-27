package tools

import (
  "os"
	"encoding/gob"
	"github.com/gorilla/sessions"
)

var (
	Store *sessions.CookieStore
)

func Init() error {
  INITIAL_COOKIE := []byte(os.Getenv("INITIAL_COOKIE"))
	Store = sessions.NewCookieStore(INITIAL_COOKIE)
	gob.Register(map[string]interface{}{})
	return nil
}
