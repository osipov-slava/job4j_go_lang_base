package base

func Mono(nums []int) bool {
	size := len(nums)
	if size == 0 || size == 1 {
		return true
	}
	direction := 0

	if nums[0] < nums[size-1] {
		direction = 1
	} else if nums[0] > nums[size-1] {
		direction = -1
	}

	for i := 1; i < size-1; i++ {
		if direction == 0 && (nums[i-1] != nums[i] || nums[i] != nums[size-1]) {
			return false
		}
		if direction == 1 && (nums[i-1] > nums[i] || nums[i] > nums[size-1]) {
			return false
		}
		if direction == -1 && (nums[i-1] < nums[i] || nums[i] < nums[size-1]) {
			return false
		}
	}
	return true
}
