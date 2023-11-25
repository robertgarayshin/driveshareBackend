package app

import (
	"driveshare_backend/internal/service"
	"driveshare_backend/internal/transport/rest"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"log"
	"net/http"
)

func CreateApp() {
	log.Println("Starting server...")
	r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../internal/static"))))
	r.HandleFunc("/", rest.HomeHandler)
	r.HandleFunc("/catalog", rest.CatalogHandler)
	r.HandleFunc("/about", rest.AboutHandler)

	r.HandleFunc("/auth/signup", rest.SignupHandler).Methods(http.MethodPost)
	r.HandleFunc("/auth/signin", rest.SigninHandler).Methods(http.MethodPost)

	r.HandleFunc("/user/{id}", service.VerifyJWT(rest.GetUserByIdHandler)).Methods(http.MethodGet)
	r.HandleFunc("/verify/{id}/token:{token}", rest.EmailConfirmHandler).Methods(http.MethodGet)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
