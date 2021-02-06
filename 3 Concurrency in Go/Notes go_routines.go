package main

import "fmt"
//import "time"
import "sync"

func foo (wg *sync.WaitGroup)  {
	fmt.Println("New_Routine")
	wg.Done()
}

func prod(v1 int, v2 int, c chan int){
	c <- v1 * v2
}

var i int = 0
var wg sync.WaitGroup
var mut sync.Mutex // Similar to Semaphore where the variable is locked so that other routines cant access it.
var on sync.Once // To initialize something once in the program with mutliple sub routines.

func setup(){
	fmt.Println("Init")
}

func inc() {
	mut.Lock()
	i = i + 1
	mut.Unlock()
	wg.Done()
}

/*
func dostuff(){  // used for initialization
	on.Do(setup)
	fmt.Println("hello")
	wg.Done()
}
*/

func dostuff(c1 chan int, c2 chan int){  // Used for Deadlock Example
	<- c1
	c2 <- 1
	wg.Done()
}

func main() {

	// goroutines are similar to thread
	// go run time scheduler does the context switching between the different go routines which are running under on os thread

	/*
	go fmt.Printf("Sub\n") // Early exit. Main routine exits before sub routine finishes
	fmt.Printf("Hello, world!\n")
	// /

	go fmt.Printf("2.2\n") // Delayed exit where we make the Main routine wait so that Sub routine executes. Not advisable
	time.Sleep(30 * time.Millisecond)
	fmt.Printf("2.1\n")
	*/

	/*
	var wg sync.WaitGroup
	wg.Add(1)
	go foo(&wg)
	wg.Wait()
	fmt.Println("Main routine")
	*/

	/*
	c := make(chan int) // Communication between Sub routines using channels (Unbuffered channels which has no capacity to hold data)
	// c := make(chan int, 3) Creates a buffered channel of size 3
	go prod(1, 2, c)
	go prod(3, 4, c)
	a := <- c
	b := <- c
	fmt.Println(a*b)
	*/

	/*
	wg.Add(2) // Example for between 2 go routines Variable sharing i.e incorrect sharing
	go inc()
	go inc()
	wg.Wait()
	fmt.Println(i)
	*/

	/*
	wg.Add(2) // Setup func gets executed only once.
	go dostuff()
	go dostuff()
	wg.Wait()
	*/

	ch1 := make(chan int) // Deadlock
	ch2 := make(chan int)
	wg.Add(2)
	go dostuff(ch1, ch2)
	go dostuff(ch2, ch1)
	wg.Wait()

}
