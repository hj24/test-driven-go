package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 6)
	want := "aaaaaa"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

/*
基准测试，测单个操作用时，b.N自动确定
运行 go test -bench=. 执行
*/
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}

// 示例，// Output: "aaaaaa"，删除该注释则不可运行
func ExampleRepeat() {
	str := Repeat("a", 6)
	fmt.Println(str)
	// Output: aaaaaa
}
