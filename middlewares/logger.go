package middlewares

import (
	"log"
	"net/http"
	"time"
)

// Logger is a middleware handler that does request logging
type Logger struct {
	handler http.Handler
}

// ServeHTTP handles the request by passing it to the real
// handler and logging the request details
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// if !AuthGuard(r) {
	// 	handlers.ReturnAuthFail(w)
	// 	return
	// }

	l.handler.ServeHTTP(w, r)

	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

// NewLogger constructs a new Logger middleware handler
func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}
