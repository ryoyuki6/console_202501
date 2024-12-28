package main

import (
	"net/http"
	"github.com/ryoyuki6/console_202501/backend/handler"
)

func main() {
	http.HandleFunc("/", handler.Router)
	http.ListenAndServe(":3000", nil)
}
