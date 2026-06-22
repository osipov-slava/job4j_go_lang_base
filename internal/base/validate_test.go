package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Validate(t *testing.T) {
	t.Parallel()

	t.Run("ValidateRequest is undefined", func(t *testing.T) {
		t.Parallel()

		var req *base.ValidateRequest
		rsl := base.Validate(req)

		assert.Equal(t, 1, len(rsl))
		assert.Equal(t, "ValidateRequest is undefined!", rsl[0])
	})

	t.Run("All fields are undefined", func(t *testing.T) {
		t.Parallel()

		req := base.ValidateRequest{
			UserID:      "",
			Title:       "",
			Description: "",
		}
		rsl := base.Validate(&req)

		assert.Equal(t, 3, len(rsl))
		assert.Equal(t, "Description is empty!", rsl[0])
		assert.Equal(t, "UserID is empty!", rsl[1])
		assert.Equal(t, "Title is empty!", rsl[2])
	})

	t.Run("Some fields are undefined", func(t *testing.T) {
		t.Parallel()

		req := base.ValidateRequest{
			UserID:      "Kate",
			Title:       "",
			Description: "",
		}
		rsl := base.Validate(&req)

		assert.Equal(t, 2, len(rsl))
		assert.Equal(t, "Description is empty!", rsl[0])
		assert.Equal(t, "Title is empty!", rsl[1])
	})

	t.Run("Correct struct", func(t *testing.T) {
		t.Parallel()

		req := base.ValidateRequest{
			UserID:      "Kate",
			Title:       "SomeTitle",
			Description: "SomeDescription",
		}
		rsl := base.Validate(&req)

		assert.Equal(t, 0, len(rsl))
	})

}
