package routes

import (
  "encoding/json"
  "github.com/scorbettUM/headspace/client_application/api_services/authorization_service/tools"
	"net/http"
)

func GetAuthorization(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
	session, err := tools.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

  js, err := json.Marshal(session.Values["login"])
  if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  w.Write(js)
}
