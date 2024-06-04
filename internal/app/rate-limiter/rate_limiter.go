package ratelimiter

import (
	"time"
)

// RPSRateLimiter определяет структуру лимитатора скорости на основе запросов в секунду.
type RPSRateLimiter struct {
	requests chan struct{}
	quit     chan struct{}
}

// NewRPSRateLimiter создает новый лимитатор скорости с указанным числом запросов в секунду.
func NewRPSRateLimiter(rps int) *RPSRateLimiter {
	limiter := &RPSRateLimiter{
		requests: make(chan struct{}, rps),
		quit:     make(chan struct{}),
	}

	// Запускаем горутину для добавления токенов в канал requests.
	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(rps))
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				select {
				case limiter.requests <- struct{}{}:
				default:
					// Если канал уже полон, пропускаем токен.
				}
			case <-limiter.quit:
				return
			}
		}
	}()

	return limiter
}

// Allow проверяет, можно ли выполнить запрос. Если можно, возвращает true.
func (l *RPSRateLimiter) Allow() bool {
	select {
	case <-l.requests:
		return true
	default:
		return false
	}
}

// Stop останавливает лимитатор скорости.
func (l *RPSRateLimiter) Stop() {
	close(l.quit)
}
