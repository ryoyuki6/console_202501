package handler

import (
	"fmt"
	"net/http"
)

func Router(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!!!")
}