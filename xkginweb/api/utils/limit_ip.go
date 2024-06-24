package utils

import (
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

type IPRateLimiter struct {
	limiterBuckets map[string]*rate.Limiter
	mu             *sync.Mutex
	r              rate.Limit
	b              int
}

func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	return &IPRateLimiter{
		limiterBuckets: make(map[string]*rate.Limiter),
		mu:             &sync.Mutex{},
		r:              r,
		b:              b,
	}
}

func (i *IPRateLimiter) AddIPX(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)
	i.limiterBuckets[ip] = limiter

	return limiter
}

func (lm *IPRateLimiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		limiter, ok := lm.limiterBuckets[ip]
		if !ok {
			limiter = lm.AddIPX(ip)
		}

		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome."))
}

// 在这个例子中，我们允许每秒放入10个令牌，并限制桶的容量为100。
// 这意味着该限制器可以处理每秒最多10个请求，
// 但如果对同一IP地址的请求达到100个，则无法通过请求。同时，我们定义了一个简单的处理程序，它将响应“欢迎”。
func main() {
	limit := rate.Limit(10) // 速率，每秒放入令牌的数量
	capacity := 100         // 容量，桶的大小

	ipRateLimiter := NewIPRateLimiter(limit, capacity)

	http.Handle("/", ipRateLimiter.Limit(http.HandlerFunc(IndexHandler)))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
