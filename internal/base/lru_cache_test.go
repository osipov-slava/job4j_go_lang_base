package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LruCache(t *testing.T) {
	t.Parallel()

	t.Run("Put and Get - simple operations", func(t *testing.T) {
		t.Parallel()

		cache := NewLruCache(3)

		cache.Put("key1", "value1")
		cache.Put("key2", "value2")
		cache.Put("key3", "value3")

		val := cache.Get("key2")
		assert.NotNil(t, val)
		assert.Equal(t, "value2", *val)

		val = cache.Get("key1")
		assert.NotNil(t, val)
		assert.Equal(t, "value1", *val)

		val = cache.Get("key3")
		assert.NotNil(t, val)
		assert.Equal(t, "value3", *val)
	})

	t.Run("Get non-existent key - returns nil", func(t *testing.T) {
		t.Parallel()

		cache := NewLruCache(2)
		cache.Put("key1", "value1")

		val := cache.Get("key2")
		assert.Nil(t, val)
	})

	t.Run("Move to head on Get", func(t *testing.T) {
		t.Parallel()

		cache := NewLruCache(3)
		cache.Put("key1", "value1")
		cache.Put("key2", "value2")
		cache.Put("key3", "value3")

		// key1 must be in the tail
		// after get() key1 move to the head
		cache.Get("key1")

		assert.Equal(t, "key1", cache.Head.Key)
		assert.Equal(t, "key3", cache.Head.Next.Key)
		assert.Equal(t, "key2", cache.Head.Next.Next.Key)
		assert.Equal(t, "key2", cache.Tail.Key)
	})

	t.Run("Put evicts oldest when full", func(t *testing.T) {
		t.Parallel()

		cache := NewLruCache(2)
		cache.Put("key1", "value1")
		cache.Put("key2", "value2")
		cache.Put("key3", "value3") // key1

		// key1 must evict
		val := cache.Get("key1")
		assert.Nil(t, val)

		val = cache.Get("key2")
		assert.NotNil(t, val)
		assert.Equal(t, "value2", *val)

		val = cache.Get("key3")
		assert.NotNil(t, val)
		assert.Equal(t, "value3", *val)
	})

	t.Run("Get moves recently used to front", func(t *testing.T) {
		t.Parallel()

		cache := NewLruCache(3)
		cache.Put("key1", "value1")
		cache.Put("key2", "value2")
		cache.Put("key3", "value3")

		cache.Get("key2")
		cache.Get("key1")

		// expected order: key1, key2, key3
		assert.Equal(t, "key1", cache.Head.Key)
		assert.Equal(t, "key2", cache.Head.Next.Key)
		assert.Equal(t, "key3", cache.Head.Next.Next.Key)

		cache.Put("key4", "value4")

		// key3 must evict
		assert.Nil(t, cache.Get("key3"))
		assert.NotNil(t, cache.Get("key1"))
		assert.NotNil(t, cache.Get("key2"))
		assert.NotNil(t, cache.Get("key4"))
	})

	t.Run("Get on empty cache returns nil", func(t *testing.T) {
		t.Parallel()

		cache := NewLruCache(5)
		val := cache.Get("any_key")
		assert.Nil(t, val)
	})

}
