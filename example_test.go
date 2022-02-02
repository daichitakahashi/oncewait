package oncewait

import (
	"fmt"
	"sync"
	"time"
)

func ExampleOnceWaiter() {
	once := New()

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

func ExampleFactory() {
	var factory Factory
	const key = "example"

	processFunc := func() {
		defer factory.Refresh(key)
		fmt.Println("start heavy process.")

		time.Sleep(time.Second) // do something

		fmt.Println("process finished.")
	}

	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			factory.Get(key).Do(processFunc)
			fmt.Println("done.")
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			factory.Get(key).Do(processFunc)
			fmt.Println("done.")
			wg.Done()
		}()

		wg.Wait()
	}

	// Output:
	// start heavy process.
	// process finished.
	// done.
	// done.
	// start heavy process.
	// process finished.
	// done.
	// done.
}
