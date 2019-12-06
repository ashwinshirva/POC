package stream



import (
	"testing"
)

func BenchmarkCallNoStream(b *testing.B) {
	CallNoStream()
}