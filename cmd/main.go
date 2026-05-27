package main

import (
	"fmt"

	"job4j.ru/go-lang-base/internal/base"
)

func main() {
	cache := base.NewLruCache(3)

	cache.Put("first", "111")
	cache.Put("second", "222")
	cache.Put("third", "333")
	cache.Put("forth", "444")
	addr := cache.Get("third")
	if addr != nil {
		res := *addr
		fmt.Println(res)
	}
}
