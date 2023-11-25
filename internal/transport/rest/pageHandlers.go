package rest

import (
	"html/template"
	"net/http"
)

// HomeHandler
//
//	@Summary		Main page handler
//	@ID				home
//	@Tags pages_demo
//	@Success		200	{string}	string	"ok"
//	@Router			/ [get]
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../../internal/static/index.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// AboutHandler
//
//	@Summary		About page handler
//	@ID				about
//	@Tags pages_demo
//	@Success		200	{string}	string	"ok"
//	@Router			/about [get]
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../../internal/static/about.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// CatalogHandler
//
//	@Summary		Catalog page handler
//	@ID				catalog
//	@Tags pages_demo
//	@Success		200	{string}	string	"ok"
//	@Router			/catalog [get]
func CatalogHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../../internal/static/catalog.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
