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

func (t *Tracker) AddItem(item Item) {
	t.items = append(t.items, item)
}

func (t *Tracker) GetItems() []Item {
	res := make([]Item, len(t.items))
	copy(res, t.items)
	return res
}

func (t *Tracker) UpdateItem(id string, newName string) bool {
	for i, item := range t.items {
		if item.ID == id {
			t.items[i].Name = newName
			return true
		}
	}
	return false
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
