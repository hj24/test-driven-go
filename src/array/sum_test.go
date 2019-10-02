package array

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{5, 6})
	want := []int{3, 11}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTail(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %d got %d", got, want)
		}
	}
	t.Run("正常的切片尾部相加", func(t *testing.T) {
		got := SumAllTail([]int{1, 2, 3}, []int{5, 6})
		want := []int{5, 6}
		checkSums(t, got, want)
	})
	t.Run("含空切片的尾部相加", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{1, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
