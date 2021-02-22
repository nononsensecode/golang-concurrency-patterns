package functions

import (
	"fmt"
	"math/rand"
	"time"
)

// Result is the result of the search
type Result string

// Search does a search and returns the Result
type Search func(query string) Result

// FakeSearch will implement a fake google search
func FakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

var (
	Web = FakeSearch("web")
	Image = FakeSearch("image")
	Video = FakeSearch("video")
)

var (
	Web1 = FakeSearch("web1")
	Web2 = FakeSearch("web2")
	Image1 = FakeSearch("image1")
	Image2 = FakeSearch("image2")
	Video1 = FakeSearch("video1")
	Video2 = FakeSearch("video2")
)

// GoogleWithoutSync does a sequential search using different kinds of search
func GoogleWithoutSync(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

// GoogleWithSync does a search in concurrent way
func GoogleWithSync(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	for i := 0; i < 3; i++ {
		results = append(results, <-c)
	}

	return
}

// GoogleSyncWithTimeout will do a synchronized search, but will timeout after a certain time
func GoogleSyncWithTimeout(query string) (results []Result) {
	c := make(chan Result)
	timeout := time.After(80 * time.Millisecond)

	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("Timed out")
			return
		}
	}

	return
}

func first(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

// GoogleWithReplicas will query a string in all categories by spawning 2 instances of 
// every category search and will accepts the first result
func GoogleWithReplicas(query string) (results []Result) {
	c := make(chan Result)
	
	go func() { c <- first(query, Web1, Web2) }()
	go func() { c <- first(query, Image1, Image2) }()
	go func() { c <- first(query, Video1, Video2) }()

	timeout := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("Timed out")
			return
		}
	}

	return
}