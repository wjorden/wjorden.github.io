package main

import (
    "net/http"

    "github.com/a-h/templ"
    "github.com/CaffeineOrDeath/CaffeineOrDeath.github.io/internal/views"
)

func main() {
    component := views.Layout("Home")
    http.Handle("/", templ.Handler(component))
}
