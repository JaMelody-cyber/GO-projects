package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected:%v,got :%v", want, got)
	}
}
func BenchmarkSplit(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Split("sha sha sha sha sah", " ")
	}
}
