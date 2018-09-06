package goro_test

import (
	"fmt"
	"sync"

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
