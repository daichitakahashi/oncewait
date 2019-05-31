# oncewait
Package oncewait offers slightly more synced sync.Once and its factory.

# Usage
```Go
func main() {
	once := oncewait.New()

	processFunc := func() {
		// do something
		time.Sleep(time.Second * 5)
		// do something

		fmt.Println("Process finished. It should be called only once!")
	}

	fmt.Println("Run channel1 and channel2.")
	fmt.Println("Maybe 5 seconds needed.")

	ch1 := make(chan struct{})
	go func() {
		once.Do(processFunc)
		fmt.Println("channel1 finished.")
		close(ch1)
	}()

	ch2 := make(chan struct{})
	go func() {
		once.Do(processFunc)
		fmt.Println("channel2 finished.")
		close(ch2)
	}()

	<-ch1
	<-ch2
}

// Run channel1 and channel2.
// Maybe 5 seconds needed.
// Process finished. It should be called only once!
// channel2 finished.
// channel1 finished.

```

# Usage(Factory)
```Go
func main() {
	factory := &oncewait.Factory{}

	updateFunc := func() {
		time.Sleep(time.Second * 5)
		fmt.Println("UPDATE SUCCEEDED!")

		factory.Refresh("update")
	}

	fmt.Println("Try update.")
	fmt.Println("Maybe 5 seconds needed.")

	ch1 := make(chan struct{})
	go func() {
		factory.Get("update").Do(updateFunc)
		fmt.Println("channel1 finished.")
		close(ch1)
	}()

	ch2 := make(chan struct{})
	go func() {
		factory.Get("update").Do(updateFunc)
		fmt.Println("channel2 finished.")
		close(ch2)
	}()

	<-ch1
	<-ch2

	fmt.Println("once again!")

	ch3 := make(chan struct{})
	go func() {
		factory.Get("update").Do(updateFunc)
		fmt.Println("last update finished.")
		close(ch3)
	}()
	<-ch3
}

// Try update.
// Maybe 5 seconds needed.
// UPDATE SUCCEEDED!
// channel2 finished.
// channel1 finished.
// once again!
// UPDATE SUCCEEDED!
// last update finished.
```