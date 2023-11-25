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

	R := mux.NewRouter()

	R.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	R.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../internal/static"))))
	R.HandleFunc("/", rest.HomeHandler)
	R.HandleFunc("/catalog", rest.CatalogHandler)
	R.HandleFunc("/about", rest.AboutHandler)

	R.HandleFunc("/auth/signup", rest.SignupHandler).Methods(http.MethodPost)
	R.HandleFunc("/auth/signin", rest.SigninHandler).Methods(http.MethodPost)

	R.HandleFunc("/user/{id}", service.VerifyJWT(rest.GetUserByIdHandler)).Methods(http.MethodGet)
	R.HandleFunc("/user/{id}", service.VerifyJWT(rest.EditProfileHandler)).Methods(http.MethodPut)
	R.HandleFunc("/user/{id}", service.VerifyJWT(rest.DeleteProfileHandler)).Methods(http.MethodDelete)

	R.HandleFunc("/verify/{id}/token:{token}", rest.EmailConfirmHandler).Methods(http.MethodGet)

	R.HandleFunc("/car/new", service.VerifyJWT(rest.PostCarHandler)).Methods(http.MethodPost)
	R.HandleFunc("/car/{id}", service.VerifyJWT(rest.GetCarByIdHandler)).Methods(http.MethodGet)
	R.HandleFunc("/car/{id}", service.VerifyJWT(rest.DeleteCarHandler)).Methods(http.MethodDelete)
	R.HandleFunc("/car", service.VerifyJWT(rest.GetAllCarsHandler)).Methods(http.MethodGet)
	R.HandleFunc("/car/book/{id}", service.VerifyJWT(rest.BookCarHandler)).Methods(http.MethodPut)
	R.HandleFunc("/car/cancel/{id}", service.VerifyJWT(rest.CancelBookHandler)).Methods(http.MethodPut)
	R.HandleFunc("/car/available", service.VerifyJWT(rest.GetAvailableCarsHandler)).Methods(http.MethodGet)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe("localhost:8080", R))
}
