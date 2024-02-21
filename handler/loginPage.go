package handler

import (
	"fin/view/layout"
	"net/http"
)

func renderLoginPage(w http.ResponseWriter, r *http.Request) {
	render(w, r, layout.Base())
}
