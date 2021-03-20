package middlewares

import (
	"log"
	"net/http"
)

var SetMiddlewareLogger = func(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)
		next(w, r)
	}
}

var SetMiddlewareJSON = func(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

type Middleware func(next http.HandlerFunc) http.HandlerFunc

func BuildMiddlewareChain(f http.HandlerFunc, m []Middleware) http.HandlerFunc {
	// if our chain is done, use the original handlerfunc
	if len(m) == 0 {
		return f
	}
	// otherwise nest the handlerfuncs
	// return m[0](BuildMiddlewareChain(f, m[1:cap(m)]...))
	return m[0](BuildMiddlewareChain(f, m[1:cap(m)]))
}
