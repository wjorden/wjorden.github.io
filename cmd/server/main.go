package main

import (
	"net/http"

	"github.com/CaffeineOrDeath/CaffeineOrDeath.github.io/internal/views"
	"github.com/a-h/templ"
)

func main() {
    component := views.Layout("Home")
    http.Handle("/", templ.Handler(component))
	
	http.ListenAndServe(":8080", nil)
}
