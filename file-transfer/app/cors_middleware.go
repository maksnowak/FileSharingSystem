package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

type CORSMiddleware struct {
	allowedOrigins   []string
	allowedMethods   []string
	allowedHeaders   []string
	allowCredentials bool
}

func NewCORSMiddleware(allowedOrigins, allowedMethods, allowedHeaders []string, allowCredentials bool) *CORSMiddleware {
	return &CORSMiddleware{
		allowedOrigins:   allowedOrigins,
		allowedMethods:   allowedMethods,
		allowedHeaders:   allowedHeaders,
		allowCredentials: allowCredentials,
	}
}

func (m *CORSMiddleware) Func() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", m.getAllowedOrigin(r))
			w.Header().Set("Access-Control-Allow-Methods", m.getAllowedMethods())
			w.Header().Set("Access-Control-Allow-Headers", m.getAllowedHeaders())

			if m.allowCredentials {
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func (m *CORSMiddleware) getAllowedOrigin(r *http.Request) string {
	origin := r.Header.Get("Origin")
	if origin == "" {
		return "*"
	}

	for _, allowedOrigin := range m.allowedOrigins {
		if allowedOrigin == "*" || allowedOrigin == origin {
			return origin
		}
	}

	return ""
}

func (m *CORSMiddleware) getAllowedMethods() string {
	if len(m.allowedMethods) == 0 {
		return "GET, POST, PUT, DELETE, OPTIONS"
	}
	return joinStrings(m.allowedMethods)
}

func (m *CORSMiddleware) getAllowedHeaders() string {
	if len(m.allowedHeaders) == 0 {
		return "Content-Type, Authorization"
	}
	return joinStrings(m.allowedHeaders)
}

func joinStrings(slice []string) string {
	result := ""
	for i, str := range slice {
		if i > 0 {
			result += ", "
		}
		result += str
	}
	return result
}

