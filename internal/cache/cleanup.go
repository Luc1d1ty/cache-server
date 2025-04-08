package cache

import (
    "log"
    "time"
)

func (c *Cache) StartTTLManager(tickInterval time.Duration, quit <-chan bool) {
    ticker := time.NewTicker(tickInterval)
    go func() {
        defer ticker.Stop() 
        for {
            select {
            case <-ticker.C:
                removed := c.CleanupExpired()
                if removed > 0 {
                    log.Printf("[TTL Manager] Removed %d expired entries", removed)
                }
            case <-quit:
                log.Println("[TTL Manager] Shutting down TTL manager")
                return
            }
        }
    }()
}
