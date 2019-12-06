package stream



import (
	"testing"
)

func BenchmarkCallCompression(b *testing.B) {
	CallCompression()
}