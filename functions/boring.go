package functions

import (
	"fmt"
	"math/rand"
	"time"
)

// Boring function will print a message with a counter in random periods
func Boring(msg string) {
	for i :=0; ; i++ {
		fmt.Printf("%s %d\n", msg, i)
		time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
	}
}

// BoringWithChannelInput will write message to the channel
func BoringWithChannelInput(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
	}
}

// BoringReturnsChannel returns a channel which will in turn stream messages
func BoringReturnsChannel(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
		}
	}()
	return c
}


// FanIn will receive two channels and will return one
func FanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() { for { c <- <-input1 } }()
	go func() { for { c <- <-input2 } }()
	return c
}

// Message struct is used dor synchronizing fan in function
type Message struct {
	Msg string
	Wait chan bool
}

// BoringMessage returns a receive only Message channel
func BoringMessage(msg string) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return c
}

// FanInMessage receives 2 receive only Message channels and returns single
func FanInMessage(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() { for { c <- <-input1 } }()
	go func() { for { c <- <-input2 } }()
	return c
}

// FanInUsingSelect uses a select block to synchronize messages from 2 channels
func FanInUsingSelect(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()

	return c
}

// BoringOnQuit stops printing messages till it receives a value on the quit channel
func BoringOnQuit(msg string, quit <-chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			case <-quit:
				return
			}
		}
	}()
	return c
}

// BoringCleanup will cleanup just before program exits
func BoringCleanup(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ;i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			case <-quit:
				cleanup()
				quit <- "See you"
				return
			}
		}
	}()

	return c
}

func cleanup() {
	fmt.Println("Cleaning up....")
	time.Sleep(1 * time.Second)
	fmt.Println("Cleanup completed")
}

