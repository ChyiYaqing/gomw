package main

import (
	"fmt"
	"time"

	"github.com/jellydator/ttlcache/v2"
)

var (
	notFound = ttlcache.ErrNotFound
	isClosed = ttlcache.ErrClosed
)

func main() {
	standalone()
}

func standalone() {
	cache := ttlcache.NewCache()
	cache.SetTTL(time.Duration(5 * time.Second))
	cache.SetCacheSizeLimit(100)
	cache.SkipTTLExtensionOnHit(true)
	cache.Set("MyKey", "MyValue")
	cache.Set("MyNumber", 1000)

	if val, err := cache.Get("MyKey"); err != notFound {
		fmt.Printf("Got it: %s\n", val)
	}
	if val, err := cache.Get("MyNumber"); err != notFound {
		fmt.Printf("Got it: %s\n", val)
	}

	time.Sleep(5 * time.Second)

	if val, err := cache.Get("MyKey"); err == notFound {
		fmt.Printf("Could not Got it: %s\n", val)
	}

	cache.Remove("MyNumber")
	cache.Purge()
	cache.Close()

}
