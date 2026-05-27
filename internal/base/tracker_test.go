package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Tracker(t *testing.T) {
	t.Parallel()

	t.Run("check link leak", func(t *testing.T) {
		t.Parallel()

		tracker := base.NewTracker()
		item := base.Item{
			ID:   "1",
			Name: "First Item",
		}
		tracker.AddItem(item)

		res := tracker.GetItems()
		res[0].Name = "Second Item"

		assert.Equal(t,
			[]base.Item{item},
			tracker.GetItems(),
		)
	})

	t.Run("NewTracker creates empty tracker", func(t *testing.T) {
		t.Parallel()

		tracker := base.NewTracker()

		assert.NotNil(t, tracker)
		items := tracker.GetItems()
		assert.Equal(t, 0, len(items))
	})

	t.Run("AddItem adds single item to tracker", func(t *testing.T) {
		t.Parallel()

		tracker := base.NewTracker()
		item := base.Item{
			ID:   "1",
			Name: "Item1",
		}

		tracker.AddItem(item)
		items := tracker.GetItems()

		assert.Equal(t, 1, len(items))
		assert.Equal(t, "1", items[0].ID)
		assert.Equal(t, "Item1", items[0].Name)
	})

	t.Run("Empty tracker GetItems returns empty slice", func(t *testing.T) {
		t.Parallel()

		tracker := base.NewTracker()
		items := tracker.GetItems()

		assert.NotNil(t, items)
		assert.Equal(t, 0, len(items))
	})

}
