package chizl

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
)

func Logger(ll zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			defer func() {
				ll.Info().
					Fields(map[string]interface{}{
						"remote_ip": r.RemoteAddr,
						"url":       r.URL.Path,
						"proto":     r.Proto,
						"method":    r.Method,
						"status":    ww.Status(),
					}).
					Msg("request")
			}()

			next.ServeHTTP(ww, r)
		})
	}
}
