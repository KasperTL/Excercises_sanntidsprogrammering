// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
    "time"
)


func numberserver(inc, dec <- chan int, get <- chan chan int){
    value := 0

    for {
        select {
        case n := <- inc:
            value += n
        case n := <- dec:
            value -= n
        case reply := <- get:
            reply <- value
            return
        }
    }
}


func incrementing(inc chan int, done chan bool) {
    //TODO: increment i 1000000 times
    for j:=0; j < 1000000; j++ {
        inc <- 1;
    }
    done <- true
}

func decrementing(dec chan int, done chan bool) {
    //TODO: decrement i 1000000 times
    for j:=0; j < 1000000; j++ {
        dec <- 1;
    }
    done <- true
}

func main() {
    // What does GOMAXPROCS do? What happens if you set it to 1?
    // GOMAXPROCS is a setting to decide how many CPU kernels the program is allowed to use. In this case you get a value with 2 kernals and 0 with 1
    runtime.GOMAXPROCS(2)    
	
    // TODO: Spawn both functions as goroutines
    inc := make(chan int)
    dec := make(chan int)
    get := make(chan chan int)
    done := make(chan bool)

    go numberserver(inc, dec, get)

    go incrementing(inc, done)
    go decrementing(dec, done)

    <- done
    <- done

    reply := make(chan int)
    get <- reply
    sum := <- reply
    // We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
    // We will do it properly with channels soon. For now: Sleep.
    time.Sleep(500*time.Millisecond)
    Println("The magic number is:", sum)
}
