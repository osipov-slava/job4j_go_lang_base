package tracker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Tracker(t *testing.T) {
	t.Parallel()

	t.Run("error update - not found", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "First Item",
		}

		err := tracker.UpdateItem(item)
		assert.ErrorIs(t, err, ErrNotFound)
	})

	t.Run("update - success", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "First Item",
		}
		err := tracker.AddItem(item)
		assert.NoError(t, err)

		item = Item{
			ID:   "1",
			Name: "Second Item",
		}
		err = tracker.UpdateItem(item)
		assert.NoError(t, err)
		assert.Len(t, tracker.items, 1)
		assert.Equal(t, item, tracker.items[0])
	})

	t.Run("error add - not unique", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "First Item",
		}

		err := tracker.AddItem(item)
		err = tracker.AddItem(item)
		assert.ErrorIs(t, err, ErrIdNoUnique)
	})
}
