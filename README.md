# oncewait
[![Go Reference](https://pkg.go.dev/badge/github.com/daichitakahashi/oncewait.svg)](https://pkg.go.dev/github.com/daichitakahashi/oncewait)

Package oncewait offers slightly more synced sync.Once and its factory.

# Usage
```go
func ExampleOnceWaiter(){
    once := oncewait.New()

    processFunc := func() {
        fmt.Println("start heavy process.")

        time.Sleep(time.Second) // do something

        fmt.Println("process finished.")
    }

    var wg sync.WaitGroup

    wg.Add(1)
    go func() {
        once.Do(processFunc)
        fmt.Println("done.")
        wg.Done()
    }()

    wg.Add(1)
    go func() {
        once.Do(processFunc)
        fmt.Println("done.")
        wg.Done()
    }()

    wg.Wait()

    // Output:
    // start heavy process.
    // process finished.
    // done.
    // done.
}
```
