package router

import (
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	return mux.NewRouter().StrictSlash(true)
}
