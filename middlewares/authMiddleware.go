package middlewares

import (
  "github.com/auth0-community/go-auth0"
  jose "gopkg.in/square/go-jose.v2"
  "net/http"
  "os"
)


func AuthMiddleware(next http.Handler) http.Handler {

  JWKS_URI := os.Getenv("JWKS_URI")
  AUTH0_API_ISSUER := os.Getenv("AUTH0_API_ISSUER")

  AUTH0_API_AUDIENCE := []string{os.Getenv("AUTH0_CLIENT_ID")}

  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: JWKS_URI}, nil)
    audience := AUTH0_API_AUDIENCE

    configuration := auth0.NewConfiguration(client, audience, AUTH0_API_ISSUER, jose.RS256)
    validator := auth0.NewValidator(configuration, nil)

    _, err := validator.ValidateRequest(r)

    if err != nil {
      next.ServeHTTP(w, r)

    } else {
      next.ServeHTTP(w, r)
    }
  })
}
