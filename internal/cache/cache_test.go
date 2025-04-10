package cache

import (
    "testing"
    "time"
)

func TestCacheSetAndGet(t *testing.T) {
    c := NewCache()
    c.Set("testKey", "testValue", 10) 

    value, ok := c.Get("testKey")
    if !ok {
        t.Error("Expected key 'testKey' to be found, but it was not")
    }
    if value != "testValue" {
        t.Errorf("Expected value 'testValue', but got '%s'", value)
    }
}

func TestCacheExpiration(t *testing.T) {
    c := NewCache()
    c.Set("expireKey", "expireValue", 1) 
    time.Sleep(2 * time.Second)
    _, ok := c.Get("expireKey")
    if ok {
        t.Error("Expected key 'expireKey' to be expired, but it was found")
    }
}

func TestCacheDelete(t *testing.T) {
    c := NewCache()
    c.Set("deleteKey", "deleteValue", 0)

    if err := c.Delete("deleteKey"); err != nil {
        t.Errorf("Error deleting key 'deleteKey': %v", err)
    }
    _, ok := c.Get("deleteKey")
    if ok {
        t.Error("Expected key 'deleteKey' to be deleted, but it was still found")
    }
}

func TestCleanupExpired(t *testing.T) {
    c := NewCache()
    c.Set("key1", "value1", 1)
    c.Set("key2", "value2", 0)
    time.Sleep(2 * time.Second)
    removed := c.CleanupExpired()
    if removed != 1 {
        t.Errorf("Expected 1 expired item to be removed, but removed %d", removed)
    }
    
    if _, ok := c.Get("key2"); !ok {
        t.Error("Expected key 'key2' to still be present, but it was not found")
    }
}
