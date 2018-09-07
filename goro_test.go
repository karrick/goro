package goro_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/karrick/goro"
)

func ExampleGoro() {
	var once goro.Once
	var counter int
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()

			if j == 25 {
				once.Reset()
			}

			once.Do(func() { counter++ })
		}(i)
	}

	fmt.Println("counter:", counter)
	// Output: counter: 2
}

func BenchmarkOnce(b *testing.B) {
	var counter int
	var once goro.Once

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			once.Do(func() { counter++ })
		}
	})

	_ = counter
}
