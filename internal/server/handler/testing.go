package handler

import (
	"io"
	"net/http"
)

func Testing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "yayayayaa")
}
