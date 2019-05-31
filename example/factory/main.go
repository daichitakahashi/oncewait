package main

import (
	"fmt"
	"time"

	"github.com/daichitakahashi/oncewait"
)

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
