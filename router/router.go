package router

import (
	"github.com/gorilla/mux"
)

func Router(mux *mux.Router) {
	User(mux)
}
