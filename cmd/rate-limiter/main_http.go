package main

import (
	"fmt"
	"net/http"
	"time"

	ratelimiter "github.com/alexeysamorodov/rate-limiter/internal/app/rate-limiter"
	"github.com/gorilla/mux"
)

func RateLimiterMiddleware(rl *ratelimiter.RPSRateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if rl.Allow() {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			}
		})
	}
}

func main2() {
	rl := ratelimiter.NewRPSRateLimiter(2) // 2 запроса в секунду
	defer rl.Stop()

	r := mux.NewRouter()
	r.Use(RateLimiterMiddleware(rl))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! %s", time.Now())
	}).Methods("GET")

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
