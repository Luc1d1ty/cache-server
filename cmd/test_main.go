/* package main

import (
	"fmt"
	"time"

	"github.com/Luc1d1ty/cache-server/internal/cache" 
)

func main() {
	c := cache.NewCache()
    quit := make(chan bool)

    c.StartTTLManager(10*time.Second, quit)

    fmt.Println("Setting key1 with TTL 5 seconds")
    c.Set("key1", "value1", 5)
    fmt.Println("Setting key2 with no TTL")
    c.Set("key2", "value2", 0)

    if val, ok := c.Get("key1"); ok {
        fmt.Printf("Immediately, key1 found: %s\n", val)
    } else {
        fmt.Println("Immediately, key1 not found")
    }

    if val, ok := c.Get("key2"); ok {
        fmt.Printf("Immediately, key2 found: %s\n", val)
    } else {
        fmt.Println("Immediately, key2 not found")
    }

    fmt.Println("Waiting 15 seconds to allow TTL manager to clean expired entries...")
    time.Sleep(15 * time.Second)

    if val, ok := c.Get("key1"); ok {
        fmt.Printf("After wait, key1 found: %s\n", val)
    } else {
        fmt.Println("After wait, key1 has expired or not found")
    }

    metrics := c.GetMetrics()
    fmt.Printf("Cache Metrics - Hits: %d, Misses: %d, Items: %d\n", metrics.Hits, metrics.Misses, metrics.ItemCount)

    quit <- true

    time.Sleep(2 * time.Second)
} */