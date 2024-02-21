package handler

import (
	"github.com/a-h/templ"
	"net/http"
)

func render(w http.ResponseWriter, r *http.Request, component templ.Component) {
	err := component.Render(r.Context(), w)
	ifErrorThenPanic(err)
}

func ifErrorThenPanic(err error) {
	if err != nil {
		panic(err)
	}
}
