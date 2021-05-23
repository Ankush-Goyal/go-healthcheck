package main

import (
	"net/http"

	_ "github.com/Ankush-Goyal/go-healthcheck/api"
	_ "github.com/Ankush-Goyal/go-healthcheck/pkg/registry"
)

func main() {
	http.ListenAndServe(":8090", nil)
}
