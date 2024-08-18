package middleware

import (
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

var (
	ipLimiters = make(map[string]*rate.Limiter)
	mu         sync.Mutex
)

func getLimiterForIP(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	if limiter, exists := ipLimiters[ip]; exists {
		return limiter
	}

	limiter := rate.NewLimiter(1, 5)
	ipLimiters[ip] = limiter

	return limiter
}

func getIP(r *http.Request) string {
	ip := r.RemoteAddr

	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		ip = forwarded
	}

	return ip
}

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getIP(r)

		limiter := getLimiterForIP(ip)

		if !limiter.Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
