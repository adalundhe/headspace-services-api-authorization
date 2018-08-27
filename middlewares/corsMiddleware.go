package middlewares

import (
  "os"
  "net/http"
)

func CorsHandler(h http.Handler) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {

    CLIENT_PORT := os.Getenv("CLIENT_PORT")
    CLIENT_HEADERS := os.Getenv("CLIENT_HEADERS")

    if (r.Method == "OPTIONS") {
      w.Header().Set("Access-Control-Allow-Origin", CLIENT_PORT)
      w.Header().Set("Access-Control-Allow-Headers", CLIENT_HEADERS)
    } else {
      w.Header().Set("Access-Control-Allow-Origin", CLIENT_PORT)
      w.Header().Set("Access-Control-Allow-Headers", CLIENT_HEADERS)
      h.ServeHTTP(w,r)
    }
  }
}
