package tracker

import (
	"fmt"
	"strings"
)

type Item struct {
	ID   string
	Name string
}

func (i Item) toString() string {
	return fmt.Sprintf("%s\t%s", i.ID, i.Name)
}

type Tracker struct {
	items []Item
}

func NewTracker() *Tracker {
	return &Tracker{}
}

func (t *Tracker) AddItem(item Item) error {
	_, ok := t.indexOf(item.ID)
	if ok {
		return ErrIdNoUnique
	}
	t.items = append(t.items, item)
	return nil
}

func (t *Tracker) GetItems() []Item {
	res := make([]Item, len(t.items))
	copy(res, t.items)
	return res
}

func (t *Tracker) UpdateItem(item Item) error {
	index, ok := t.indexOf(item.ID)
	if !ok {
		return ErrNotFound
	}
	t.items[index] = item
	return nil
}

func (t *Tracker) indexOf(id string) (int, bool) {
	for i, item := range t.items {
		if item.ID == id {
			return i, true
		}
	}
	return -1, false
}

func (t *Tracker) DeleteItem(id string) bool {
	for i, item := range t.items {
		if item.ID == id {
			t.items = append(t.items[:i], t.items[i+1:]...)
			return true
		}
	}
	return false
}

func (t *Tracker) FindItemsByFragment(fragment string) []Item {
	items := make([]Item, 0)
	for _, item := range t.items {
		if strings.Contains(item.Name, fragment) {
			items = append(items, item)
		}
	}
	return items
}
