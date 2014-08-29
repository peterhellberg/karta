package main

import (
	"log"
	"net/http"
)

type context struct {
	Logger *log.Logger
}

func (c *context) log(a ...interface{}) {
	c.Logger.Println(a...)
}

func (c *context) logf(f string, a ...interface{}) {
	c.Logger.Printf(f, a...)
}

type handler func(*context, *http.Request, http.ResponseWriter) error

func baseHandler(ctx *context, fn handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := fn(ctx, r, w)
		if err != nil {
			ctx.Logger.Printf("baseHandler: uri=%s err=%s", r.RequestURI, err)
			w.WriteHeader(500)
		}
	})
}

// ListenAndServe creates a context, registers all handlers
// using the provided routes function, and starts listening on the port
func ListenAndServe(l *log.Logger, port string, routes func(*context)) error {
	routes(&context{Logger: l})

	l.Printf("Listening on http://0.0.0.0:%s\n", port)

	return http.ListenAndServe(":"+port, nil)
}
