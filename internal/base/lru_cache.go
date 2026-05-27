package base

type Node struct {
	Key   string
	Value string
	Prev  *Node
	Next  *Node
}

type LruCache struct {
	size int
	busy int
	Head *Node
	Tail *Node
}

func NewLruCache(size int) *LruCache {
	return &LruCache{
		size: size,
	}
}

func (l *LruCache) Put(key string, value string) {
	newNode := &Node{
		Key:   key,
		Value: value,
		Prev:  nil,
		Next:  nil,
	}

	if l.busy == 0 {
		l.Head = newNode
		l.Tail = l.Head
		l.busy++
		return
	}
	if l.busy > 0 && l.busy < l.size {
		oldHead := l.Head
		l.Head = newNode
		l.Head.Next = oldHead
		newNode.Next = oldHead
		oldHead.Prev = newNode
		l.busy++
		return
	}
	if l.busy == l.size {
		oldHead := l.Head
		l.Head = newNode
		l.Head.Next = oldHead
		newNode.Next = oldHead
		oldHead.Prev = newNode
		l.Tail = l.Tail.Prev
		l.Tail.Next = nil
		return
	}
}

func (l *LruCache) Get(key string) *string {
	cursorNode := l.Head
	for i := 0; i < l.busy; i++ {
		if cursorNode.Key == key {
			result := cursorNode.Value

			// AI code
			if l.Head != cursorNode {
				if cursorNode.Prev != nil {
					cursorNode.Prev.Next = cursorNode.Next
				}
				if cursorNode.Next != nil {
					cursorNode.Next.Prev = cursorNode.Prev
				}

				if l.Tail == cursorNode {
					l.Tail = cursorNode.Prev
				}

				oldHead := l.Head
				l.Head = cursorNode
				cursorNode.Next = oldHead
				cursorNode.Prev = nil
				oldHead.Prev = cursorNode
			}

			return &result

		}
		cursorNode = cursorNode.Next
	}
	return nil
}
