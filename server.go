package main

import (
  "github.com/scorbettUM/headspace/client_application/api_services/authorization_service/middlewares"
  "github.com/scorbettUM/headspace/client_application/api_services/authorization_service/routes"
  "github.com/urfave/negroni"
	"github.com/gorilla/mux"
  "golang.org/x/crypto/acme/autocert"
	"net/http"
  "os"
  "fmt"
)

func StartServer(){

  AUTHORIZATION_SERVICE_PORT := os.Getenv("AUTHORIZATION_SERVICE_PORT")

  r := mux.NewRouter()
  loginRoutes := mux.NewRouter().PathPrefix("/api/auth").Subrouter().StrictSlash(true)

  loginRoutes.Handle("/", negroni.New(
    negroni.Wrap(http.HandlerFunc(routes.GetAuthorization)),
  ))

  r.HandleFunc("/api/auth/callback", routes.Callback)

  r.PathPrefix("/api/auth").Handler(negroni.New(
    negroni.Wrap(middlewares.AuthMiddleware(middlewares.CorsHandler(loginRoutes))),
  ))

  https_listener := autocert.NewListener(AUTHORIZATION_SERVICE_PORT)

  server := &http.Server{
		Handler: middlewares.CorsHandler(r),
	}

	fmt.Println(server.Serve(https_listener))
}
