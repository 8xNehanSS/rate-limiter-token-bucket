package main

import (
	"fmt"
	tokenBucket "rateLimiterTokenBucket/tokenbucket"
	"time"
)

func main() {
	tokenBucket.InitRateLimiter(12, 1)

	// simulate multiple users making requests
	ips := []string{"192.168.1.1", "192.168.1.2", "192.168.1.3", "192.168.1.4", "192.168.1.5"}

	for _, ip := range ips {
		go func(ip string) {
			bucket := tokenBucket.NewTokenBucket(ip) // capacity 10 tokens, refill rate 5 tokens/sec
			for i := 0; i < 15; i++ {
				if bucket.Allow() {
					fmt.Println(i, " : Request from", ip, "allowed")
				} else {
					fmt.Println(i, " : Request from", ip, "denied")
				}
				time.Sleep(100 * time.Millisecond)
			}
		}(ip)
	}

	time.Sleep(2 * time.Second)
}
