package main

import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}
