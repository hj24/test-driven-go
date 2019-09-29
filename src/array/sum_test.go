package array

import "testing"

func TestSum(t *testing.T) {
	t.Run("使用数组", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("want %d got %d", got, want)
		}
	})
	t.Run("使用切片", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		got := SliceSum(slice)
		want := 15
		if got != want {
			t.Errorf("want %d got %d", got, want)
		}
	})
}
