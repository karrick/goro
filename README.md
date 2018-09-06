# goro

Go library that provides sync.Once capability that can be reset.

## Usage

Documentation is available via
[![GoDoc](https://godoc.org/github.com/karrick/goro?status.svg)](https://godoc.org/github.com/karrick/goro).

```Go
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
```
