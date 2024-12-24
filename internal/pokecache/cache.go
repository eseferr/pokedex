package pokecache

import (
	"sync"
	"time"
)

type Cache struct{
	cacheData map[string]cacheEntry
	mutex sync.Mutex
}
type cacheEntry struct {
	createdAt time.Time
	val []byte
	
}
func (c *Cache)reapLoop(interval time.Duration){
	ticker := time.NewTicker(interval)
	for range ticker.C{
		c.mutex.Lock()
	for key,value := range c.cacheData{
		if time.Since(value.createdAt) > interval{
			delete(c.cacheData, key)
		}
	}
	c.mutex.Unlock()
	}
}
func NewCache(interval time.Duration) *Cache{
	c := &Cache{
		cacheData: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
} 
func (c *Cache)Add(key string, val []byte){
	c.mutex.Lock()
	c.cacheData[key] = cacheEntry{createdAt: time.Now(), val: val}
	c.mutex.Unlock()
}
func (c *Cache)Get(key string)([]byte,bool){
	c.mutex.Lock()
	entry, exists := c.cacheData[key]
	c.mutex.Unlock()
	if !exists{
		return nil, false
	}
	return entry.val, true
}
