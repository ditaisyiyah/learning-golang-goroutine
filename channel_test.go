package goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

/*
Channels are a typed conduit through
which you can send and receive values with the channel operator (<-)

By default, sends and receives block until the other side is ready
*/

/*
func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // finally, close channel

	go func() {
		channel <- "Dita Larasati" // in channel
		fmt.Println("Sending data inside channel")
	}()

	data := <-channel // out channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}
*/

// channel by default passing by reference
func GiveMeResponse(channel chan string) {
	channel <- "Dita Larasati" // in channel
	fmt.Println("Sending data inside channel")
}

/*
func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // finally, close channel

	go GiveMeResponse(channel)

	data := <-channel // out channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}
*/

// channel behave could be defined while it is as a paramter
func OnlyIn(channel chan<- string) {
	channel <- "Dita Larasati" // in channel
	fmt.Println("Sending data inside channel")
}

func OnlyOut(channel <-chan string) {
	data := <-channel // out channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // finally, close channel

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// buffered channel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 1) // set the size of channel
	defer close(channel)            // finally, close channel

	channel <- "Dita Larasati 1"
	channel <- "Dita Larasati 2" // will be error

	// data := <-channel // out channel
	// fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// range channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 1; i <= 10; i++ {
			channel <- "Dita Larasati " + strconv.Itoa(i)
		}

		// close the channel after finish send all the values
		close(channel)
	}()

	// then receive all the values by looping the channel
	for data := range channel {
		fmt.Println(data)
	}

	time.Sleep(5 * time.Second)
}

// (default) select-case channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data masuk dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data masuk dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}

		if counter == 2 {
			break
		}
	}

	time.Sleep(5 * time.Second)
}
